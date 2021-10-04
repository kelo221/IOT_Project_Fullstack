package main

func main() {
	go handleDatabase()
	go handleMQTT()
	handleHTTP()
}
