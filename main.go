package main

func main() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			// break
			continue
		}
		println("continuing...")
	}
}

/*
$ go run main.go // break
0
continuing...
1
continuing...
2

$ go run main.go // continue

0
continuing...
1
continuing...
2
3
continuing...
4
continuing...
*/
