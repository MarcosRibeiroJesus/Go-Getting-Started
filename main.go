package main

import "fmt"

func main() {
	const pi = 3.1415
	fmt.Println(pi)

	const c int = 3
	fmt.Println(c + 3)

	// fmt.Println(c + 1.2)  1.2 (untyped float constant) truncated to int
	fmt.Println(float32(c) + 1.2)
}

/*
$ go run main.go
3.1415
6
4.2
*/
