package app

import (
	"net/http"

	"github.com/PaulTabaco/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/delete/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
}
