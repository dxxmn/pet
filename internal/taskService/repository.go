package taskService

import (
	"gorm.io/gorm"
	"pet/internal/database"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, newTask Task) (Task, error) {
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

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	err := database.DB.First(&task, id).Error
	if err != nil {
		return err
	}
	r.db.Delete(&task)
	return nil
}
