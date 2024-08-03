package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/controllers"
	"github.com/kika1s1/task_manager/middleware"
)

func SetupRouter()*gin.Engine {
	router :=gin.Default()
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	// Protected routes
	protected := router.Group("/tasks")
	protected.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		protected.GET("/tasks", controllers.GetAllTasks)
		protected.GET("/task/:id", controllers.GetTask)
		protected.PUT("/task/:id", controllers.UpdateTask)
		protected.DELETE("/task/:id", controllers.DeleteTask)
		protected.POST("/task", controllers.CreateTask)
	}
	return router
	
	

}