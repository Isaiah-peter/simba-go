package models

import (
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
	Name         string
	Email        string
	Password     string
	DollarAcount int
	EuroAccount  int
	PoundsAcount int
}

type LoginResponse struct {
	userId       int
	name         string
	email        string
	dollarAcount int
	euroAccount  int
	poundsAcount int
	accessToken  string
	errorMessage string
	status       bool
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
	db.Find(user, id)
	return &user, db
}

func FindUser(email, password string) *LoginResponse {
	var newUser *User
	if err := db.Find(&newUser, "email = ?", email); err != nil {
		return &LoginResponse{
			errorMessage: "unkown email",
			status:       false,
		}
	}

	expireAt := time.Now().Add(time.Hour * 2)

	errp := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if errp != nil && errp == bcrypt.ErrMismatchedHashAndPassword {
		return &LoginResponse{
			errorMessage: "password unkown or mismarch",
			status:       false,
		}
	}

	tk := &Token{
		UserId: int(newUser.ID),
		Email:  newUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expireAt,
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, tk)
	tokenString, err := token.SignedString([]byte("qwertyuiopoiuytr"))
	if err != nil {
		panic(err)
	}
	result := &LoginResponse{
		userId:       int(newUser.ID),
		name:         newUser.Name,
		email:        newUser.Email,
		dollarAcount: newUser.DollarAcount,
		poundsAcount: newUser.PoundsAcount,
		euroAccount:  newUser.EuroAccount,
		accessToken:  tokenString,
		errorMessage: "login successful",
		status:       true,
	}

	return result
}
