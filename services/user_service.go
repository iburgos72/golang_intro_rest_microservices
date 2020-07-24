package services

import (
	"golang_intro_rest_microservices/domain"
	"golang_intro_rest_microservices/utils"
)

type usersService struct {}

var(
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError)  {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}