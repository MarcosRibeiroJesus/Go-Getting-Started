package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, functions!")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port, numbersOfRetries int) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Numbers of retries", numbersOfRetries)
}

/*
go run .

Hello, functions!
Starting server...
Server started!
*/
