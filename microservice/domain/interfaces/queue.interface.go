package interfaces

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type IQueueAdapter interface {
	GetConfigMap() *ckafka.ConfigMap
}

type IMessageHandler interface {
	Handle(msg *ckafka.Message) error
}
