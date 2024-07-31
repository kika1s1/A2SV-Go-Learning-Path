package controllers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager_api/data"
	"github.com/kika1s1/task_manager_api/models"

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


 func UpdateTask(ctx *gin.Context) {
    id := ctx.Param("id")

    var updatedTask models.Task

    if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, task := range data.Tasks {
        if task.ID == id {
            // Update only the specified fields
            if updatedTask.Title != "" {
                data.Tasks[i].Title = updatedTask.Title
            }
            if updatedTask.Description != "" {
                data.Tasks[i].Description = updatedTask.Description
            }
            ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
            return
        }
    }

    ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}



func DeleteTask(ctx *gin.Context) {
    id := ctx.Param("id")

    for i, val := range data.Tasks {
        if val.ID == id {
            data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
            ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
            return
        }
    }

    ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}


func CreateTask(ctx *gin.Context) {
    var newTask models.Task

    if err := ctx.ShouldBindJSON(&newTask); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    data.Tasks = append(data.Tasks, newTask)
    ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}