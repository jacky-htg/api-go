package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jacky-htg/api-go/02-config-file/config"
	"github.com/jacky-htg/api-go/03-mvc/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsersEndPoint).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetUserEndPoint).Methods("GET")

	http.ListenAndServe(config.GetString("server.address"), router)
}