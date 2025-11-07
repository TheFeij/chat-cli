package mongodb

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"message-persist/internal/app/core/models"
	"message-persist/internal/app/core/ports"
	"message-persist/internal/app/infrastructure/config"
	"time"
)

type MessageRepository struct {
	collection *mongo.Collection
}

var _ ports.MessageRepository = (*MessageRepository)(nil)

func NewMessageRepository(client *mongo.Client, config *config.AppConfig) *MessageRepository {
	collection := client.Database("TODO").Collection(config.Mongo.Collection)

	return &MessageRepository{collection: collection}
}

func (repo *MessageRepository) Create(data models.Message) error {
	doc := map[string]interface{}{
		"id":        uuid.New().String(),
		"sender_id": data.SenderID,
		"chat_id":   data.ChatID,
		"content":   data.Content,
		"timestamp": time.Now(),
	}

	_, err := repo.collection.InsertOne(context.Background(), doc)

	return err
}
