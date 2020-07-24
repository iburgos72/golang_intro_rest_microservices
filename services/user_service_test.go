package services

import (
	"github.com/stretchr/testify/assert"
	"golang_intro_rest_microservices/domain"
	"golang_intro_rest_microservices/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T)  {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 does not exist",
		}
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T)  {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 125,
		}, nil
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 125, user.Id)
}