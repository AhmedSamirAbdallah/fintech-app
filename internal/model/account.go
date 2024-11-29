package model

import (
	"fin-tech-app/pkg/enum"
)

type Account struct {
	ID       string             `bson:"_id,omitempty" json:"id"`
	UserID   string             `bson:"userId" json:"userId"`
	Status   enum.AccountStatus `bson:"status" json:"status"`
	Balance  float64            `bson:"balance" json:"balance"`
	Currency enum.Currency      `bson:"currency" json:"currency"`
}
