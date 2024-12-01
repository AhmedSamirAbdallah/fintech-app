package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	// TO BE READ FROM ENV

	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Printf("Error loading environment file: %v\n", err)
	}

	// Read the Mongo URI from the environment
	mongoURI := os.Getenv("mongoURI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("Connection Sucess ! %v", client)

	return client, nil
}
