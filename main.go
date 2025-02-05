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
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task Task
	err := DB.First(&task, id).Error
	if err != nil {
		fmt.Fprint(w, "Задача не найдена")
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	DB.Save(&task)
	fmt.Fprint(w, "Задача обновлена\n")
	json.NewEncoder(w).Encode(task)
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task Task
	err := DB.First(&task, id).Error
	if err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTasksHandler).Methods("POST")
	router.HandleFunc("/api/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", UpdateTasksHandler).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", DeleteTasksHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
