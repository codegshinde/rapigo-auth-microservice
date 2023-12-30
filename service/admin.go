// service/admin.go
package service

import (
	"context"
	"rapigo/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Admin struct {
	AdminId  string `json:"adminId"`
	Password string `json:"password"`
}

func GetAdminByID(adminID string) (Admin, error) {
	var admin Admin
	collection := db.GetCollection("ProbysTestDb", "admins")

	filter := bson.D{{
		Key: "adminId", Value: adminID,
	}}

	err := collection.FindOne(context.TODO(), filter).Decode(&admin)
	return admin, err
}
