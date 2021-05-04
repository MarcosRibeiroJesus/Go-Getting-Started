package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	slice = append(slice, 4, 27, 12)

	fmt.Println(slice)

	s2 := slice[1:] // slice the variable starting at item index 1 and proceeding to the end.
	s3 := slice[:2] // set an ending index (not incluiding the index 2, just 0 and 1)
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)
}

/*
$ go run main.go
[1 2 3]
[1 2 3 4 27 12]
[2 3 4 27 12] [1 2] [2]
*/
