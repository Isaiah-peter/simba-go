package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")
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
	if dbport == "" {
		dbport = "16020"
	}
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, dbport)
	fmt.Print(dns)
	d, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
