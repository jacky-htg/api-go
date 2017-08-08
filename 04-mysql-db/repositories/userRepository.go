package repositories

import (
	"github.com/jacky-htg/api-go/04-mysql-db/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/api-go/04-mysql-db/libraries"
)

func GetUsers() ([]models.User) {
	rows, err := db.Query("SELECT `id`, `name`, `email`, `password` FROM users")

	libraries.CheckError(err)

	var users []models.User

	for rows.Next() {
		var id int64
		var name string
		var email string
		var password string
		err = rows.Scan(&id, &name, &email, &password)
		libraries.CheckError(err)
		users = append(users, models.User{ID:id, Name:name, Email:email})
	}

	return users
}

func CreateUser(user models.User) (models.User)  {
	stmt, err := db.Prepare("INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, NOW())")
	libraries.CheckError(err)

	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	libraries.CheckError(err)

	id, err := res.LastInsertId()
	libraries.CheckError(err)

	user.ID = id
	user.Password = nil

	return user
}

func GetUser(paramId int64) (models.User)  {
	rows, err := db.Query("SELECT id, name, email, password FROM users WHERE id=?", paramId)
	libraries.CheckError(err)
	defer rows.Close()

	var id int64
	var name string
	var email string
	var password string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &password)
		libraries.CheckError(err)
	}

	err = rows.Err()
	libraries.CheckError(err)

	return models.User{ID:id, Name:name, Email:email}
}