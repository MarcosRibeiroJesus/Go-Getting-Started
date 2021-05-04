package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, functions!")
	startWebServer()
}

func startWebServer() {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started!")
}

/*
go run .

Hello, functions!
Starting server...
Server started!
*/
