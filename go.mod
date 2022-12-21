module github.com/PaulTabaco/bookstore_items-api

go 1.19

require (
	github.com/PaulTabaco/bookstore_oauth v0.0.0-20221216155034-d5bcb1d6f7aa
	github.com/PaulTabaco/bookstore_utils v0.0.0-20221216140554-1772c03f800c
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.0.0-20211216131617-bbee439d559c // indirect
	github.com/elastic/go-elasticsearch v0.0.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.5.0 // indirect
	github.com/mercadolibre/golang-restclient v0.0.0-20170701022150-51958130a0a0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

// replace (
// 	github.com/PaulTabaco/bookstore_oauth => ../bookstore_oauth
// 	github.com/PaulTabaco/bookstore_utils => ../bookstore_utils // otherwise from $GOPATH - ../../../../pkg/mod/github.com/!paul!tabaco/bookstore_utils@v0.0.0-20221212224443-19484854a26a
// )
