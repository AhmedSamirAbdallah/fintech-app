package main

import (
	"fin-tech-app/internal/db"
	"fmt"
	"log"
)

func main() {
	client, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)

	}
	fmt.Println("MongoDB client connected:", client)

}
