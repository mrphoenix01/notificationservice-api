// kafka/producer.go
package kafka

import (
	"encoding/json"
	"log"
	"notification_service/model"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

// InitializeKafkaProducer initializes the Kafka producer
func InitializeKafkaProducer(brokers []string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	log.Println("Kafka producer initialized")
}

// PublishMessage publishes a notification message to the specified Kafka topic
func PublishMessage(topic string, notification model.Notification) error {
	message, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	})
	return err
}
