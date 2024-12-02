package main

import (
	"fin-tech-app/config"
	"fin-tech-app/internal/db"
	"fin-tech-app/internal/handlers"
	"fin-tech-app/internal/repository"
	"fin-tech-app/internal/router"
	"fin-tech-app/internal/service"
	"fin-tech-app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() (*mux.Router, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading environment file: %v\n", err)
	}
	client, err := db.ConnectMongo(config.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	fmt.Println("MongoDB client connected:", client)

	userRepo := repository.NewUserRepository(client, config.DatabaseName)
	accountRepo := repository.NewAccountRepository(client, config.DatabaseName)
	transactionRepo := repository.NewTransactionRepository(client, config.DatabaseName)

	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo, userRepo)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo)

	userHandler := handlers.NewUserHandler(userService)
	accountHandler := handlers.NewAccountHandler(accountService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	r := mux.NewRouter()

	router.RegisterUserRoutes(r, userHandler)
	router.RegisterAccountRoutes(r, accountHandler)
	router.RegisterTranscationRoutes(r, transactionHandler)
	utils.RegisterHealthCheckRoutes(r, client, config)

	return r, nil
}

func main() {
	r, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server running on port 8080")

}
