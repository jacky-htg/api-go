package repositories

import (
	"database/sql"
	"github.com/jacky-htg/api-go/05-token/config"
	"github.com/jacky-htg/api-go/05-token/libraries"
)

var db *sql.DB
var err error

func init() {
	// Create an sql.DB and check for errors
	db, err = sql.Open(config.GetString("database.driverName"), config.GetString("database.dataSourceName"))
	libraries.CheckError(err)

	// Test the connection to the database
	err = db.Ping()
	libraries.CheckError(err)
}

