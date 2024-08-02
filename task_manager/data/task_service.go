package data

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/kika1s1/task_manager/models"
	"github.com/kika1s1/task_manager/config"
)

func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	collection := config.GetCollection("tasks")
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return collection.InsertOne(context.Background(), task)
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	collection := config.GetCollection("tasks")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	collection := config.GetCollection("tasks")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func UpdateTask(id primitive.ObjectID, updatedData models.Task) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("tasks")
	updatedData.UpdatedAt = time.Now()
	return collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"title":       updatedData.Title,
				"description": updatedData.Description,
				"dueDate":     updatedData.DueDate,
				"status":      updatedData.Status,
				"updatedAt":   updatedData.UpdatedAt,
			},
		},
	)
}

func DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := config.GetCollection("tasks")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
