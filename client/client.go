package client

import (
	"concept/a.zankowitch/seal-edge-lib-go/mqtt"
	"fmt"
	"log"
)

type ConnectionState string

const lastWillMessage = `{"moduleStatus":"OFFLINE"}`
const onlineMessage = `{"moduleStatus":"ONLINE"}`

const (
	ConnectionStateNeverConnected ConnectionState = "NEVER_CONNECTED"
	ConnectionStateConnecting     ConnectionState = "CONNECTING"
	ConnectionStateConnected      ConnectionState = "CONNECTED"
	ConnectionStateDisconnecting  ConnectionState = "DISCONNECTING"
	ConnectionStateDisconnected   ConnectionState = "DISCONNECTED"
)

type SealTopics struct {
}

type SealEdgeApi struct {
	State    ConnectionState
	metadata sealEdgeApiMetadata
	client   *mqtt.MqttClient
}

type sealEdgeApiMetadata struct {
	namespace     string
	componentName string
	moduleName    string
	moduleVersion string
}

func new(b *SealEdgeApiBuilder) *SealEdgeApi {
	return &SealEdgeApi{
		State: ConnectionStateNeverConnected,
		metadata: sealEdgeApiMetadata{
			namespace:     b.Namespace,
			componentName: b.ComponentName,
			moduleName:    b.ModuleName,
			moduleVersion: b.ModuleVersion,
		},
	}
}

func (seal *SealEdgeApi) baseTopic() string {
	return fmt.Sprintf(
		"%s/%s/api/v1/%s/%s",
		seal.metadata.namespace,
		seal.metadata.componentName,
		seal.metadata.moduleName,
		seal.metadata.moduleVersion,
	)
}

func (seal *SealEdgeApi) statusTopic() string {
	return seal.baseTopic() + "/status"
}

func (seal *SealEdgeApi) endpointTopic() string {
	return seal.baseTopic() + "/endpoint"
}

func (seal *SealEdgeApi) Connect(host string, port uint16) error {
	seal.State = ConnectionStateConnecting

	seal.client = mqtt.New(host, port)
	seal.client.Opts.SetWill(seal.statusTopic(), lastWillMessage, 1, true)

	err := seal.client.Connect()
	if err != nil {
		return err
	}

	// Should this be configured as OnConnect?
	err = seal.client.Publish(seal.statusTopic(), 1, false, onlineMessage)
	if err != nil {
		return err
	}

	return nil
}

func (seal *SealEdgeApi) Disconnect() {
	log.Println("Will dicsonnect Seal edge lib client")
	seal.client.Client.Disconnect(1000)
}
