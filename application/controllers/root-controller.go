package controllers

import (
	"github.com/gorilla/mux"

)



func NewRootController() *mux.Router  {
	postController := NewPostController()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",postController.GetAll).Methods("GET")
	router.HandleFunc("/{id}", postController.GetById).Methods("GET")
	router.HandleFunc("/",postController.Create).Methods("POST")
	router.HandleFunc("/{id}",postController.Delete).Methods("DELETE")

	return router
}