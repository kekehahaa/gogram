package gogram

import (
	"log"
	"xrayBot/internal/lib/grpc/gogram/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func NewTelegramBot(token string) {
	if token == "" {
		log.Fatal("token is nil")
	}

	tgClient := telegram.New(tgBotHost, token)
}
