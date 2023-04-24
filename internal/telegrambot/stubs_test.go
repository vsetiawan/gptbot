package telegrambot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type botAPIStub struct{}

func (b *botAPIStub) GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	c := make(chan tgbotapi.Update)
	return c
}

func (b *botAPIStub) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return tgbotapi.Message{}, nil
}
