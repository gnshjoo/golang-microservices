package services

import (
	"github.com/gnshjoo/golang-microservices/mvc/utils"
	"github.com/gnshjoo/golang-microservices/mvc/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User,*utils.ApplicationError)
)

type usersDaoMock struct {}

func init() {
	domain.UserDao = &userDaoMock
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User,*utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User,*utils.ApplicationError){
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 was not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNotError(t *testing.T) {
	getUserFunction = func(userId int64)(*domain.User,*utils.ApplicationError){
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}