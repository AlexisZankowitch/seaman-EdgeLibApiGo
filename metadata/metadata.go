package metadata

import "concept/a.zankowitch/seal-edge-lib-go/client"

type Metadata struct {
	namespace     string
	componentName string
	moduleName    string
	moduleVersion string
}

func New(b client.SealEdgeApiBuilder) Metadata {
	return Metadata{
		namespace:     b.Namespace,
		componentName: b.ComponentName,
		moduleName:    b.ModuleName,
		moduleVersion: b.ModuleVersion,
	}
}

func (s *Metadata) GetNameSpace() string {
	return s.namespace
}
