package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

type dataPackageOut struct { //Either send speed or pressure
	Auto     bool `json:"auto,omitempty"`
	Pressure int  `json:"pressure,omitempty"`
	Speed    int  `json:"speed,omitempty"`
}

var newAuto bool
var newPressure int
var newSpeed int

//test
func handleMQTTOut() {

	newAuto = true
	newPressure = 0
	newSpeed = 0

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")

	opts.SetClientID("CLIENT_ID")

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		for {

			message := dataPackageOut{
				Auto:     newAuto,
				Pressure: newPressure,
				Speed:    newSpeed,
			}
			fmt.Println("sending this: ", message)
			messageJSON, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			token := c.Publish("controller/settings", 0, false, messageJSON)
			token.Wait()
			time.Sleep(time.Second * 1)
		}
	}()

	select {}

}
