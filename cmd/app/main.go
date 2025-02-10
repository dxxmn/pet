package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet/internal/database"
	"pet/internal/handlers"
	"pet/internal/taskService"
	"pet/internal/userService"
	"pet/internal/web/tasks"
	"pet/internal/web/users"
)

func main() {
	database.InitDB()

	TaskRepo := taskService.NewTaskRepository(database.DB)
	TaskService := taskService.NewTaskService(TaskRepo)
	TaskHandler := handlers.NewTaskHandler(TaskService)

	UserRepo := userService.NewUserRepository(database.DB)
	UserService := userService.NewUserService(UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(TaskHandler, nil)
	strictUserHandler := users.NewStrictHandler(UserHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
