package app

import (
	"log"
	"net/http"
	"time"

	"github.com/PaulTabaco/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApp() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	log.Fatal(srv.ListenAndServe())
}
