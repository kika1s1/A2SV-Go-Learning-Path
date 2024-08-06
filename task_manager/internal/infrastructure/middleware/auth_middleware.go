package middleware

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecret), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(*jwt.StandardClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            c.Abort()
            return
        }

        c.Set("userID", claims.Subject)
        c.Next()
    }
}
