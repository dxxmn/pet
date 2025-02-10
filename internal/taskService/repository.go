package taskService

import (
	"gorm.io/gorm"
)

type TaskRepositoryInt interface {
	GetTasks() ([]Task, error)
	PostTasks(task Task) (Task, error)
	GetTasksId(id uint) (Task, error)
	UpdateTasksId(id uint, task Task) (Task, error)
	DeleteTasksId(id uint) error
}

type TaskRepositoryStr struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryStr {
	return &TaskRepositoryStr{db: db}
}

func (r *TaskRepositoryStr) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryStr) PostTasks(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryStr) GetTasksId(id uint) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *TaskRepositoryStr) UpdateTasksId(id uint, newTask Task) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return Task{}, err
	}

	task.Task = newTask.Task
	task.IsDone = newTask.IsDone

	err1 := r.db.Save(&task).Error
	if err1 != nil {
		return Task{}, err1
	}
	return task, nil
}

func (r *TaskRepositoryStr) DeleteTasksId(id uint) error {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}
	err1 := r.db.Delete(&task).Error
	if err1 != nil {
		return err1
	}
	return nil
}
