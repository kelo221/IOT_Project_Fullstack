package main

func main() {

	handleDatabase()
	createAccounts()
	go handleMQTTIn()

	handleHTTP()

}
