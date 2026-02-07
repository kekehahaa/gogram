package gogram

import (
	"fmt"

	tgclient "github.com/kekehahaa/gogram/clients/telegram"
	event_consumer "github.com/kekehahaa/gogram/consumer/event-consumer"
	"github.com/kekehahaa/gogram/events/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

type Bot struct {
	token     string
	batchSize int
}

func New(token string, batchSize int) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token is nil")
	}

	return &Bot{
		token:     token,
		batchSize: batchSize,
	}, nil
}

func (b *Bot) StartTelegramBot() error {

	eventsProcessor := telegram.New(tgclient.New(tgBotHost, b.token))

	// log

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, b.batchSize)
	if err := consumer.Start(); err != nil {
		return fmt.Errorf("token is nil")
	}

	return nil
}
