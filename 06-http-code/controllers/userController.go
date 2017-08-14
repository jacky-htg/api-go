package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-go/06-http-code/models"
	"github.com/jacky-htg/api-go/06-http-code/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/jacky-htg/api-go/06-http-code/libraries"
)

var err error

func GetUsersEndPoint(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Token"]) == 0 {
		http.Error(w, "Please suplay valid token", http.StatusBadRequest)
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		return
	}

	users, err := repositories.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func CreateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Token"]) == 0 {
		http.Error(w, "Please suplay valid token", http.StatusBadRequest)
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		return
	}

	var user models.User

	user.Name   = req.FormValue("name")
	user.Email  = req.FormValue("email")
	password   := req.FormValue("password")

	// Validation
	if len(user.Name) == 0 {
		http.Error(w, "Please suplay valid name", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Please suplay valid email", http.StatusBadRequest)
		return
	}

	if len(password) == 0 {
		http.Error(w, "Please suplay valid password", http.StatusBadRequest)
		return
	}

	user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := repositories.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Token"]) == 0 {
		http.Error(w, "Please suplay valid token", http.StatusBadRequest)
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(req)
	id, err  := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := repositories.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID <= 0 {
		http.Error(w, "User not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func EditUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Token"]) == 0 {
		http.Error(w, "Please suplay valid token", http.StatusBadRequest)
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(req)
	id, err  := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := repositories.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Name   = req.FormValue("name")
	user.Email  = req.FormValue("email")
	password   := req.FormValue("password")

	// Validation
	if len(user.Name) == 0 {
		http.Error(w, "Please suplay valid name", http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Please suplay valid email", http.StatusBadRequest)
		return
	}

	if len(password) == 0 {
		http.Error(w, "Please suplay valid password", http.StatusBadRequest)
		return
	}

	user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err = repositories.EditUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUserEndPoint(w http.ResponseWriter, req *http.Request) {
	if len(req.Header["Token"]) == 0 {
		http.Error(w, "Please suplay valid token", http.StatusBadRequest)
		return
	}

	isTokenValid, _ := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(req)
	id, err  := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ , err = repositories.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("User has been deleted")
}