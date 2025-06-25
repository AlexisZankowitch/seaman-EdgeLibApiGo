package client

import (
	"concept/a.zankowitch/seal-edge-lib-go/builder"
)

type ConnectionState string

const (
	ConnectionStateNeverConnected ConnectionState = "NEVER_CONNECTED"
	ConnectionStateConnecting     ConnectionState = "CONNECTING"
	ConnectionStateConnected      ConnectionState = "CONNECTED"
	ConnectionStateDisconnecting  ConnectionState = "DISCONNECTING"
	ConnectionStateDisconnected   ConnectionState = "DISCONNECTED"
)

type SealEdgeApi struct {
	State  ConnectionState
	metada sealEdgeApiMetadata
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
		metada: sealEdgeApiMetadata{
			namespace:     b.Namespace,
			componentName: b.ComponentName,
			moduleName:    b.ModuleName,
			moduleVersion: b.ModuleVersion,
		},
	}
}

func (seal *SealEdgeApi) Connect(host string, port uint16) error {
	f := &builder.MyFoo{}

	f.GetNameSpace()

	seal.State = ConnectionStateConnecting

	seal.metada.namespace = "foo"

	return nil
}
