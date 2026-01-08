package routes

import (
	"Go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router){
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
}
