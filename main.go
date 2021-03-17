package main

import (
	"log"
	"net/http"
)

func main() {
	dbConfig()
	handleRequests()
}

func handleRequests() {
	r := HandleRoot()
	log.Fatal(http.ListenAndServe(":8080", r))
}
