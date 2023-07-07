package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sejal254/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
