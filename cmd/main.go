package main

import (
    "Go-bookstore/pkg/config"
    "Go-bookstore/pkg/utils"
    "Go-bookstore/pkg/routes"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
	//Initalize Redis
	utils.InitRedis()
    // Initialize MySQL
    config.GetDB()

    // Initialize Redis
    utils.InitRedis()

    // Router
    router := mux.NewRouter()
    routes.RegisterBookStoreRoutes(router)

	log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)
}
