package controllers

import (
	"encoding/json"
	"github.com/gnshjoo/golang-microservices/mvc/services"
	"github.com/gnshjoo/golang-microservices/mvc/utils"
	_ "log"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request)  {
	userIdParam := request.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		// Just return the Bad Request to the client
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the err adn return to the client
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
