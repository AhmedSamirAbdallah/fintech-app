package model

import "fin-tech-app/pkg/enum"

type Transaction struct {
	ID              string               `bson:"_id,omitempty" json:"id"`
	AccountID       string               `bson:"accountId" json:"accountId"`
	Amount          float64              `bson:"amount" json:"amount"`
	TransactionType enum.TransactionType `bson:"transactionType" json:"transactionType"`
}
