package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vsetiawan/gptbot/internal/telegrambot"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handleTelegramWebhookRequest handles the incoming webhook request
func handleTelegramWebhookRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var telegramRequest telegrambot.TelegramWebhookRequest
	if err := json.Unmarshal([]byte(request.Body), &telegramRequest); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	// Reply to the message
	message := fmt.Sprintf("Hi %s, you said: %s", telegramRequest.Message.From.FirstName, telegramRequest.Message.Text)
	reply := map[string]string{"method": "sendMessage", "chat_id": fmt.Sprintf("%d", telegramRequest.Message.Chat.ID), "text": message}
	replyJSON, err := json.Marshal(reply)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	// Send the reply back to Telegram
	response := telegrambot.TelegramWebhookResponse{StatusCode: http.StatusOK, Body: string(replyJSON)}
	return events.APIGatewayProxyResponse{StatusCode: response.StatusCode, Body: response.Body}, nil
}

func main() {
	lambda.Start(handleTelegramWebhookRequest)
}
