package userService

import "pet/internal/taskService"

type UserService struct {
	repo UserRepositoryInt
}

func NewUserService(repo UserRepositoryInt) *UserService { return &UserService{repo: repo} }

func (s *UserService) GetUsers() ([]User, error) { return s.repo.GetUsers() }

func (s *UserService) PostUsers(user User) (User, error) { return s.repo.PostUsers(user) }

func (s *UserService) GetUsersId(id uint) (User, error) { return s.repo.GetUsersId(id) }

func (s *UserService) UpdateUsersId(id uint, newUser User) (User, error) {
	return s.repo.UpdateUsersId(id, newUser)
}

func (s *UserService) DeleteUsersId(id uint) error { return s.repo.DeleteUsersId(id) }

func (s *UserService) GetTasksForUser(id uint) ([]taskService.Task, error) {
	return s.repo.GetTasksForUser(id)
}
