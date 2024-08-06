package entities

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Permission int

const (
	NONE  Permission = 0
	READ  Permission = 1
	EDIT  Permission = 2
	ADMIN Permission = 3
)

type Translate struct {
	Raw       string `json:"raw"`
	Translate string `json:"translate"`
}
