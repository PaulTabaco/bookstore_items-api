package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApp() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 45 * time.Second,
		ReadTimeout:  45 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	// log.Fatal(srv.ListenAndServe())
}
