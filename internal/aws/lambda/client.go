package lambda

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vsetiawan/gptbot/internal/telegram"
	"net/http"
	"os"
)

// handleTelegramWebhookRequest handles the incoming webhook request
func handleTelegramWebhookRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var telegramRequest telegram.TelegramWebhookRequest
	if err := json.Unmarshal([]byte(request.Body), &telegramRequest); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}
	bot, err := telegram.NewBot(os.Getenv("HELLO_BOT_TOKEN"))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, errors.New("internal server error")
	}

	// Reply to the message
	message := fmt.Sprintf("Hi %s, you said: %s", telegramRequest.Message.From.FirstName, telegramRequest.Message.Text)
	err = bot.SendMessage(&telegram.Message{
		Content: message,
		ChatID:  int64(telegramRequest.Message.Chat.ID),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}, errors.New("internal server error")
	}

	response := telegram.TelegramWebhookResponse{
		StatusCode: http.StatusOK,
		Body:       "",
	}
	return events.APIGatewayProxyResponse{
		StatusCode: response.StatusCode, Body: response.Body,
	}, nil
}

func main() {
	ctx, cancelFunc := newDefaultContext()
	defer cancelFunc()
	lambda.StartWithOptions(handleTelegramWebhookRequest, lambda.WithContext(ctx))
}
