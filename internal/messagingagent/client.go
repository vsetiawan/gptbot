package messagingagent

type MessagingAgent interface {
	GetNextMessage() (string, error)
	Send(message string) error
}
