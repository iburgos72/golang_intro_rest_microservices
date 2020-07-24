package domain

import (
	"fmt"
	"golang_intro_rest_microservices/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		123: {Id: 1, FirstName: "Ivan", LastName: "Burgos", Email: "ivan.burgos@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64)(*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("user %v does not exist", userId),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}
	return user, nil
}
