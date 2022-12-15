package services

import (
	"net/http"

	"github.com/PaulTabaco/bookstore_items-api/domain/items"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
)

var (
	// To make evailable to Outside of module
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, *rest_errors.NewRestError("should be implemented", http.StatusNotImplemented, "not implemented", nil)
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, *rest_errors.NewRestError("should be implemented", http.StatusNotImplemented, "not implemented", nil)
}
