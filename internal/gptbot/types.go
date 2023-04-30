package gptbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vsetiawan/gptbot/internal/telegram"
)

type tgBot interface {
	GetUpdatesChan() <-chan tgbotapi.Update
	SendMessage(message *telegram.Message) error
}

type chatGPT interface {
	Answer(message string) (string, error)
}
