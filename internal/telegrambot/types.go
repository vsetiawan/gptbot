package telegrambot

type Update struct {
	Message Message
	ChatID  string
}

type Message struct {
}

type Response struct {
	Content string
	ChatID  string
}
