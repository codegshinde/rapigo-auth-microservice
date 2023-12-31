package service

import (
	"context"
	"rapigo/internal/db"
	"rapigo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

type FeildFilter struct {
	Key   string
	Value string
}

func GetFindMyField(fieldKey string, feildValue string) (models.AdminResponse, error) {
	var admin models.AdminResponse
	collection := db.GetCollection("ProbysTestDb", "admins")

	filter := bson.D{{
		Key: fieldKey, Value: feildValue,
	}}

	err := collection.FindOne(context.TODO(), filter).Decode(&admin)
	return admin, err
}
