package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/jacky-htg/api-go/05-token/repositories"
	"github.com/jacky-htg/api-go/05-token/libraries"
)

func GetTokenHandler(w http.ResponseWriter, req *http.Request) {
	inputEmail := req.FormValue("email")
	inputPassword := req.FormValue("password")

	if inputEmail == "" || inputPassword == "" {
		json.NewEncoder(w).Encode("Invalid email or password")
		return
	}

	databasePassword, err := repositories.GetPwdByEmail(inputEmail)
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputPassword))

	if err == nil {
		// Sign the token with our secret
		tokenString, err := libraries.ClaimToken(inputEmail)

		if err != nil {
			json.NewEncoder(w).Encode("Auth failed")
		}

		// Finally, write the token to the browser window
		json.NewEncoder(w).Encode(tokenString)
	} else {
		json.NewEncoder(w).Encode("finally Auth failed")
	}
}
