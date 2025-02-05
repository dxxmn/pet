package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pet/internal/database"
	"pet/internal/handlers"
	"pet/internal/taskService"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskByIDHandler).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskByIDHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
