package models

import (
	"fmt"
	"simba-clone/pkg/config"
	utils "simba-clone/pkg/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DollarAcount int    `json:"dollar_acc"`
	EuroAccount  int    `json:"euro_acc"`
	PoundsAcount int    `json:"pounds_acc"`
}

type Token struct {
	UserId int
	Email  string
	jwt.RegisteredClaims
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{}, &Transaction{})
}

func (u *User) CreateUser() *User {
	hash := utils.HashPassword(u.Password)
	u.Password = hash
	db.Create(u)
	return u
}

func GetUserById(id int) (*User, *gorm.DB) {
	var user User
	db.Find(&user, id)
	return &user, db
}
func GetUser() []User {
	var User []User
	db.Find(&User)
	return User
}

func FindOne(email string, password string) map[string]interface{} {
	var newUser *User
	if err := db.Where("email = ?", email).First(&newUser).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expireAt := time.Now().Add(time.Hour * 2)
	errf := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "password": newUser.Password, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	fmt.Println(errf)
	tk := &Token{
		UserId: int(newUser.ID),
		Email:  newUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expireAt,
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{
		"status":     true,
		"message":    "logged in",
		"token":      tokenString,
		"name":       newUser.Name,
		"email":      newUser.Email,
		"pounds_acc": newUser.PoundsAcount,
		"euro_acc":   newUser.EuroAccount,
		"dollar_acc": newUser.DollarAcount,
	}
	return resp
}
