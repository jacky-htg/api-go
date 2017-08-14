package main

import (
	"net/http"
	"github.com/jacky-htg/api-go/06-http-code/config"
	"github.com/jacky-htg/api-go/06-http-code/routing"
)

func main() {
	router := routing.NewRouter()
	http.ListenAndServe(config.GetString("server.address"), router)
}