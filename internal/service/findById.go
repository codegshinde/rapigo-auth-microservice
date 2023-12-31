package service

import (
	"context"
	"rapigo/internal/db"
	"rapigo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAdminByID(adminId string) (models.AdminResponse, error) {
	var admin models.AdminResponse
	collection := db.GetCollection("ProbysTestDb", "admins")

	filter := bson.D{{
		Key: "adminId", Value: adminId,
	}}

	err := collection.FindOne(context.TODO(), filter).Decode(&admin)
	return admin, err
}
