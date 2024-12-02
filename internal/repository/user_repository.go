package repository

import (
	"context"
	"fin-tech-app/internal/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Database
}

func NewUserRepository(client *mongo.Client, databaseName string) *UserRepository {
	return &UserRepository{
		Collection: client.Database(databaseName),
	}
}

func (r *UserRepository) getUsersCollection() *mongo.Collection {
	return r.Collection.Collection("users")
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	usersCollection := r.getUsersCollection()

	log.Printf("Inserting User: %v\n", user)
	result, err := usersCollection.InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to insert user %v: %w", user.ID, err)
	}

	log.Printf("User created successfully with ID: %v", result.InsertedID)

	return nil

}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}
	usersCollection := r.getUsersCollection()
	err = usersCollection.FindOne(ctx, map[string]interface{}{"_id": objectID}).Decode(&user)
	if err != nil {
		log.Printf("Error fetching user with ID %v: %v", id, err)
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user with ID %v not found", id)
		}
		return nil, fmt.Errorf("failed to fetch user with ID %v: %w", id, err)
	}
	return &user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	usersCollection := r.getUsersCollection()

	cursor, err := usersCollection.Find(ctx, map[string]interface{}{})
	if err != nil {
		log.Printf("Error fetching users %v", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User

		err := cursor.Decode(&user)
		if err != nil {
			log.Printf("Error decoding user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	err = cursor.Err()
	if err != nil {
		log.Printf("Cursor iteration error: %v", err)
		return nil, err
	}
	return users, err

}
