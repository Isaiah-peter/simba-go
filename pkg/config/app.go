package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dns := "host=localhost user=postgres password=etanuwoma dbname=simba port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
