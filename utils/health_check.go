package utils

import (
	"context"
	"encoding/json"
	"fin-tech-app/config"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type HealthCheck struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	Service   string `bson:"service" json:"service"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
}

type HealthCheckResponse struct {
	Status       string                 `json:"status"`
	UpTime       string                 `json:"upTime"`
	Dependancies map[string]interface{} `json:"dependancies"`
}

func CheckDatabase(client *mongo.Client) bool {
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Printf("Error pinging MongoDB: %v", err)
		return false
	}
	return true
}

func CheckReadOnDB(client *mongo.Client, databaseName string) bool {
	collection := client.Database(databaseName).Collection("healthcheck")
	var result HealthCheck
	err := collection.FindOne(context.Background(), map[string]interface{}{}).Decode(&result)
	if err != nil {
		log.Printf("Error reading from healthcheck collection: %v", err)
		return false
	}
	log.Printf("%v", result)
	return true
}

func CheckWriteOnDB(client *mongo.Client, databaseName string) bool {
	collection := client.Database(databaseName).Collection("healthcheck")
	healthCheck := HealthCheck{
		Service:   "healthcheck",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	_, err := collection.InsertOne(context.Background(), healthCheck)
	if err != nil {
		log.Printf("Error writing to healthcheck collection: %v", err)
		return false
	}
	return true
}
func CheckKafka() bool {
	return true
}

func HealthCheckHandler(client *mongo.Client, config *config.Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		kafkaStatus := CheckKafka()
		upTime := time.Now().String()

		dbStatus := map[string]interface{}{
			"connection": CheckDatabase(client),
			"read":       CheckReadOnDB(client, config.DatabaseName),
			"write":      CheckWriteOnDB(client, config.DatabaseName),
		}

		response := HealthCheckResponse{
			Status: "UP",
			UpTime: upTime,
			Dependancies: map[string]interface{}{
				"database": dbStatus,
				"kafka":    kafkaStatus,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func RegisterHealthCheckRoutes(mux *mux.Router, client *mongo.Client, config *config.Config) {
	mux.HandleFunc("/api/health-check", HealthCheckHandler(client, config))
}
