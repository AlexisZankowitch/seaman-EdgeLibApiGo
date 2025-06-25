package client

import (
	"errors"
	"fmt"
	"strings"
)

type SealEdgeApiBuilder struct {
	Namespace     string
	ComponentName string
	ModuleName    string
	ModuleVersion string
}

func (c *SealEdgeApiBuilder) WithNamespace(namespace string) *SealEdgeApiBuilder {
	c.Namespace = namespace
	return c
}

func (c *SealEdgeApiBuilder) WithComponentName(component string) *SealEdgeApiBuilder {
	c.ComponentName = component
	return c
}

func (c *SealEdgeApiBuilder) WithModuleName(moduleName string) *SealEdgeApiBuilder {
	c.ModuleName = moduleName
	return c
}

func (c *SealEdgeApiBuilder) WithModuleVersion(moduleVersion string) *SealEdgeApiBuilder {
	c.ModuleVersion = moduleVersion
	return c
}

func (c *SealEdgeApiBuilder) assertNamespace() error {
	if c.Namespace == "" {
		return errors.New("Namespace must be defined")
	}
	return nil
}

func (c *SealEdgeApiBuilder) assertComponentName() error {
	if c.ComponentName == "" {
		return errors.New("ComponentName must be defined")
	}
	return nil
}

func (c *SealEdgeApiBuilder) assertModuleName() error {
	if c.ModuleName == "" {
		return errors.New("ModuleName must be defined")
	}
	return nil
}

func (c *SealEdgeApiBuilder) assertModuleVersion() error {
	if c.ModuleVersion == "" {
		return errors.New("ModuleVersion must be defined")
	}
	return nil
}

func (c *SealEdgeApiBuilder) Build() (*SealEdgeApi, error) {
	errors := []string{}

	if err := c.assertNamespace(); err != nil {
		errors = append(errors, err.Error())
	}
	if err := c.assertComponentName(); err != nil {
		errors = append(errors, err.Error())
	}
	if err := c.assertModuleName(); err != nil {
		errors = append(errors, err.Error())
	}
	if err := c.assertModuleVersion(); err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
	}

	return new(c), nil
}
