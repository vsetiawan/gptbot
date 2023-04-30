package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	botAPI      botAPI
	updatesChan <-chan tgbotapi.Update
}

type botAPI interface {
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

func NewBot(token string) (*Bot, error) {
	tgBotAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	updatesChan := makeTgUpdatesChan(tgBotAPI)

	return &Bot{
		botAPI:      tgBotAPI,
		updatesChan: updatesChan,
	}, nil
}

func (b *Bot) GetUpdatesChan() <-chan tgbotapi.Update {
	return b.updatesChan
}

func (b *Bot) SendMessage(message *Message) error {
	msg := tgbotapi.NewMessage(message.ChatID, message.Content)
	_, err := b.botAPI.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
