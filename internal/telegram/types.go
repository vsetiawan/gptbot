package telegram

type Update struct {
	Message Message
	ChatID  string
}

type Message struct {
}

type Response struct {
	Content string
	ChatID  int64
}