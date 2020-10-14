package messenger

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID `json:"id"`
	From        int       `json:"from"`
	Destination int       `json:"to"`
	Sent        time.Time `json:"sent"`
	Message     string    `json:"msg"`
}

type NewMessageRequest struct {
	From    int    `json:"from"`
	To      string `json:"to"`
	Message string `json:"msg"`
}

func NewMessage(from int, to int) *Message {
	return &Message{
		ID:          uuid.New(),
		From:        from,
		Destination: to,
		Sent:        time.Now(),
	}
}

func NewMessageWithMessage(from int, to int, msg string) *Message {
	return &Message{
		ID:          uuid.New(),
		From:        from,
		Destination: to,
		Sent:        time.Now(),
		Message:     msg,
	}
}
