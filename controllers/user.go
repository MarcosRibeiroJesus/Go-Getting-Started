package controllers

import "net/http"

type userController struct{}

func (uc userController) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
