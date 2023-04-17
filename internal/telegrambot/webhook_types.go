package telegrambot

// TelegramWebhookRequest is the request structure sent by Telegram to the webhook URL
type TelegramWebhookRequest struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			Type      string `json:"type"`
			Title     string `json:"title"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

// TelegramWebhookResponse is the response structure sent back to Telegram
type TelegramWebhookResponse struct {
	StatusCode int    `json:"-"`
	Body       string `json:"body"`
}
