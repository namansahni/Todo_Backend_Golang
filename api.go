package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	db := dbConfig()

	var todos []Todo
	db.Find(&todos)

	fmt.Println(todos)
	fmt.Println(reflect.TypeOf(todos))
	json.NewEncoder(w).Encode(todos)
}
