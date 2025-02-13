package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserColl *mongo.Collection
var AttendColl *mongo.Collection

func InitDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	UserColl = client.Database("attendance").Collection("users")
	AttendColl = client.Database("attendance").Collection("attendance")
}
