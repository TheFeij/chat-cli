package models

import "time"

type Message struct {
	ID       string
	SenderID string
	ChatID   string
	Content  string
	Time     time.Time
}
