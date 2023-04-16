package app

import (
	"github.com/vsetiawan/gptbot/internal/chatbot"
	"github.com/vsetiawan/gptbot/internal/chatgpt"
	"os"
)

type gptBot struct {
	chatAgent chatbot.Agent
	chatGPT   *chatgpt.Client
}

func NewGPTBot() *gptBot {
	telegramToken := os.Getenv("HELLO_BOT_TOKEN")
	chatAgent, err := chatbot.NewTelegramBot(telegramToken)
	chatGPTToken := os.Getenv("OPENAI_API_KEY")
	chatGPT := chatgpt.NewClient(chatGPTToken)
	return &gptBot{
		chatAgent: chatAgent,
		chatGPT:   chatGPT,
	}
}

func (g *gptBot) Start() {

}
