package service

import (
	"gin-temp/internal/model"
	"sync"
)

type UserService struct{}

var (
	userService     *UserService
	onceUserService sync.Once
)

func GetUserService() *UserService {
	onceUserService.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (userService *UserService) GetUserByPhone(phone string) (*model.User, error) {
	userDao := model.GetUserDao()
	return userDao.GetByPhone(phone)
}
