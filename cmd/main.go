package main

import (
	"concept/a.zankowitch/seal-edge-lib-go/client"
	"log"
)

func main() {
	builder := &client.SealEdgeApiBuilder{}

	sealClient, err := builder.
		WithComponentName("component_name").
		WithModuleName("module_name").
		WithModuleVersion("module_version").
		WithNamespace("namespace").
		Build()

	if err != nil {
		println(err.Error())
	}

	err = sealClient.Connect("localhost", 1883)

	if err != nil {
		log.Print(err)
		return
	}
}
