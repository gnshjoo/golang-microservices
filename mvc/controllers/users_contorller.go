package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gnshjoo/golang-microservices/mvc/services"
	"github.com/gnshjoo/golang-microservices/mvc/utils"
	_ "log"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		// Just return the Bad Request to the client
		//c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		// Handle the err adn return to the client
		//c.JSON(apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)
		return
	}
	// return user to client
	//c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}
