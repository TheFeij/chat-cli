package ports

import "message-persist/internal/app/core/models"

type MessageRepository interface {
	Create(data models.Message) error
}

type MessagePersistService interface {
	Persist(data models.Message) error
}
