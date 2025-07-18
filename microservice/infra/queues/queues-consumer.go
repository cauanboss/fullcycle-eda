package queues

import (
	"fmt"
	"log"
	"microservice/domain/interfaces"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
	Handler   interfaces.IMessageHandler
}

func NewConsumer(queueAdapter interfaces.IQueueAdapter, topics []string, handler interfaces.IMessageHandler) *Consumer {
	configMap := queueAdapter.GetConfigMap()
	if configMap == nil {
		log.Fatal("ConfigMap is nil")
	}

	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
		Handler:   handler,
	}
}

func (c *Consumer) Consume() error {
	fmt.Println("Consuming topics: ", c.Topics)

	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topics: %v", err)
	}

	fmt.Printf("Consumer started successfully for topics: %v\n", c.Topics)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Message received: %s\n", string(msg.Value))
		if c.Handler != nil {
			if err := c.Handler.Handle(msg); err != nil {
				fmt.Printf("Error processing message: %v\n", err)
			}
		}
	}
}

func (c *Consumer) Start() {
	go func() {
		err := c.Consume()
		if err != nil {
			log.Printf("Consumer error: %v", err)
		}
	}()
}
