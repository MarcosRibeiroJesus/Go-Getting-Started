package main

func main() {
	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

/*
$ go run main.go

0
1
2
3
4
*/
