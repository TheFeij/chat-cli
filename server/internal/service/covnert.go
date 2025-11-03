package service

import (
	"server/internal/pb/chatpb"
	"server/internal/repository"
)

func convertMessageProtoToModel(in *chatpb.ChatMessage) *repository.Message {
	return &repository.Message{
		SenderID:   in.SenderId,
		ReceiverID: in.ReceiverId,
		Content:    in.Content,
	}
}
