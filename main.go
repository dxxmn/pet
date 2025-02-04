package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type TaskRequest struct {
	Task string `json:"task"`
}

var task string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, "+task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var t TaskRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}
	task = t.Task
	fmt.Fprintln(w, "Task written: "+task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
