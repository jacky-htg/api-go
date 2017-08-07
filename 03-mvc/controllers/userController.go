package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-go/03-mvc/models"
	"github.com/jacky-htg/api-go/03-mvc/repositories"
)

func GetUsersEndPoint(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(repositories.GetUsers())
}

func CreateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	var user models.User

	user.Name   = req.FormValue("name")
	user.Email  = req.FormValue("email")

	// Validation
	if len(user.Name) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid name")
		return
	}

	if len(user.Email) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid email")
		return
	}

	json.NewEncoder(w).Encode(repositories.CreateUser(user))
}

func GetUserEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _  := strconv.ParseInt(params["id"], 10, 64)

	json.NewEncoder(w).Encode(repositories.GetUser(id))
}