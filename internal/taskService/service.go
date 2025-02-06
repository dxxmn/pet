package taskService

type TaskService struct {
	repo TaskRepositoryInt
}

func NewService(repo TaskRepositoryInt) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, newTask Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, newTask)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
