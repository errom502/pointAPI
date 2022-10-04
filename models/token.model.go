package models

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
