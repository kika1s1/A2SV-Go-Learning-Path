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
	r.GET("/tasks", authMiddleware, controllers.GetTasks)
	r.GET("/tasks/:id", authMiddleware, controllers.GetTaskByID)

	// Admin routes
	adminMiddleware := middleware.AdminMiddleware()
	r.PUT("/promote/:username",authMiddleware, adminMiddleware, controllers.Promote)
	r.DELETE("/tasks/:id", authMiddleware,adminMiddleware, controllers.DeleteTask)
	r.POST("/tasks", authMiddleware,adminMiddleware, controllers.CreateTask)
	r.PUT("/tasks/:id", authMiddleware,adminMiddleware, controllers.UpdateTask)
	return r
}