package main

import (
	"fmt"

	"github.com/MarcosRibeiroJesus/Go-Getting-Started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Luan",
		LastName:  "Samuel",
	}
	fmt.Println(u)
}

/*
go run github.com/MarcosRibeiroJesus/Go-Getting-Started/ or go run .

{2 Luan Samuel}
*/
