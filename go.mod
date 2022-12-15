module github.com/PaulTabaco/bookstore_items-api

go 1.19

require (
	github.com/PaulTabaco/bookstore_oauth v0.0.0-20221213140644-b465c0adf12c
	github.com/PaulTabaco/bookstore_utils v0.0.0-20221212224443-19484854a26a
)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mercadolibre/golang-restclient v0.0.0-20170701022150-51958130a0a0 // indirect
)

replace (
    github.com/PaulTabaco/bookstore_utils  => ../bookstore_utils // otherwise from $GOPATH - ../../../../pkg/mod/github.com/!paul!tabaco/bookstore_utils@v0.0.0-20221212224443-19484854a26a
    // github.com/PaulTabaco/bookstore_oauth => ../bookstore_oauth
)
