package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/models"
)

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
		
        // Type assert the user to the claims type
        claims, ok := user.(*models.Claims)
        if !ok || claims.Role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
        }

        c.Next()
    }
}
