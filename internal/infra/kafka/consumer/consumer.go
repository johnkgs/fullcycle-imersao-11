package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/kafka/factory"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

func Consume(topics []string, servers string, msgChannel chan *kafka.Message) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "gostats",
		"auto.offset.reset": "earliest",
	}
	kafkaConsumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChannel <- msg
		}
	}
}

func ProcessEvents(ctx context.Context, msgChannel chan *kafka.Message, uwo uow.UowInterface) {
	for msg := range msgChannel {
		fmt.Println("Received message", string(msg.Value), "on topic", *msg.TopicPartition.Topic)
		strategy := factory.CreateProcessMessageStrategy(*msg.TopicPartition.Topic)
		err := strategy.Process(ctx, msg, uwo)
		if err != nil {
			fmt.Println(err)
		}
	}
}
