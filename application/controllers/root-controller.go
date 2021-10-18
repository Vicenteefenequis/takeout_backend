package controllers

import (
	"github.com/gorilla/mux"
)

func NewRootController() *mux.Router {
	postController := NewPostController()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/post", postController.GetAll).Methods("GET")
	router.HandleFunc("/post", postController.Create).Methods("POST")
	router.HandleFunc("/post/{id}", postController.GetById).Methods("GET")
	router.HandleFunc("/post/{id}", postController.Delete).Methods("DELETE")
	router.HandleFunc("/post/{id}", postController.Update).Methods("PUT")
	return router
}
