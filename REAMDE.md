Questions:

- [ ] Are the namespace, component name, module name and module version always mandatory? If yes why using the builder pattern and not a simple constructor that takes all those arguments?
  e.g:
``` go
sealClient, _ := builder.
	WithComponentName("component_name").
	WithModuleName("module_name").
	WithModuleVersion("module_version").
	WithNamespace("namespace").
	Build()
// could be simply
sealClient := New("component_name", "module_name", "module_version", "namespace")
sealClient.Connect()
```
- [ ] Because the arguments to the builder are used in the topics, we should probably validate them: forbid `#` or `+` usage
- [ ] Do we want dev to be able to change metadata of the seal client once built?
- [ ] In the Seal API > message specification "could look like the following" does it mean the `.build()` is optional ?
- [ ] In the Connect: There has to be the option to set a specific host and port for the Mqtt Client. Doea it mean host and port are optional ? 😅
- [ ] We are using uuid as client ID for MQTT: nice because all client id will be unique, but not nice for identification in the broker logs
- [ ] Do we need to make sure the messages arrive in the correct order? In the doc of go PAHO: `max_inflight_messages - Unless this is set to 1 mosquitto does not guarantee ordered delivery of messages.`
- [ ] Last will - lacking a section in the documentation imo. Not clear what the QoS should be or if it should be retained.
- [ ] In the start flow: `publish to endpoint topic all the methods of the module` ... but when we just created the client, there aren't any methods yet, no ? ... Okay there are alerady existing because of the decorators you guys are using in the Java & Python lib... so the chart is a bit "lying"
- [ ] Are the request/response payload of the methods define somewhere? I guess it will have to be done directly on the seal service 🤔
- [ ] What is the purpose of publishing the methods to the endpoints topic?
- [ ] I understand that on the edge using MQTT might be easier for device to device communication because we don't always know the IP adresses of the devices so how do we tell a service to send an http request to another one? But for device to cloud communication why not using some simple https calls? The edge device proxy method call and forward it to the cloud. The service on the cloud would run on a sub domain of whatever: myService.fht.gea.com ? ... And the edge gateway is simply a gateway for all gea.com request ?
