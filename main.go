package main

import "fmt"

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 33
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}

/*
$ go run main.go

map[foo:42]
42
map[foo:33]
map[]
*/
