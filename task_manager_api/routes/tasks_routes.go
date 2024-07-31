package routes

import (
    "github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager_api/controllers"

    
)

func Routes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)

}