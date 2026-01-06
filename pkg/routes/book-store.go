package routes

import (
	"github.com/gorilla/mux"
	"//pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandlerFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

//purpose of this file is to register all the routes related to book store
