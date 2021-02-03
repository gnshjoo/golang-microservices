package services

import (
	"github.com/gnshjoo/golang-microservices/mvc/domain"
	"github.com/gnshjoo/golang-microservices/mvc/utils"
	"net/http"
)

type itemService struct {}

var ItemService itemService

func GetItem(itemId string) (*domain.Item, *utils.ApplicationError)  {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}