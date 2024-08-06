package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	IsAdmin     bool `bson:"isAdmin" json:"isAdmin"`
}

type Claims struct {
	Username string `json:"username"`
	IsAdmin     bool `json:"isAdmin"`
	jwt.StandardClaims
	ExpiresAt int64  `json:"expiresAt"`
}
