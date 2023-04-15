package chatbot

type Update struct {
	message string
	chatID  string
}

type Response struct {
	content string
	chatID  string
}

type Bot interface {
	GetUpdatesChan() <-chan *Update
}
