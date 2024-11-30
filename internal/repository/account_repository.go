package repository

import (
	"context"
	"fin-tech-app/internal/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	Collection *mongo.Collection
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account *model.Account) error {

	log.Printf("Inserting Account ... %v", account)
	res, err := r.Collection.InsertOne(ctx, account)
	if err != nil {
		return fmt.Errorf("failed to insert account %v: %w", account.ID, err)
	}
	log.Printf("Account created successfully with ID: %v", res.InsertedID)
	return nil
}

func (r *AccountRepository) GetAccountById(ctx context.Context, id string) (*model.Account, error) {
	var account model.Account
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}
	err = r.Collection.FindOne(ctx, map[string]interface{}{"_id": objectID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account with ID %v not found", id)
		}
		return nil, fmt.Errorf("failed to fetch account %v: %w", id, err)
	}
	return &account, nil
}

func (r *AccountRepository) GetAccounts(ctx context.Context) ([]model.Account, error) {
	var accounts []model.Account

	cursor, err := r.Collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch accounts %w", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var account model.Account

		err = cursor.Decode(&account)
		if err != nil {
			return nil, fmt.Errorf("failed to decode account: %w", err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil

}

func (r *AccountRepository) DeleteAccount(ctx context.Context, id string) error {
	res, err := r.Collection.DeleteOne(ctx, map[string]string{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete account with ID %v: %w", id, err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("account with ID %v not found", id)
	}
	log.Printf("Account with ID %v successfully deleted", id)
	return nil
}

func (r *AccountRepository) GetAccountBalance(ctx context.Context, id string) (float64, error) {
	ObjecrID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := map[string]interface{}{
		"_id": ObjecrID,
	}
	var account model.Account

	err = r.Collection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("account with id %s not found", id)
		}
		return 0, fmt.Errorf("failed to get balance for account %s: %v", id, err)
	}

	return account.Balance, nil
}

func (r *AccountRepository) UpdateAccount(ctx context.Context, id string, updatedAccount *model.Account) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}
	updateFields := map[string]interface{}{}
	if updatedAccount.Balance != 0 {
		updateFields["balance"] = updatedAccount.Balance
	}
	if updatedAccount.Status != "" {
		updateFields["status"] = updatedAccount.Status
	}
	if len(updateFields) == 0 {
		return fmt.Errorf("no fields to update")
	}
	filter := map[string]interface{}{
		"_id": objectId,
	}
	update := map[string]interface{}{
		"$set": updateFields,
	}
	res, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update account with ID %v: %w", id, err)
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("account with ID %v not found", id)
	}
	log.Printf("Account with ID %v successfully updated", id)

	return nil
}
