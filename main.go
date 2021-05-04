package main

import (
	"net/http"

	"github.com/MarcosRibeiroJesus/Go-Getting-Started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

/*
$ go build .
$ ./Go-Getting-Started

http://localhost:3000/users

Hello from the User Controller!
*/
