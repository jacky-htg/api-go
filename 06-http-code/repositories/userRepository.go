package repositories

import (
	"github.com/jacky-htg/api-go/06-http-code/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/api-go/06-http-code/libraries"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := db.Query("SELECT `id`, `name`, `email`, `password` FROM `users`")

	libraries.CheckError(err)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var id int64
		var name string
		var email string
		var password string
		err = rows.Scan(&id, &name, &email, &password)
		libraries.CheckError(err)
		users = append(users, models.User{ID:id, Name:name, Email:email})
	}

	return users, err
}

func CreateUser(user models.User) (models.User, error)  {
	stmt, err := db.Prepare("INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, NOW())")
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	id, err := res.LastInsertId()
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	user.ID = id
	user.Password = nil

	return user, err
}

func GetUser(paramId int64) (models.User, error)  {
	rows, err := db.Query("SELECT id, `name`, `email`, password FROM users WHERE id=?", paramId)
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var id int64
	var name string
	var email string
	var password string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &password)
		libraries.CheckError(err)

		if err != nil {
			return models.User{}, err
		}
	}

	err = rows.Err()
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	return models.User{ID:id, Name:name, Email:email}, err
}

func GetPwdByEmail(inputEmail string)(string, error){
	rows, err := db.Query("SELECT `password` FROM users WHERE email=?", inputEmail)
	libraries.CheckError(err)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var databasePassword string

	for rows.Next() {
		err := rows.Scan(&databasePassword)
		libraries.CheckError(err)
		if err != nil {
			return "", err
		}
	}

	err = rows.Err()
	libraries.CheckError(err)
	if err != nil {
		return "", err
	}

	return databasePassword, err
}

func EditUser(user models.User) (models.User, error)  {
	stmt, err := db.Prepare("UPDATE `users` SET `name`=?, `email`=?, `password`=?, `updated_at`=NOW() WHERE `id`=?")
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.ID)
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	user.Password = nil

	return user, err
}

func DeleteUser(id int64) (bool, error)  {
	stmt, err := db.Prepare("DELETE FROM `users` WHERE `id`=?")
	libraries.CheckError(err)

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(id)
	libraries.CheckError(err)

	if err != nil {
		return false, err
	}

	return true, err
}