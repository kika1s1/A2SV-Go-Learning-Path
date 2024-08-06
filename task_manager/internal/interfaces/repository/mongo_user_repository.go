package repository

import (
	"context"
	"errors"

	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

// FindByUsername implements repository.UserRepository.
func (r *MongoUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
    err := r.Collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    if err!= nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, nil // No user found
        }
        return nil, err // An error occurred
    }
    return &user, nil
    
}

func (r *MongoUserRepository) FindByEmail(username string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // No user found
		}
		return nil, err // An error occurred
	}
	return &user, nil
}

func (r *MongoUserRepository) Save(user *models.User) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	return err
}
func NewMongoUserRepository(db *mongo.Database) repository.UserRepository {
	return &MongoUserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) Create(user *models.User) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoUserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}
