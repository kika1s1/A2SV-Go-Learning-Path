package data

import (
	"context"
	"github.com/kika1s1/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	userCollection = client.Database("task_manager").Collection("users")
}

// CreateUser adds a new user to the database
func CreateUser(user models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}
