package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dbname := os.Getenv("db_name")
	if dbname == "" {
		dbname = "simba"
	}

	dbuser := os.Getenv("db_username")
	if dbuser == "" {
		dbuser = "isaiah"
	}
	dbpass := os.Getenv("db_")
	if dbpass == "" {
		dbpass = "Etanuwoma18"
	}
	dns := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpass, dbname)
	d, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
