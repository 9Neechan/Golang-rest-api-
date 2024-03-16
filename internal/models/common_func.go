package models

import (
	"context"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

// Возвращает несколько документов из коллекции по заданному filter и opts
func FindFewTasksInMomgo(filter primitive.D, opts *options.FindOptions, coll *mongo.Collection) ([]Task, error) { // []any
	returnTasks := make([]Task, 0)

	cur, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return make([]Task, 0), err
	}

	// проходим по всем items и добавляем в массив
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result Task

		err := cur.Decode(&result)
		if err != nil {
			return make([]Task, 0), err
		}

		returnTasks = append(returnTasks, result)
	}
	if err := cur.Err(); err != nil {
		return make([]Task, 0), err
	}

	return returnTasks, nil
}