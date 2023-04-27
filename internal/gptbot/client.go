package gptbot

import (
	"context"
	"github.com/vsetiawan/gptbot/internal/chatgpt"
	"github.com/vsetiawan/gptbot/internal/telegrambot"
	"os"
)

type GPTBot struct {
	tgBot   tgBot
	chatGPT chatGPT
}

func NewGPTBot() (*GPTBot, error) {
	tgToken := os.Getenv("HELLO_BOT_TOKEN")
	tgBot, err := telegrambot.NewBot(tgToken)
	if err != nil {
		return nil, err
	}
	chatGPTToken := os.Getenv("OPENAI_API_KEY")
	chatGPT := chatgpt.NewClient(chatGPTToken)
	return &GPTBot{
		tgBot:   tgBot,
		chatGPT: chatGPT,
	}, nil
}

func (g *GPTBot) Start(ctx context.Context) {
	go g.startReceivingUpdates(ctx)
}
