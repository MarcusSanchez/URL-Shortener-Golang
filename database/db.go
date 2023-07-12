package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Route = "127.0.0.1:3000"
var mongoClient *mongo.Client

func SetupMongoDB() {
	var err error
	var instance = "mongodb://localhost:27017"
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(instance))
	if err != nil {
		log.Fatal(err)
	}
}

func ShutDownMongoDB() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database("URLShortener").Collection(name)
}
