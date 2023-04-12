package chatgpt

import (
	"os"
)

type Client struct {
	Token string
}

func NewClient() *Client {
	return &Client{
		Token: os.Getenv("OPENAI_API_KEY"), // TODO: move this to be done on app start
	}
}

type requestBody struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Client) Answer(message string) (string, error) {
	resp, err := c.chatCompletionAPI(message)
	if err != nil {
		return "", err
	}
	return resp.getAnswer(), nil
}
