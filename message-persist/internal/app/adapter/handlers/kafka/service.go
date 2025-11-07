package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"message-persist/internal/app/core/models"
	"message-persist/internal/app/core/ports"
)

type KafkaConsumerService struct {
	reader         *kafka.Reader
	persistService ports.MessagePersistService
}

func NewKafkaConsumerService(reader *kafka.Reader, dcs ports.MessagePersistService) *KafkaConsumerService {
	return &KafkaConsumerService{reader: reader, persistService: dcs}
}

func (k *KafkaConsumerService) Start() error {
	log.Println("Starting Kafka consumer...")
	go func() {
		for {
			m, err := k.reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Failed to read messages: %v", err)
				continue
			}
			log.Println("Received message: ", string(m.Key))

			// unmarshal msg

			err = k.persistService.Persist(models.Message{})
			if err != nil {
				log.Printf("Failed to persist message: %v", err)
			}
		}
	}()
	return nil
}

func (k *KafkaConsumerService) Stop() error {
	log.Println("Stopping Kafka consumer...")
	return k.reader.Close()
}
