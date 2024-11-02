package services

import (
	"beego-poc/models"

	"github.com/beego/beego/v2/core/logs"
)

type UserService interface {
	GetByID(id int) *models.User
	Save(user *models.User)
}

type userService struct {
	storage map[int]*models.User
}

func newUserService() UserService {
	return &userService{
		storage: map[int]*models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
			2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
		},
	}
}

func (u *userService) GetByID(id int) *models.User {
	logs.Info("Get user by ID: id = %d", id)
	return u.storage[id]
}

func (u *userService) Save(user *models.User) {
	u.storage[user.ID] = user
}
