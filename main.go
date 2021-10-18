package main

import (
	"log"
	"net/http"
	"takeout-backend/application/controllers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", controllers.NewRootController()))
}
