package routes

import (
	"github.com/gorilla/mux"
	"github.com/sejal254/go-bookstore/pkg/controllers"
)

var RegisterBooksStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("POST")

}
