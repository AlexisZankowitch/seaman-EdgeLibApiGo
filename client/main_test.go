package client

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	sealMqtt "concept/a.zankowitch/seal-edge-lib-go/mqtt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_Integration(t *testing.T) {
	ctx := context.Background()
	// Mosquitto configuration to allow anonymous connections
	mosquittoConf := "listener 1883\n" +
		"allow_anonymous true\n"

	req := testcontainers.ContainerRequest{
		Image:        "eclipse-mosquitto:2.0.20",
		ExposedPorts: []string{"1883/tcp"},
		WaitingFor:   wait.ForListeningPort("1883/tcp"),
		// Provide the custom configuration to the container
		Files: []testcontainers.ContainerFile{
			{
				Reader:            strings.NewReader(mosquittoConf),
				ContainerFilePath: "/mosquitto/config/mosquitto.conf",
				FileMode:          0o644,
			},
		},
		// Use the provided configuration file
		Cmd: []string{"mosquitto", "-c", "/mosquitto/config/mosquitto.conf"},
	}
	mosquitto, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer mosquitto.Terminate(ctx)

	if err != nil {
		log.Print(err)
		return
	}

	port, err := mosquitto.MappedPort(ctx, "1883")
	p := uint16(port.Int())
	// p := uint16(1883)
	require.NoError(t, err)

	t.Run("Connect to broker", func(t *testing.T) {
		builder := SealEdgeApiBuilder{}
		messageReceived := false
		sealClient, _ := builder.
			WithComponentName("component_name").
			WithModuleName("module_name").
			WithModuleVersion("module_version").
			WithNamespace("namespace").
			Build()
		defer sealClient.Disconnect()

		mqttClient := sealMqtt.New("127.0.0.1", p)
		mqttClient.Opts.DefaultPublishHandler = func(client mqtt.Client, msg mqtt.Message) {
			messageReceived = string(msg.Payload()) == `{"moduleStatus":"ONLINE"}`
			fmt.Printf("Subscriber received: %s from topic: %s\n", msg.Payload(), msg.Topic())
		}

		mqttClient.Opts.OnConnect = func(client mqtt.Client) {
			fmt.Println("Subscriber Connected to MQTT broker!")
			// Subscribe to the topic once connected
			token := client.Subscribe("#", 1, nil) // topic, qos, messageHandler (nil uses default)
			token.Wait()
			if token.Error() != nil {
				log.Fatalf("Subscriber Failed to subscribe: %v", token.Error())
			}
			fmt.Printf("Subscriber subscribed to topic: %s\n", "#")
		}

		mqttClient.Connect()

		err := sealClient.Connect("127.0.0.1", p)

		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			return messageReceived
		}, 5*time.Second, 100*time.Millisecond, "Should receive MQTT message")
	})
}
