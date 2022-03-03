package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount       int    `json:"amount"`
	CurrencyType string `json:"currency_type"`
	EmailTo      string `json:"email_to"`
	EmailFrom    string `json:"email_form"`
	Status       string `json:"status"`
}

func (u *Transaction) CreateTrancaction() *Transaction {
	db.Create(u)
	return u
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	user := new(User)
	db := db.Where("email = ?", email).Find(user)
	return user, db
}
