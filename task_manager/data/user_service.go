package data

import (
	"context"
	"github.com/kika1s1/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserCollection(client *mongo.Client) {
	userCollection = client.Database("task_manager").Collection("users")
}

func CreateUser(user models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}