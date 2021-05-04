package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	// fmt.Println(arr[4]) // ./main.go:16:17: invalid array index 4 (out of bounds for 3-element array)
}

/*
$ go run main.go
[1 2 3]
[1 2 3]
*/
