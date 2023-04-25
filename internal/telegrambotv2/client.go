package telegrambotv2

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	botAPI      TelegramBotAPI
	updatesChan <-chan tgbotapi.Update
}

type TelegramBotAPI interface {
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

func (t Bot) GetNextMessage() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t Bot) SendMessage(message string) error {
	//TODO implement me
	panic("implement me")
}
