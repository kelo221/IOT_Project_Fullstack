package main

func main() {

	handleDatabase()
	go handleMQTTIn()
	//go handleMQTTOut()

	//	aql("FOR x IN IOT_DATA_SENSOR RETURN x")

	handleHTTP()

}
