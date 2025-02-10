package taskService

type TaskService struct {
	repo TaskRepositoryInt
}

func NewTaskService(repo TaskRepositoryInt) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) PostTasks(task Task) (Task, error) {
	return s.repo.PostTasks(task)
}

func (s *TaskService) GetTasksID(id uint) (Task, error) { return s.repo.GetTasksId(id) }

func (s *TaskService) UpdateTasksID(id uint, newTask Task) (Task, error) {
	return s.repo.UpdateTasksId(id, newTask)
}

func (s *TaskService) DeleteTasksID(id uint) error {
	return s.repo.DeleteTasksId(id)
}
