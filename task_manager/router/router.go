package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/controllers"
	"github.com/kika1s1/task_manager/data"
	"github.com/kika1s1/task_manager/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)
func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Initialize collections
	data.InitTaskCollection(client)
	data.InitUserCollection(client)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
            "message": "Welcome to the Task Manager API",
        })
	})
	// Public routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// Auth middleware
	authMiddleware := middleware.AuthMiddleware()

	// Task routes
	r.POST("/tasks", authMiddleware, controllers.CreateTask)
	r.GET("/tasks", authMiddleware, controllers.GetTasks)
	r.PUT("/tasks/:id", authMiddleware, controllers.UpdateTask)
	r.GET("/tasks/:id", authMiddleware, controllers.GetTaskByID)
	r.DELETE("/tasks/:id", authMiddleware, controllers.DeleteTask)

	// Admin routes
	adminMiddleware := middleware.AdminMiddleware()
	r.DELETE("/tasks/:id/admin",authMiddleware, adminMiddleware, controllers.DeleteTask)

	return r
}