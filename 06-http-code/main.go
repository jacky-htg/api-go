package main

import (
	"net/http"
	"github.com/jacky-htg/api-go/06-redis/config"
	"github.com/jacky-htg/api-go/06-redis/routing"
)

func main() {
	router := routing.NewRouter()
	http.ListenAndServe(config.GetString("server.address"), router)
}