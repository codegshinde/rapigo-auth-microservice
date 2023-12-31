package service

import (
	"context"
	"rapigo/internal/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(data interface{}) (*mongo.InsertOneResult, error) {
	collection := db.GetCollection("ProbysTestDb", "admins")

	// Insert the document into the collection
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
