package utils

import (
	"context"
	"encoding/json"
	"fin-tech-app/config"
	"fin-tech-app/internal/kafka"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/IBM/sarama"
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
func CheckKafka(brokers string) bool {
	brokerList := strings.Split(brokers, ",")
	log.Printf("brokers : %v", brokerList)
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.ClientID = "health-check-client"

	admin, err := sarama.NewClusterAdmin(brokerList, config)
	if err != nil {
		log.Printf("failed to create Kafka admin client: %v", err)
		return false
	}
	defer admin.Close()

	brokersList, _, err := admin.DescribeCluster()
	if err != nil {
		log.Printf("failed to describe cluster: %v", err)
		return false
	}
	log.Printf("Kafka health check successful: found %d broker(s)\n", len(brokersList))
	return true
}

func CheckProduce(brokers string, topic string) bool {
	producer, err := kafka.CreateProducer(brokers)
	if err != nil {
		log.Printf("Error creating Kafka producer: %v", err)
		return false
	}
	defer producer.Close()
	err = kafka.SendMessage(producer, topic, "produce within the health check")
	if err != nil {
		log.Printf("Error sending message to Kafka topic %s: %v", topic, err)
		return false
	}
	return true
}

func CheckConsume(brokers string, topic string) bool {
	return true
}

func HealthCheckHandler(client *mongo.Client, config *config.Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		upTime := time.Now().String()

		dbStatus := map[string]interface{}{
			"connection": CheckDatabase(client),
			"read":       CheckReadOnDB(client, config.DatabaseName),
			"write":      CheckWriteOnDB(client, config.DatabaseName),
		}

		kafkaStatus := map[string]interface{}{
			"connection": CheckKafka(config.KafkaBroker),
			"produce":    CheckProduce(config.KafkaBroker, config.KafkaTopic),
			"consume":    CheckConsume(config.KafkaBroker, config.KafkaTopic),
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
