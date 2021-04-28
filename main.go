package main

import "fmt"

func main() {
	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}

/*
$ go run main.go
Arthur
0xc000010200 Arthur
0xc000010200 Tricia
*/
