package repository

import (
	"context"

	"github.com/kika1s1/task_manager/internal/domain/models"
	"github.com/kika1s1/task_manager/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
    Collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) repository.TaskRepository {
    return &MongoTaskRepository{
        Collection: db.Collection("tasks"),
    }
}

func (r *MongoTaskRepository) Create(task *models.Task) error {
    _, err := r.Collection.InsertOne(context.Background(), task)
    return err
}

func (r *MongoTaskRepository) GetByID(id string) (*models.Task, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var task models.Task
    err = r.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
    return &task, err
}

func (r *MongoTaskRepository) Update(task *models.Task) error {
    objID, err := primitive.ObjectIDFromHex(task.ID.Hex())
    if err != nil {
        return err
    }
    _, err = r.Collection.ReplaceOne(context.Background(), bson.M{"_id": objID}, task)
    return err
}

func (r *MongoTaskRepository) Delete(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}

func (r *MongoTaskRepository) GetAll() ([]*models.Task, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    var tasks []*models.Task
    if err = cursor.All(context.Background(), &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}
