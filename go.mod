module github.com/PaulTabaco/bookstore_items-api

go 1.19

require (
	github.com/PaulTabaco/bookstore_oauth v0.0.0-20221224155501-ebd57bb1442a
	github.com/PaulTabaco/bookstore_utils v0.0.0-20221224154750-06edb834515f
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/elastic/go-elasticsearch/v8 v8.5.0
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
	github.com/mercadolibre/golang-restclient v0.0.0-20170701022150-51958130a0a0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

// replace (
// 	// github.com/PaulTabaco/bookstore_oauth => ../bookstore_oauth
// 	github.com/PaulTabaco/bookstore_utils => ../bookstore_utils // otherwise from $GOPATH - ../../../../pkg/mod/github.com/!paul!tabaco/bookstore_utils@v0.0.0-20221212224443-19484854a26a
// )
