package models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Personalidade struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Personaliades []Personalidade

var CreateToken = func(u string, p string) string {

	var SampleSecret = []byte("secret")

	user := User{
		Username: u,
		Password: p,
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	}).SignedString(SampleSecret)
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
