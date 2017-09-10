package service

import (
	"context"
	"fmt"
)

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

type stubMessageService struct{}

// New return a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New() (s MessageService) {
	s = &stubMessageService{}
	return s
}

// Implement the business logic of GetMessages
func (s *stubMessageService) GetMessages(ctx context.Context) (ms []Message, e error) {
	ms = []Message{
		Message{
			ID:   "1",
			Body: "first message",
		},
		Message{
			ID:   "2",
			Body: "second message",
		},
	}
	return ms, nil
}

// Implement the business logic of GetMessage
func (s *stubMessageService) GetMessage(ctx context.Context, mID string) (m Message, e error) {
	m = Message{
		ID:   mID,
		Body: mID + " message",
	}
	return m, nil
}

// Implement the business logic of PostMessage
func (s *stubMessageService) PostMessage(ctx context.Context, m Message) (e error) {
	fmt.Println(m.ID)
	fmt.Println(m.Body)
	fmt.Println("Saved")
	return nil
}
