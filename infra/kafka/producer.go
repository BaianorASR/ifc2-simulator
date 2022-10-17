package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// NewKafkaProducer creates a new kafka producer
func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return p
}

// Publish publishes a message to kafka
func Publish(msg string, topic string, producer *kafka.Producer) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
