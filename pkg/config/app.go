package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	host := os.Getenv("Host")
	user := os.Getenv("User")
	password := os.Getenv("Password")
	dbname := os.Getenv("Database")
	var dbport, _ = strconv.Atoi(os.Getenv("Port"))
	if host == "" {
		host = "localhost"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "etanuwoma"
	}
	if dbname == "" {
		dbname = "simba"
	}
	if dbport == 0 {
		dbport = 5500
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)
	d, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
