package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	mongoURI := "mongodb://admin:adminpassword@localhost:27017/?authSource=admin"

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("Connection Sucess ! %v", client)

	return client, nil
}
