package notifier

import (
	"log"
)

func NewMock() Notifier {
	return &mockClient{}
}

type mockClient struct {
}

func (c *mockClient) Notify(message Message) error {
	log.Println(message)

	return nil
}
