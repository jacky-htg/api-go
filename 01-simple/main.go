package main

import (
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"encoding/json"
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
    
    router.HandleFunc("/api/users", GetUserEndPoint).Methods("GET")
    
    http.ListenAndServe(":9090", router)
}

func GetUserEndPoint(w http.ResponseWriter, _ *http.Request) {
    json.NewEncoder(w).Encode(users)
}
