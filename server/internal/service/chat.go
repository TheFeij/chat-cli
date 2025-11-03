package service

import (
	"google.golang.org/grpc"
	"log"
	"server/internal/pb/chatpb"
)

type Repository struct {
	messages  chan *chatpb.ChatMessage
	usersChan []chan *chatpb.ChatMessage
}

func (r *Repository) Chat(g grpc.BidiStreamingServer[chatpb.ChatMessage, chatpb.ChatMessage]) error {
	go func() {
		msg, err := g.Recv()
		if err != nil {
			log.Fatal(err) // TODO: what to do with this error?
		}

		r.messages <- msg
	}()

	for {
		message := <-r.messages

		err := g.Send(message)
		if err != nil {
			return err
		}
	}
}

func chatWorker(r *Repository) {
	for {
		msg := <-r.messages

		for _, usersChan := range r.usersChan {
			usersChan <- msg
		}
	}
}
