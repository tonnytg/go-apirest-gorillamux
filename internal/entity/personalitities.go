package entity

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Personality struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Personalities []Personality

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
