package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
	ExpiresAt int64  `json:"expiresAt"`
}
