package main

import "fmt"

func main() {
	fmt.Println("hello2")
	go handleDatabase()
	handleHTTP()
}
