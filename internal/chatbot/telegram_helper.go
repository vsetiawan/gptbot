package chatbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func newTelegramUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	return updateConfig
}

func makeUpdatesChan(tgBotAPI *tgbotapi.BotAPI) <-chan *Update {
	updateConfig := newTelegramUpdateConfig()
	tgUpdatesChan := tgBotAPI.GetUpdatesChan(updateConfig)
	updatesChan := make(chan *Update)
	go pipeTelegramChanToUpdateChan(&tgUpdatesChan, updatesChan)
	return updatesChan
}

func pipeTelegramChanToUpdateChan(tgUpdatesChan *tgbotapi.UpdatesChannel, updatesChan chan<- *Update) {
	for telegramUpdate := range *tgUpdatesChan {
		update := &Update{
			message: telegramUpdate.Message.Text,
			chatID:  strconv.FormatInt(telegramUpdate.Message.Chat.ID, 10),
		}
		updatesChan <- update
	}
}
