package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/usecase"
)

type AuthHandler struct {
    UserUsecase *usecase.UserUsecase
}

func NewAuthHandler(router *gin.Engine, us *usecase.UserUsecase) {
    handler := &AuthHandler{
        UserUsecase: us,
    }

    router.POST("/auth/register", handler.Register)
    router.POST("/auth/login", handler.Login)
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.UserUsecase.Register(&user); err != nil {
        if err.Error() == "user already exists" {
            c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully !"})
}
func (h *AuthHandler) Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.UserUsecase.Login(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
