package services

import (
	"github.com/gnshjoo/golang-microservices/mvc/domain"
	"github.com/gnshjoo/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// Get User From Databases
	return domain.GetUser(userId)
}
