package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shouldHaveNamespace(t *testing.T) {
	c := &SealEdgeApiBuilder{}
	r := c.WithNamespace("test-namespace")

	assert.Equal(t, "test-namespace", c.Namespace)
	assert.Same(t, c, r)
}

func Test_shouldHaveComponent(t *testing.T) {
	c := &SealEdgeApiBuilder{}
	r := c.WithComponentName("test-component")

	assert.Equal(t, "test-component", c.ComponentName)
	assert.Same(t, c, r)
}

func Test_shouldHaveModuleName(t *testing.T) {
	c := &SealEdgeApiBuilder{}
	r := c.WithModuleName("test-Module_Name")

	assert.Equal(t, "test-Module_Name", c.ModuleName)
	assert.Same(t, c, r)
}

func Test_shouldHaveModuleVersion(t *testing.T) {
	c := &SealEdgeApiBuilder{}
	r := c.WithModuleVersion("test-module_version")

	assert.Equal(t, "test-module_version", c.ModuleVersion)
	assert.Same(t, c, r)
}

func Test_Build(t *testing.T) {
	tests := []struct {
		name          string
		Namespace     string
		ComponentName string
		ModuleName    string
		ModuleVersion string
		shouldSuccess bool
	}{
		{
			shouldSuccess: false,
			name:          "empty",
		},
		{
			Namespace:     "namespace",
			shouldSuccess: false,
			name:          "with namespace",
		},
		{
			Namespace:     "namespace",
			ComponentName: "component",
			shouldSuccess: false,
			name:          "with namespace & component",
		},
		{
			Namespace:     "namespace",
			ComponentName: "component",
			ModuleName:    "module",
			shouldSuccess: false,
			name:          "with namespace & component & module name",
		},
		{
			Namespace:     "namespace",
			ComponentName: "component",
			ModuleName:    "module",
			ModuleVersion: "version",
			shouldSuccess: true,
			name:          "with namespace & component & module name & module version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SealEdgeApiBuilder{}

			seal, err := c.WithNamespace(tt.Namespace).
				WithComponentName(tt.ComponentName).
				WithModuleName(tt.ModuleName).
				WithModuleVersion(tt.ModuleVersion).
				Build()

			if !tt.shouldSuccess {
				assert.Error(t, err, "Build should return an error")
				return
			}

			if tt.shouldSuccess {
				assert.Nil(t, err)
				assert.NotNil(t, seal, "Seal is defined")
			}

			seal.Connect("host string", 1)

		})
	}
}
