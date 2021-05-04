package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Marcos"
	u.LastName = "Ribeiro"
	fmt.Println(u.FirstName + " " + u.LastName)

	u2 := user{
		ID:        1,
		FirstName: "Marcos",
		LastName:  "Ribeiro",
	}
	fmt.Println(u2)
}

/*
$ go run main.go

Marcos Ribeiro
{1 Marcos Ribeiro}
*/
