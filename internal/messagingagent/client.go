package messagingagent

type MessagingAgent interface {
	GetNextMessage() (string, error)
	SendMessage(message string) error
}
