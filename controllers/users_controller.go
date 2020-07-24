package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_intro_rest_microservices/services"
	"golang_intro_rest_microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}