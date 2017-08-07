package main

import (
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/jacky-htg/api-go/02-config-file/config"
)

type User struct {
	ID	        int64
	Name        string
	Email       string
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeletedAt   time.Time
}

var users []User

func main() {
    router := mux.NewRouter()
    users = append(users, User{ ID: 1, Name: "Jacky Chan"})
    users = append(users, User{ ID: 2, Name: "Jet Lee", Email: "jet@lee.com"})
    
    router.HandleFunc("/api/users", GetUsersEndPoint).Methods("GET")
    router.HandleFunc("/api/users", CreateUserEndPoint).Methods("POST")
    router.HandleFunc("/api/users/{id}", GetUserEndPoint).Methods("GET")

    http.ListenAndServe(config.GetString("server.address"), router)
}

func GetUsersEndPoint(w http.ResponseWriter, _ *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func CreateUserEndPoint(w http.ResponseWriter, req *http.Request) {
    var user User
    
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
     
    user.ID         = int64(len(users)+1)
    user.CreatedAt  = time.Now()
    
    users = append(users, user)    
    json.NewEncoder(w).Encode(user)
}

func GetUserEndPoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id, _  := strconv.ParseInt(params["id"], 10, 64)
    
    for _, v := range users {
        if v.ID == id {
            json.NewEncoder(w).Encode(v)
            return
        }
    }      
    
    json.NewEncoder(w).Encode(&User{})
}
