package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"yonje/baseframework/config"
)

type User struct {
	Id       int64  `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
	Name     string `json:"name" form:"name"`
}

func (self *User) GenerateJwtToken() string {
	tokenStr, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.FormatInt(self.Id, 10),
	}).SignedString(config.SigningSecret)

	return tokenStr
}

func GeneratePasswordHash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed[:])
}

func CheckHashWithPassWord(hashed string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
