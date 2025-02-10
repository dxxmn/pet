package userService

import (
	"gorm.io/gorm"
)

type UserRepositoryInt interface {
	GetUsers() ([]User, error)
	PostUsers(user User) (User, error)
	GetUsersId(id uint) (User, error)
	UpdateUsersId(id uint, newUser User) (User, error)
	DeleteUsersId(id uint) error
}

type UserRepositoryStr struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryStr { return &UserRepositoryStr{db: db} }

func (r *UserRepositoryStr) GetUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepositoryStr) PostUsers(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepositoryStr) GetUsersId(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepositoryStr) UpdateUsersId(id uint, newUser User) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	user.Email = newUser.Email
	user.Password = newUser.Password
	err1 := r.db.Save(&user).Error
	if err1 != nil {
		return User{}, err1
	}
	return user, nil
}

func (r *UserRepositoryStr) DeleteUsersId(id uint) error {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}
	err1 := r.db.Delete(&user).Error
	if err1 != nil {
		return err1
	}
	return nil
}
