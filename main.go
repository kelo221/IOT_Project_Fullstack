package main

func main() {

	handleDatabase()
	go handleMQTT()

	//aql("FOR x IN IOT_DATA_SENSOR RETURN x")

	handleHTTP()

}
