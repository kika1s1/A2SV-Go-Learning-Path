package main

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/kika1s1/task_manager/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil{
     log.Fatalf("Error loading .env file: %s", err)
    }
    uri := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	r := router.SetupRouter(client)

	// Start server
	r.Run(":3000")
}