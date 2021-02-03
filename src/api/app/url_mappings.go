package app

import (
	"github.com/gnshjoo/golang-microservices/src/api/contollers/polo"
	"github.com/gnshjoo/golang-microservices/src/api/contollers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}