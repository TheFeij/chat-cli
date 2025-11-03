package repository

import "time"

type Message struct {
	SenderID   string    `bson:"sender_id"`
	ReceiverID string    `bson:"receiver_id"`
	Content    string    `bson:"content"`
	Timestamp  time.Time `bson:"timestamp"`
}
