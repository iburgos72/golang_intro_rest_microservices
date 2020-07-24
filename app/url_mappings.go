package app

import (
	"golang_intro_rest_microservices/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}