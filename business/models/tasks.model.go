package models

import (
	"context"

	task "github.com/matheuscaet/go-api-template/business/types"
	"github.com/matheuscaet/go-api-template/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	databaseName   = "go-api-template"
	collectionName = "tasks"
)

func GetTasks(ctx context.Context) ([]task.Task, error) {
	client := database.Connect()
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []task.Task{}, err
	}
	defer cursor.Close(ctx)

	var tasks []task.Task
	for cursor.Next(ctx) {
		var ntask task.Task
		if err := cursor.Decode(&ntask); err != nil {
			return []task.Task{}, err
		}
		tasks = append(tasks, ntask)
	}
	return tasks, nil
}

func CreateTask(ctx context.Context, task task.Task) (task.Task, error) {
	client := database.Connect()
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func UpdateTask(ctx context.Context, task task.Task) (task.Task, error) {
	client := database.Connect()
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)
	_, err := collection.UpdateOne(ctx, bson.M{"id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return task, err
	}
	return task, nil
}

func DeleteTask(ctx context.Context, id string) error {
	client := database.Connect()
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
