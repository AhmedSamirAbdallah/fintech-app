package repository

import (
	"context"
	"fin-tech-app/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {

	log.Printf("Inserting User: %v\n", user)
	_, err := r.Collection.InsertOne(ctx, user)

	if err != nil {
		// handle specific errors
		// if mongo.IsDuplicateKeyError(err) {
		// 	return fmt.Errorf("account already exists: %w", err)
		// }
		return err
	}
	log.Print("User created successfully")
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
