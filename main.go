package main

import (
	"bufio"
	"context"
	"github.com/vsetiawan/gptbot/internal/gptbot"
	"log"
	"os"
)

func main() {
	bot, err := gptbot.NewGPTBot()
	if err != nil {
		panic("Error creating GPTBot")
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	bot.Start(ctx)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Println(err.Error())
	}
	cancel()
}
