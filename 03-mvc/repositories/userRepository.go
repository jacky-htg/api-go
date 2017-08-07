package repositories

import (
	"time"
	"github.com/jacky-htg/api-go/03-mvc/models"
)

var users []models.User

func GetUsers()([]models.User) {
	return users
}

func CreateUser(user models.User)(models.User) {
	user.ID         = int64(len(users)+1)
	user.CreatedAt  = time.Now()

	users = append(users, user)
	return user
}

func GetUser(id int64)(models.User) {
	for _, v := range users {
		if v.ID == id {
			return v
		}
	}

	return models.User{}
}