package telegrambotv2

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func makeTgUpdatesChan(botAPI TelegramBotAPI) <-chan tgbotapi.Update {
	updateConfig := newTelegramUpdateConfig()
	tgUpdatesChan := botAPI.GetUpdatesChan(updateConfig)
	return tgUpdatesChan
}

func newTelegramUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(telegramUpdateOffset)
	updateConfig.Timeout = telegramUpdateTimeout
	return updateConfig
}
