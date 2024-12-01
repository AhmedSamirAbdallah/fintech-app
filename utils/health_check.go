package utils

import (
	"context"
	"encoding/json"
	"fin-tech-app/internal/db"
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

func CheckDatabase() (*mongo.Client, bool) {

	client, err := db.ConnectMongo()
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return nil, false
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Printf("Error pinging MongoDB: %v", err)
		return nil, false
	}

	return client, true
}

func CheckReadOnDB(client *mongo.Client) bool {

	collection := client.Database("fintech").Collection("healthcheck")

	var result HealthCheck
	err := collection.FindOne(context.Background(), map[string]interface{}{}).Decode(&result)

	if err != nil {
		log.Printf("Error reading from healthcheck collection: %v", err)
		return false
	}
	return true
}

func CheckWriteOnDB(client *mongo.Client) bool {
	collection := client.Database("fintech").Collection("healthcheck")
	healthCheck := HealthCheck{
		Service:   "fintech",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	_, err := collection.InsertOne(context.Background(), healthCheck)
	if err != nil {
		log.Printf("Error writing to healthcheck collection: %v", err)
		return false
	}
	// _, err = collection.DeleteOne(context.Background(), )
	// if err != nil {
	// 	log.Printf("Error deleting test document from healthcheck collection: %v", err)
	// 	return false
	// }
	return true

}
func CheckKafka() bool {
	return true
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	kafkaStatus := CheckKafka()
	upTime := time.Now().String()
	client, dbConnection := CheckDatabase()

	dbStatus := map[string]interface{}{
		"connection": dbConnection,
		"read":       CheckReadOnDB(client),
		"write":      CheckWriteOnDB(client),
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

func RegisterHealthCheckRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/health-check", HealthCheckHandler)
}
