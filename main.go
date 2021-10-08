package main

func main() {
	//getPng()				//TODO	Update image when new sample is received
	go handleDatabase()
	go handleMQTT()
	handleHTTP()
}
