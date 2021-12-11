package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GetDB() (db *sql.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file!")
	}

	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	mySqlUrl := dbUser + ":" + dbPass + "@(127.0.0.1:3306" + ")/" + dbName + "?parseTime=true"

	db, err = sql.Open(dbDriver, mySqlUrl)

	return
}
