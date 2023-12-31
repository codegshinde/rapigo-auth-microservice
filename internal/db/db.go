// db/db.go
package db

import (
	"context"
	"rapigo/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() error {
	clientOptions := options.Client().ApplyURI(config.GetEnvVariable("MONGODB_URI"))
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	return err
}

// GetCollection returns a MongoDB collection
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func GetClient() *mongo.Client {
	return client
}
