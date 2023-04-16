package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type TelegramBot struct {
	botAPI      *tgbotapi.BotAPI
	updatesChan <-chan *Update
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	tgBotAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	updatesChan := makeUpdatesChan(tgBotAPI)

	return &TelegramBot{
		botAPI:      tgBotAPI,
		updatesChan: updatesChan,
	}, nil
}

func (t *TelegramBot) GetUpdatesChan() <-chan *Update {
	return t.updatesChan
}

func (t *TelegramBot) SendResponse(response *Response) error {
	chatIDInt, err := strconv.ParseInt(response.chatID, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatIDInt, response.content)
	_, err = t.botAPI.Send(msg)
	return nil
}
