package router

import (
    "github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/controllers"

    
)

func SetRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.CreateTask)
	

}