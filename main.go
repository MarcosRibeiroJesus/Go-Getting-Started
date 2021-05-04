package main

func main() {
	// wellKnownPorts := map[string]int{"http": 80, "https": 443}
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
}

/*
$ go run main.go with map

https 443
http 80

$ go run main.go with slice

0 1
1 2
2 3
*/
