package main

func main() {
	println("Starting web server...")

	// do important things
	panic("Something bad just happened...")

	println("Web server started") // unreachable code
}

/*
go run main.go // panic

Starting web server...
panic: Something bad just happened...

goroutine 1 [running]:
main.main()
        /home/marcosribeiro/Documents/Learning/Go/Go-Getting-Started/main.go:7 +0x5c
exit status 2

go run main.go // without panic

Starting web server...
Web server started
*/
