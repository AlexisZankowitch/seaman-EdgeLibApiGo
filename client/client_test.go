package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeConnecting(t *testing.T) {
	seal := &SealEdgeApi{}

	seal.Connect("test", 65535)

	assert.Equal(t, seal.metadata.namespace, "foo")
}

func Test_TopicsShouldBeCorrect(t *testing.T) {
	builder := SealEdgeApiBuilder{}

	sealClient, _ := builder.
		WithComponentName("component_name").
		WithModuleName("module_name").
		WithModuleVersion("module_version").
		WithNamespace("namespace").
		Build()

	assert.Equal(t, "namespace/component_name/api/v1/module_name/module_version", sealClient.baseTopic())
	assert.Equal(t, "namespace/component_name/api/v1/module_name/module_version/status", sealClient.statusTopic())
}
