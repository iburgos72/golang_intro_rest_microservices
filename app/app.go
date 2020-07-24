package app

import (
	"golang_intro_rest_microservices/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
