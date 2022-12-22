package main

import (
	"os"

	"github.com/PaulTabaco/bookstore_items-api/app"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApp()
}
