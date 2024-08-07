package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kika1s1/task_manager/data"
	"github.com/kika1s1/task_manager/models"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Check if username already exists
    existingUser, err := data.GetUserByUsername(user.Username)
    if err == nil && existingUser != (models.User{}) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    // Check password hardness
    if err := checkPasswordHardness(user.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if there are any users in the database
    userCount, err := data.GetUserCount()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check user count"})
        return
    }

    // If no users exist, set the new user as an admin
    if userCount == 0 {
        user.IsAdmin = true
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }
    user.Password = string(hashedPassword)

    err = data.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":  "User created",
        "username": user.Username,
        "isAdmin":  user.IsAdmin,
    })
}

// checkPasswordHardness checks if the password meets the required hardness criteria using regex
func checkPasswordHardness(password string) error {
    var (
        minLen    = regexp.MustCompile(`.{8,}`)
        upperCase = regexp.MustCompile(`[A-Z]`)
        lowerCase = regexp.MustCompile(`[a-z]`)
        number    = regexp.MustCompile(`[0-9]`)
        special   = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",.<>?\\|/~` + "`" + `]`)
    )

    if !minLen.MatchString(password) {
        return fmt.Errorf("password must be at least 8 characters long")
    }
    if !upperCase.MatchString(password) {
        return fmt.Errorf("password must contain at least one uppercase letter")
    }
    if !lowerCase.MatchString(password) {
        return fmt.Errorf("password must contain at least one lowercase letter")
    }
    if !number.MatchString(password) {
        return fmt.Errorf("password must contain at least one number")
    }
    if !special.MatchString(password) {
        return fmt.Errorf("password must contain at least one special character")
    }

    return nil
}

// Login handles user login and returns JWT token
func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	storedUser, err := data.GetUserByUsername(user.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		fmt.Println()
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateJWT(storedUser.Username, storedUser.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GenerateJWT creates a new JWT token
func GenerateJWT(username string, isAdmin bool) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	JWT_SECRET := os.Getenv("JWT_SECRET")
	claims := models.Claims{
		Username:  username,
		IsAdmin:      isAdmin,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}
// Promote admin role
func Promote(c *gin.Context) {
	username := c.Param("username")
	if err := data.Promote(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not promote user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User is promoted Admin successfully",
	})
}
// CreateTask handles task creation
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := data.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created",})
}

// GetTasks handles fetching all tasks
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles fetching a single task by ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask handles updating a task
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := data.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTask handles deleting a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task is deleted successfully",
	})
}
