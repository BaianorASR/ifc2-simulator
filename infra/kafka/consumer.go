package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// Consumer is a kafka consumer
type Consumer struct {
	MsgChan chan *kafka.Message
}

// NewKafkaConsumer creates a new kafka consumer
func NewKafkaConsumer(msgChan chan *kafka.Message) *Consumer {
	return &Consumer{
		MsgChan: msgChan,
	}
}

// Consumer is a kafka consumer constructor
func (k *Consumer) Consumer() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("error consuming kafka message: " + err.Error())
	}

	topics := []string{os.Getenv("KAFKA_READ_TOPIC")}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		return
	}
	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			k.MsgChan <- msg
		}
	}
}
