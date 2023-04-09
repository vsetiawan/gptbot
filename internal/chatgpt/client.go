package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Client struct {
	Token string
}

type Sender interface {
	Send(message string) (string, error)
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

func (c *Client) Send(message string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	data := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []requestBody{
			{
				Role:    "user",
				Content: message,
			},
		},
		//"temperature": 1,
		//"n":           1,
		//"stop":        "\n",
	}
	payload, err := json.Marshal(data)
	log.Print(string(payload))
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("An error occured: %s", err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return "", err
	}
	chatCompletion := &ChatCompletion{}
	err = json.Unmarshal(body, chatCompletion)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf(chatCompletion.Choices[0].Message.Content)
	return chatCompletion.Choices[0].Message.Content, nil
}
