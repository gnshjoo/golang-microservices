package services

import (
	"github.com/gnshjoo/golang-microservices/mvc/domain"
	"github.com/gnshjoo/golang-microservices/mvc/utils"

)

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// Get User From Databases
	//return domain.UserDao.GetUser(userId)
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil

}
