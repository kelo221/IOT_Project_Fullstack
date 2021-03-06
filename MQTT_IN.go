package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"time"
)

type dataPackageIn struct {
	Nr       int  `json:"Nr,omitempty"`
	Speed    int  `json:"speed,omitempty"`
	Setpoint int  `json:"Setpoint,omitempty"`
	Pressure int  `json:"pressure,omitempty"`
	Auto     bool `json:"auto,omitempty"`
	Err      bool `json:"err,omitempty"`
	UnixTime int  `json:"UnixTime,omitempty"`
}

var MQTTpackage dataPackageIn

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	MQTTpackage = dataPackageIn{
		Nr:       0,
		Speed:    0,
		Setpoint: 0,
		Pressure: 0,
		Auto:     false,
		Err:      false,
		UnixTime: 0,
	}

	err := json.Unmarshal(msg.Payload(), &MQTTpackage)
	if err != nil {
		logrus.Error(err)
	} else {

		//Add current time to MQTTpackage
		MQTTpackage.UnixTime = int(time.Now().Unix())

		fmt.Println(MQTTpackage)
		appendToDB(MQTTpackage)

	}

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client) {

	for {
		token := client.Publish("controller/status", 0, false, nil)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "controller/status"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}

func handleMQTTIn() {

	var broker = "localhost"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	client.Disconnect(250)

}
