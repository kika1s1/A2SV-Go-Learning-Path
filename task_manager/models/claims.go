package models

import "github.com/dgrijalva/jwt-go"

// Claims defines the structure of the JWT claims.
type Claims struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Role     string `bson:"role" json:"role"`
    jwt.StandardClaims
}