package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kika1s1/task_manager/internal/infrastructure/db"
	"github.com/kika1s1/task_manager/internal/infrastructure/middleware"
	"github.com/kika1s1/task_manager/internal/interfaces/http"
	"github.com/kika1s1/task_manager/internal/interfaces/repository"
	"github.com/kika1s1/task_manager/internal/usecase"
)

func main() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    dbURI := os.Getenv("MONGO_URI")
    jwtSecret := os.Getenv("JWT_SECRET")

    mongoDB := db.NewMongoClient(dbURI)

    taskRepo := repository.NewMongoTaskRepository(mongoDB)
    taskUsecase := usecase.NewTaskUsecase(taskRepo)

    userRepo := repository.NewMongoUserRepository(mongoDB)
    userUsecase := usecase.NewUserUsecase(userRepo, jwtSecret)

    router := gin.Default()

    

    http.NewAuthHandler(router, userUsecase)
	router.Use(middleware.AuthMiddleware(jwtSecret))
    http.NewTaskHandler(router, taskUsecase)

    log.Fatal(router.Run(":3000"))
}
