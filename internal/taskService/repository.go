package taskService

import (
	"gorm.io/gorm"
	"pet/internal/database"
)

type TaskRepositoryInt interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type TaskRepositoryStr struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryStr {
	return &TaskRepositoryStr{db: db}
}

func (r *TaskRepositoryStr) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *TaskRepositoryStr) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryStr) UpdateTaskByID(id uint, newTask Task) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return Task{}, err
	}

	task.Task = newTask.Task
	task.IsDone = newTask.IsDone

	r.db.Save(&task)
	return task, err
}

func (r *TaskRepositoryStr) DeleteTaskByID(id uint) error {
	var task Task
	err := database.DB.First(&task, id).Error
	if err != nil {
		return err
	}
	r.db.Delete(&task)
	return nil
}
