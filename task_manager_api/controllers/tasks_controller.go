package controllers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager_api/data"

)

func GetAllTasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, data.Tasks)
}


func GetTask(c *gin.Context) {
    id := c.Param("id")
    for _, task := range data.Tasks {
        if task.ID == id {
            c.IndentedJSON(http.StatusOK, task)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}