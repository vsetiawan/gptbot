package chatbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
	updatesChan <-chan *Update
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	tgBotAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	updatesChan := makeUpdatesChan(tgBotAPI)

	return &TelegramBot{
		BotAPI:      tgBotAPI,
		updatesChan: updatesChan,
	}, nil
}

func (t *TelegramBot) GetUpdatesChan() <-chan *Update {
	return t.updatesChan
}
