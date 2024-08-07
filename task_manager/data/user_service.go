package data

import (
	"context"
	"time"

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
// GetUserCount returns the number of users in the database
func GetUserCount() (int64, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    count, err := userCollection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return 0, err
    }
    return count, nil
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
func Promote(username string)  error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"isAdmin": true}})
	return err
}