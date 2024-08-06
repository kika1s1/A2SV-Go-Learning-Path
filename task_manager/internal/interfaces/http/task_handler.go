package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)


type TaskHandler struct {
    TaskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(router *gin.Engine, us *usecase.TaskUsecase) {
    handler := &TaskHandler{
        TaskUsecase: us,
    }

    router.POST("/tasks",  handler.CreateTask)
    router.GET("/tasks/:id", handler.GetTaskByID)
    router.PUT("/tasks/:id", handler.UpdateTask)
    router.DELETE("/tasks/:id", handler.DeleteTask)
    router.GET("/tasks", handler.GetAllTasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.TaskUsecase.CreateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    task, err := h.TaskUsecase.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}



func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID, _ = primitive.ObjectIDFromHex(id)

    if err := h.TaskUsecase.UpdateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id := c.Param("id")

    if err := h.TaskUsecase.DeleteTask(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
    tasks, err := h.TaskUsecase.GetAllTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

