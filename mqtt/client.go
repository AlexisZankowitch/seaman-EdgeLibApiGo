package mqtt

import (
	"fmt"
	"log"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MqttClient struct {
	Client mqtt.Client
	Opts   *mqtt.ClientOptions
}

func New(host string, port uint16) *MqttClient {
	opts := mqtt.NewClientOptions()
	brokerUrl := "tcp://" + host + ":" + strconv.FormatUint(uint64(port), 10)

	log.Printf("Broker URL to connect to: %s", brokerUrl)

	opts.AddBroker(brokerUrl)
	uuidStr := uuid.New().String()
	opts.SetClientID(uuidStr)

	return &MqttClient{
		Opts: opts,
	}
}

func (c *MqttClient) Connect() error {
	c.Client = mqtt.NewClient(c.Opts)

	token := c.Client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Client Failed to connect to MQTT broker: %v", token.Error())
		return token.Error()
	}
	log.Println("Connected to MQTT broker.")
	return nil
}

func (c *MqttClient) Publish(topic string, qos byte, retained bool, payload any) error {
	token := c.Client.Publish(topic, qos, retained, payload)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Publisher Failed to publish message: %v\n", token.Error())
		return token.Error()
	}
	fmt.Println("Publisher sent a messaege on topic %v", topic)
	return nil
}

func (c *MqttClient) Subscribe(topic string, qos byte) error {
	token := c.Client.Subscribe(topic, qos, nil)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Subscriber Failed to subscribe: %v\n", token.Error())
		return token.Error()
	}
	fmt.Println("Subscriber subscribed to topic %v", topic)
	return nil
}

func (c *MqttClient) Disconnect() {
	c.Client.Disconnect(1000)
}
