package main

func main() {

	handleDatabase()
	handleLogin()
	go handleMQTTIn()

	handleHTTP()

}
