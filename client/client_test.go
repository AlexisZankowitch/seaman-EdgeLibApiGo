package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeConnecting(t *testing.T) {
	seal := &SealEdgeApi{}

	seal.Connect("test", 65535)

	assert.Equal(t, seal.metada.namespace, "foo")
}
