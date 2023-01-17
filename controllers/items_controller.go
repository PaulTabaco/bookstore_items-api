package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PaulTabaco/bookstore_items-api/domain/items"
	"github.com/PaulTabaco/bookstore_items-api/services"
	"github.com/PaulTabaco/bookstore_items-api/utils/http_utils"
	"github.com/PaulTabaco/bookstore_oauth/oauth"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
	"github.com/gorilla/mux"
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
		// http_utils.RespondError(w, err)
		return
	}
	sellerID := oauth.GetCallerId(r)

	if sellerID == 0 {
		http_utils.RespondError(w, rest_errors.NewUnauthorizedError("unauthorized"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http_utils.RespondError(w, rest_errors.NewBadRequestError("invalid request body"))
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		http_utils.RespondError(w, rest_errors.NewBadRequestError("invalid item json body"))
		return
	}

	itemRequest.Seller = sellerID

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
