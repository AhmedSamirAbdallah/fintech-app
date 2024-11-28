package model

import "time"

type User struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	FirstName      string    `bson:"firstName" json:"firstName"`
	LastName       string    `bson:"lastName" json:"lastName"`
	Email          string    `bson:"email" json:"email"`
	Ssn            string    `bson:"ssn" json:"ssn"`
	PhoneNumber    string    `bson:"phoneNumber" json:"phoneNumber"`
	BirthDate      time.Time `bson:"birthDate" json:"birthDate"`
	Address        string    `bson:"address" json:"address"`
	SecurityAnswer string    `bson:"securityAnswer" json:"securityAnswer"`
}

// {
//     "firstName": "Ahmed",
//     "lastName": "Samir",
//     "email": "ahmedsamir661998@gmail.com",
//     "ssn": "123-45-36789",
//     "phoneNumber": "+12345672890",
// "birthDate": "1990-01-01T00:00:00Z",
//     "address": "123 Main St, Anytown, USA",
//     "securityAnswer": "My first pet's name"
// }
