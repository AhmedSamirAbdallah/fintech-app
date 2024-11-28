package repository

import (
	"context"
	"fin-tech-app/internal/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {

	log.Printf("Inserting User: %v\n", user)
	result, err := r.Collection.InsertOne(ctx, user)

	if err != nil {
		// handle specific errors
		// if mongo.IsDuplicateKeyError(err) {
		// 	return fmt.Errorf("account already exists: %w", err)
		// }
		return fmt.Errorf("failed to insert user %v: %w", user.ID, err)
	}

	log.Printf("User created successfully with ID: %v", result.InsertedID)

	return nil

}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	err := r.Collection.FindOne(ctx, map[string]string{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := r.Collection.Find(ctx, map[string]interface{}{})
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
