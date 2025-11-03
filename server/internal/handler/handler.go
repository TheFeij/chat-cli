package handler

import (
	"google.golang.org/grpc"
	"server/internal/pb/chatpb"
	"server/internal/service"
)

type ChatServer struct {
	chatpb.UnimplementedChatServiceServer
}

func (c ChatServer) Chat(g grpc.BidiStreamingServer[chatpb.ChatMessage, chatpb.ChatMessage]) error {
	return service.Chat(g)
}

var _ chatpb.ChatServiceServer = (*ChatServer)(nil)
