package handlers

import (
	"context"
	"pet/internal/taskService"
	"pet/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.Id)

	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}

	return response, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body

	task, err := h.Service.GetTaskByID(request.Id)
	if err != nil {
		return nil, err
	}

	if taskRequest.Task != nil {
		task.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		task.IsDone = *taskRequest.IsDone
	}
	if taskRequest.Task == nil && taskRequest.IsDone == nil {
		response := tasks.PatchTasksId400Response{}
		return response, nil
	}
	newTask, err1 := h.Service.UpdateTaskByID(request.Id, task)
	if err1 != nil {
		return nil, err1
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &newTask.ID,
		Task:   &newTask.Task,
		IsDone: &newTask.IsDone,
	}

	return response, nil
}

func (h *Handler) GetTasksId(ctx context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.Service.GetTaskByID(request.Id)
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}
	return response, nil
}
