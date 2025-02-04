package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	fmt.Fprint(w, "Успешно добавлено\n")
	json.NewEncoder(w).Encode(task)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	fmt.Fprint(w, "Все задания\n")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	initDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTasksHandler).Methods("POST")
	router.HandleFunc("/api/tasks", GetTasksHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
