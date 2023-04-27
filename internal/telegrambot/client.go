package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	botAPI      BotAPI
	updatesChan <-chan tgbotapi.Update
}

type BotAPI interface {
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

func NewBot(token string) (*TelegramBot, error) {
	tgBotAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	updatesChan := makeTgUpdatesChan(tgBotAPI)

	return &TelegramBot{
		botAPI:      tgBotAPI,
		updatesChan: updatesChan,
	}, nil
}

func (t *TelegramBot) GetUpdatesChan() <-chan tgbotapi.Update {
	return t.updatesChan
}

func (t *TelegramBot) SendResponse(response *Response) error {
	msg := tgbotapi.NewMessage(response.ChatID, response.Content)
	_, err := t.botAPI.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
