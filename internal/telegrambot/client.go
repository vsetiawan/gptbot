package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type TelegramBot struct {
	botAPI      *tgbotapi.BotAPI
	updatesChan <-chan tgbotapi.Update
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
	chatIDInt, err := strconv.ParseInt(response.ChatID, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatIDInt, response.Content)
	_, err = t.botAPI.Send(msg)
	return nil
}
