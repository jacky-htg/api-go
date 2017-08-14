package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/jacky-htg/api-go/06-http-code/repositories"
	"github.com/jacky-htg/api-go/06-http-code/libraries"
)

func GetTokenHandler(w http.ResponseWriter, req *http.Request) {
	inputEmail := req.FormValue("email")
	inputPassword := req.FormValue("password")

	if inputEmail == "" || inputPassword == "" {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	databasePassword, err := repositories.GetPwdByEmail(inputEmail)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(databasePassword) == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputPassword))

	if err != nil {
		libraries.CheckError(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}


	// Sign the token with our secret
	tokenString, err := libraries.ClaimToken(inputEmail)

	if err != nil {
		http.Error(w, "Auth failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tokenString)
}
