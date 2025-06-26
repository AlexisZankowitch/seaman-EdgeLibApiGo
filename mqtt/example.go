package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	brokerURL = "tcp://10.128.2.198:1883" // Your MQTT broker address and port
	clientID  = "go_mqtt_publisher"
	topic     = "test/topic"
	qos       = 1 // Quality of Service: 0 (at most once), 1 (at least once), 2 (exactly once)
)

// Define a message handler for incoming messages (not strictly needed for a publisher, but good practice)
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Publisher received message: %s on topic: %s\n", msg.Payload(), msg.Topic())
}

// Define a connection handler
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Publisher Connected to MQTT broker!")
}

// Define a connection lost handler
var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Publisher Connection Lost: %v\n", err)
}

func main() {
	// Set up MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler
	opts.SetCleanSession(true) // Start a clean session

	// Create and connect the client
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Publisher Failed to connect to MQTT broker: %v", token.Error())
	}

	fmt.Printf("Publisher is sending messages to topic '%s'...\n", topic)

	// Publish messages every second
	for i := 0; ; i++ {
		text := fmt.Sprintf("Hello from Go Publisher! Message %d", i)
		token = client.Publish(topic, qos, false, text) // topic, qos, retained, payload
		token.Wait()
		if token.Error() != nil {
			fmt.Printf("Publisher Failed to publish message: %v\n", token.Error())
		} else {
			fmt.Printf("Publisher sent: %s\n", text)
		}
		time.Sleep(1 * time.Second)
	}
}
