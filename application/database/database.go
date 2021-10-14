package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
	Context context.Context
}

func NewConnection() *MongoDB {

	mongoClient := &MongoDB{}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@takeout_db:27017/?maxPoolSize=20&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	mongoClient.Context = context.Background()

	err = client.Connect(mongoClient.Context)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient.Client = client

	return mongoClient

}
