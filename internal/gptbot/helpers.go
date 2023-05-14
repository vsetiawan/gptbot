package gptbot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vsetiawan/gptbot/internal/telegram"
	"log"
)

func (g *GPTBot) startReceivingUpdates(ctx context.Context) {
	updates := g.tgBot.GetUpdatesChan()
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			g.handleUpdate(ctx, update)
		}
	}
}

func (g *GPTBot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		g.handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		g.handleButton(update.CallbackQuery)
		break
	}
}

func (g *GPTBot) handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}
	log.Printf("%s wrote: %s", user.FirstName, text)
	resp, err := g.chatGPT.Answer(text)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		tgbotapi.NewMessage(message.Chat.ID, "ChatGPT is busy right now. :(")
	}
	log.Printf("%s wrote: %s", "chatgpt", resp)
	tgMessage := &telegram.Message{
		Content: resp,
		ChatID:  message.Chat.ID,
	}
	if err := g.tgBot.SendMessage(tgMessage); err != nil {
		log.Printf("An error occurred: %s", err.Error())
	}
}

func (g *GPTBot) handleButton(query *tgbotapi.CallbackQuery) {}

func (g *GPTBot) handleCommand() {}
