package main

import (
	"github.com/gorilla/mux"
)

func HandleRoot() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/getAllTodos", getAllTodos).Methods("GET", "OPTIONS")

	return r
}
