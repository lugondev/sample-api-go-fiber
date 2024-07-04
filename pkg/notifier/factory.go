package notifier

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Adapter string

const (
	Slack    Adapter = "slack"
	Mock     Adapter = "mock"
	Telegram Adapter = "telegram"
)

type Config struct {
	Adapter  Adapter
	Slack    SlackConfig
	Telegram TelegramConfig
}

func NewNotifier() (Notifier, error) {
	var config Config
	err := viper.UnmarshalKey("notifier", &config)
	if err != nil {
		return nil, err
	}

	switch config.Adapter {
	case Mock:
		log.Println("Initialize Mock notifier")
		return NewMock(), nil
	case Slack:
		log.Println("Initialize Slack notifier")
		return NewSlack(config.Slack), nil
	case Telegram:
		log.Println("Initialize Telegram notifier")
		return NewTelegram(config.Telegram), nil
	default:
		return nil, errors.New("unable to find adapter")
	}
}
