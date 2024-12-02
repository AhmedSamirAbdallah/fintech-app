package kafka

import (
	"context"
	"log"
	"strings"

	"github.com/IBM/sarama"
)

// StartConsumer starts a Kafka consumer group
func StartConsumer(groupID, topic string) {
	brokers := strings.Split(config.GetEnv("KAFKA_BROKERS", "localhost:9092"), ",")

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Failed to create consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx := context.Background()

	for {
		// Pass the consumer handler to consume messages
		handler := &ConsumerGroupHandler{}
		if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

// ConsumerGroupHandler implements the sarama.ConsumerGroupHandler interface
type ConsumerGroupHandler struct{}

// Setup is run at the beginning of a new session
func (h *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session
func (h *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes Kafka messages
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, topic = %s, partition = %d, offset = %d\n",
			string(message.Value), message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}
	return nil
}
