package queues

import (
	"microservice/domain/interfaces"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type QueueAdapter struct {
	ConfigMap *ckafka.ConfigMap
}

func NewQueueAdapter() interfaces.IQueueAdapter {

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	return &QueueAdapter{
		ConfigMap: &configMap,
	}
}

func (q *QueueAdapter) GetConfigMap() *ckafka.ConfigMap {
	return q.ConfigMap
}
