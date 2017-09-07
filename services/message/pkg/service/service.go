package service

import "context"

// MessageService provides operations on messages.
type MessageService interface {
	GetMessages(ctx context.Context) ([]Message, error)
	GetMessage(ctx context.Context, mID string) (Message, error)
	PostMessage(ctx context.Context, m Message) error
}

// Message represents a single message.
// ID should be globally unique.
type Message struct {
	ID   string `json:"id"`
	Body string `json:"body,omitempty"`
}
