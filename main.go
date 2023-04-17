package main

import (
	"bufio"
	"context"
	"github.com/vsetiawan/gptbot/internal/gptbot"
	"log"
	"os"
)

func main() {
	cancelFunc := startBot()
	log.Println("Start listening for updates. Press enter to stop")
	waitForNewlineAndCancelContext(cancelFunc)
}

func waitForNewlineAndCancelContext(cancelFunc context.CancelFunc) {
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Println(err.Error())
	}
	cancelFunc()
}

func startBot() context.CancelFunc {
	bot, err := gptbot.NewGPTBot()
	if err != nil {
		panic("Error creating GPTBot")
	}
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	bot.Start(ctx)
	return cancelFunc
}
