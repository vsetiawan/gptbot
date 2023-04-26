package messagingagent

type Message interface {
	GetContent() string
	GetRecipient() string
}
