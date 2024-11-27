package model

type Transaction struct {
	ID        string  `bson:"_id,omitempty" json:"id"`
	AccountID string  `bson:"accountId" json:"accountId"`
	Amount    float64 `bson:"amount" json:"amount"`
	Type      string  `bson:"type" json:"type"` // "credit" or "debit"
}
