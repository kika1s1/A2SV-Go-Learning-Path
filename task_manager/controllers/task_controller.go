package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/config"
	"github.com/kika1s1/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTasks(c *gin.Context){
    var tasks []models.Task
    collection := config.GetCollection("tasks")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var task models.Task
        cursor.Decode(&task)
        tasks = append(tasks, task)
    }

    c.JSON(http.StatusOK, tasks)
}



func GetTask(c *gin.Context) {
    idParam := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var task models.Task
    collection := config.GetCollection("tasks")
    err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&task)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    idParam := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var updatedTask models.Task
    err = c.BindJSON(&updatedTask)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    collection := config.GetCollection("tasks")
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    result := collection.FindOneAndUpdate(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": updatedTask}, opts)
    if result.Err() != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    var updatedTaskResult models.Task
    err = result.Decode(&updatedTaskResult)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedTaskResult)
}


func DeleteTask(c *gin.Context) {
    idParam := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    collection := config.GetCollection("tasks")
    result, _ := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func CreateTask(c *gin.Context) {
    var newTask models.Task
    err := c.BindJSON(&newTask)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newTask.ID = primitive.NewObjectID()

    collection := config.GetCollection("tasks")
    _, err = collection.InsertOne(context.Background(), newTask)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newTask)
}