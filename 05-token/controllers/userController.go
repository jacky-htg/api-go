package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-go/05-token/models"
	"github.com/jacky-htg/api-go/05-token/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/jacky-htg/api-go/05-token/libraries"
)

func GetUsersEndPoint(w http.ResponseWriter, req *http.Request) {
	if (len(req.Header["Token"]) == 0) {
		json.NewEncoder(w).Encode("Please suplay valid token")
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		json.NewEncoder(w).Encode("token tidak valid")
		return
	}

	json.NewEncoder(w).Encode(repositories.GetUsers())
}

func CreateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if (len(req.Header["Token"]) == 0) {
		json.NewEncoder(w).Encode("Please suplay valid token")
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		json.NewEncoder(w).Encode("token tidak valid")
		return
	}

	var user models.User

	user.Name   = req.FormValue("name")
	user.Email  = req.FormValue("email")
	password   := req.FormValue("password")

	// Validation
	if len(user.Name) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid name")
		return
	}

	if len(user.Email) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid email")
		return
	}

	if len(password) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid password")
		return
	}

	user.Password, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	json.NewEncoder(w).Encode(repositories.CreateUser(user))
}

func GetUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if (len(req.Header["Token"]) == 0) {
		json.NewEncoder(w).Encode("Please suplay valid token")
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		json.NewEncoder(w).Encode("token tidak valid")
		return
	}

	params := mux.Vars(req)
	id, _  := strconv.ParseInt(params["id"], 10, 64)

	json.NewEncoder(w).Encode(repositories.GetUser(id))
}