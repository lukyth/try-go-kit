package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

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
	if mID != "1" {
		return Message{}, ErrNotFound
	}
	m = Message{
		ID:   mID,
		Body: mID + " message",
	}
	return m, nil
}

// Implement the business logic of PostMessage
func (s *stubMessageService) PostMessage(ctx context.Context, m Message) (e error) {
	url := "http://localhost:8081/count"

	var jsonStr = []byte(fmt.Sprintf(`{"s":"%s"}`, m.Body))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(m.ID)
	fmt.Println(m.Body)
	fmt.Println("Message count: " + string(body))
	return nil
}
