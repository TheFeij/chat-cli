package services

import (
	"fmt"
	"message-persist/internal/app/core/models"
	"message-persist/internal/app/core/ports"
)

type MessagePersistService struct {
	repository ports.MessageRepository
}

func NewMessagePersistService(repository ports.MessageRepository) *MessagePersistService {
	return &MessagePersistService{
		repository: repository,
	}
}

func (s *MessagePersistService) Persist(data models.Message) error {
	return fmt.Errorf("TODO")
}
