package services

import (
	"golang_intro_rest_microservices/domain"
	"golang_intro_rest_microservices/utils"
	"net/http"
)

type itemsServices struct {}

var (
	ItemsService itemsServices
)

func (i *itemsServices) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}