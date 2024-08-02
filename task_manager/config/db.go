package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
func ConnectDB() {
    err := godotenv.Load(".env")
    if err != nil{
     log.Fatalf("Error loading .env file: %s", err)
    }
    uri := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")

    Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("task_manager").Collection(collectionName)
}

func DisconnectDB() {
    err := Client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Disconnected from MongoDB")
}