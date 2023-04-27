package gptbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vsetiawan/gptbot/internal/telegrambot"
)

type tgBot interface {
	GetUpdatesChan() <-chan tgbotapi.Update
	SendResponse(response *telegrambot.Response) error
}

type chatGPT interface {
	Answer(message string) (string, error)
}
