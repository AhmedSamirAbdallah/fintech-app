package repository

import (
	"context"
	"fin-tech-app/internal/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	Collection *mongo.Collection
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *model.Transaction) error {
	log.Printf("Inserting Transaction ... %v", transaction)
	res, err := r.Collection.InsertOne(ctx, &transaction)
	if err != nil {
		return fmt.Errorf("failed to insert transaction %v: %w", transaction.ID, err)
	}
	log.Printf("Transaction created successfully with ID: %v", res.InsertedID)
	return nil
}

func (r *TransactionRepository) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	filter := map[string]interface{}{}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions %w", err)

	}
	defer cursor.Close(ctx)
	var transactions []model.Transaction
	for cursor.Next(ctx) {
		var transaction model.Transaction

		err = cursor.Decode(&transaction)
		if err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %w", err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
