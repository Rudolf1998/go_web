package main

import (
	"net/http"
	"main/controller"
)

func main(){
	controller.RegisterRoutes()
	http.ListenAndServe(":8282", nil)
}
