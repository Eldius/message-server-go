package messenger

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID      uuid.UUID
	From    int
	To      int
	Sent    time.Time
	Message string
}

func NewMessage(from int, to int) *Message {
	return &Message{
		ID:   uuid.New(),
		From: from,
		To:   to,
		Sent: time.Now(),
	}
}

func NewMessageWithMessage(from int, to int, msg string) *Message {
	return &Message{
		ID:      uuid.New(),
		From:    from,
		To:      to,
		Sent:    time.Now(),
		Message: msg,
	}
}
