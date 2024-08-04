package data

import (
	"context"
	"errors"

	"github.com/kika1s1/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var taskCollection *mongo.Collection

func InitTaskCollection(client *mongo.Client) {
	taskCollection = client.Database("task_manager").Collection("tasks")
}

func CreateTask(task models.Task) error {
	_, err := taskCollection.InsertOne(context.Background(), task)
	return err
}

func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, errors.New("invalid ID format")
	}

	err = taskCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, err
	}
	return task, nil
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTask(task models.Task) error {
	_, err := taskCollection.ReplaceOne(context.Background(), bson.M{"_id": task.ID}, task)
	return err
}

func DeleteTask(id string) error {
	_, err := taskCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}