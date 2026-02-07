package gogram

import (
	"fmt"
	"log/slog"

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
	log       slog.Logger
}

func New(token string, batchSize int, logger slog.Logger) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token is nil")
	}

	return &Bot{
		token:     token,
		batchSize: batchSize,
		log:       logger,
	}, nil
}

func (b *Bot) StartTelegramBot() error {

	eventsProcessor := telegram.New(tgclient.New(tgBotHost, b.token))

	b.log.Info("Starting bot")

	consumer := event_consumer.New(b.log, eventsProcessor, eventsProcessor, b.batchSize)
	if err := consumer.Start(); err != nil {
		return fmt.Errorf("token is nil")
	}

	return nil
}
