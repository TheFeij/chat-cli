package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"message-persist/internal/app/infrastructure/config"
)

func NewKafkaReader(config *config.AppConfig) *kafka.Reader {
	kafkaURL := config.Kafka.Broker
	if kafkaURL == "" {
		panic("kafka broker url config is required")
	}

	topic := config.Kafka.Topic
	if topic == "" {
		panic("kafka topic config is required")
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1e9,
		GroupID:  config.Kafka.GroupId,
	})
}

func registerKafkaConsumer(lifecycle fx.Lifecycle, kcs *KafkaConsumerService) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return kcs.Start()
		},
		OnStop: func(ctx context.Context) error {
			return kcs.Stop()
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		NewKafkaReader,
		NewKafkaConsumerService,
	),
	fx.Invoke(registerKafkaConsumer),
)
