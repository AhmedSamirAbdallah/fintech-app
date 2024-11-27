package model

import (
	"fin-tech-app/pkg/enum"
)

type Account struct {
	ID           string             `bson:"_id" json:"id"`
	UserID       string             `bson:"userId" json:"userId"`
	AccountName  string             `bson:"accountName" json:"accountName"`
	Status       enum.AccountStatus `bson:"status" json:"status"`
	Balance      float64            `bson:"balance" json:"balance"`
	Transactions []Transaction      `bson:"transactions" json:"transactions"`
}
