package controllers

import (
	"fmt"
	"net/http"

	"github.com/PaulTabaco/bookstore_items-api/domain/items"
	"github.com/PaulTabaco/bookstore_items-api/services"
	"github.com/PaulTabaco/bookstore_oauth/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: - return error to user (by w)
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: - return error json to user (by w)
		return
	}

	fmt.Println(result)
	// TODO: - return result item json to w , and HTTP status 201
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
