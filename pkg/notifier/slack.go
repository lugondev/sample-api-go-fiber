package notifier

import (
	"fmt"
	"github.com/slack-go/slack"
)

type SlackConfig struct {
	WebhookUrl string
}

func NewSlack(configuration SlackConfig) *SlackClient {
	return &SlackClient{configuration}
}

type SlackClient struct {
	config SlackConfig
}

func (c *SlackClient) Notify(message Message) error {
	webhook := slack.WebhookMessage{
		Text: message.Content,
	}

	if err := slack.PostWebhook(c.config.WebhookUrl, &webhook); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
