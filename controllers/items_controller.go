package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PaulTabaco/bookstore_items-api/domain/items"
	"github.com/PaulTabaco/bookstore_items-api/services"
	"github.com/PaulTabaco/bookstore_items-api/utils/http_utils"
	"github.com/PaulTabaco/bookstore_oauth/oauth"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
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
		http_utils.RespondError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http_utils.RespondError(w, rest_errors.NewBadRequestError("invalid reques body"))
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		http_utils.RespondError(w, rest_errors.NewBadRequestError("invalid item json body"))
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
