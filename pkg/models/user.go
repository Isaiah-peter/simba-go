package models

import (
	"simba-clone/pkg/config"
	utils "simba-clone/pkg/util"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	DollarAcount int
	EuroAccont   int
	PoundsAcount int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	hash := utils.HashPassword(u.Password)
	u.Password = hash
	db.Create(u)
	return u
}

func GetUserById(id int) (*User, *gorm.DB) {
	var user User
	db.Find(user, id)
	return &user, db
}
