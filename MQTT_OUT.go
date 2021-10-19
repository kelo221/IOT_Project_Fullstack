package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func handleMQTTOut(JSONPackage string) {

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")

	opts.SetClientID("CLIENT_ID")

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := c.Publish("controller/settings", 0, false, JSONPackage)
	token.Wait()

}
