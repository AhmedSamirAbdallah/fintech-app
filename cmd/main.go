package main

import (
	"fin-tech-app/internal/db"
	"fin-tech-app/internal/handlers"
	"fin-tech-app/internal/repository"
	"fin-tech-app/internal/router"
	"fin-tech-app/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	client, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)

	}
	fmt.Println("MongoDB client connected:", client)

	// Initialize the UserRepository with the MongoDB client
	userRepo := &repository.UserRepository{
		Collection: client.Database("fintech").Collection("users"), // use the correct database and collection name
	}

	// Initialize the UserService with the UserRepository
	userService := &service.UserService{
		UserRepo: userRepo,
	}

	// Initialize the UserHandler with the UserService
	userHandler := &handlers.UserHandler{
		UserService: userService,
	}

	r := mux.NewRouter()

	router.RegisterUserRoutes(r, userHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server running on port 8080")

}
