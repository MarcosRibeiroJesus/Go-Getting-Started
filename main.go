package main

import "fmt"

const (
	first  = iota + 6  // first = 6 && iota = 0 prints 6
	second = 2 << iota // second = 2 << 2 || 2 * iota(1) * 2 = 4 && iota = 1 + 6(first) prints 4
	third  = 3 << iota // third 3 << 3 ??? prints 12
	fourth             // ?????? prints 24
)

const (
	fifth = iota
	sixth
)

func main() {
	fmt.Println(first, second, third, fourth, fifth, sixth)
}

/*
$ go run main.go
6 4 12 24 0 1
*/
