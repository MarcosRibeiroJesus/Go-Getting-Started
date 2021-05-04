package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	// println(i) // ./main.go:7:10: undefined: i
}

/*
$ go run main.go

0
1
2
3
4
*/
