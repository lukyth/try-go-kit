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

type stubMessageService struct{}

// Get a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New() (s *stubMessageService) {
	s = &stubMessageService{}
	return s
}

// Implement the business logic of GetMessages
func (me *stubMessageService) GetMessages(ctx context.Context) (M0 []Message, e1 error) {
	return M0, e1
}

// Implement the business logic of GetMessage
func (me *stubMessageService) GetMessage(ctx context.Context, mID string) (M0 Message, e1 error) {
	return M0, e1
}

// Implement the business logic of PostMessage
func (me *stubMessageService) PostMessage(ctx context.Context, m Message) (e0 error) {
	return e0
}
