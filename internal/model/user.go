package model

type User struct {
	ID       string    `bson:"_id" json:"id"`
	Name     string    `bson:"name" json:"name"`
	Accounts []Account `bson:"accounts" json:"accounts"`
}
