package services

import (
	"github.com/PaulTabaco/bookstore_items-api/domain/items"
	"github.com/PaulTabaco/bookstore_items-api/domain/queries"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
)

var (
	// To make evailable to Outside of module
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Update(string, []byte) (*string, rest_errors.RestErr)
	Delete(string) (*string, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemsService) Update(id string, request []byte) (*string, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Update(request); err != nil {
		return nil, err
	}
	return &item.Id, nil
}

func (s *itemsService) Delete(id string) (*string, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Delete(); err != nil {
		return nil, err
	}
	return &item.Id, nil
}
