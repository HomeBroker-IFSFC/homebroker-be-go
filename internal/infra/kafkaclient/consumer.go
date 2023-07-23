package kafkaclient

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *kafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *kafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *kafka.Message) {
	consumer, err := kafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			msgChan <- msg
		}
	}

}
