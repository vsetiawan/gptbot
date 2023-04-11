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

const (
	chatCompletionURL = "https://api.openai.com/v1/chat/completions"
)

type Client struct {
	Token string
}

type ResponseGetter interface {
	GetResponse(message string) (string, error)
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

func (c *Client) GetResponse(message string) (string, error) {
	req, err := c.constructHttpRequest(message)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer closeRequestBody(resp.Body)

	messageResponse, err := getMessageContentFromResponseBody(resp.Body)
	if err != nil {
		return "", err
	}
	log.Print(messageResponse)
	return messageResponse, nil
}

func (c *Client) constructHttpRequest(message string) (*http.Request, error) {
	payload, err := constructPayload(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", chatCompletionURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))
	return req, nil
}

func constructPayload(message string) ([]byte, error) {
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
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return nil, err
	}
	return payload, err
}

func closeRequestBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func getMessageContentFromResponseBody(body io.ReadCloser) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return "", err
	}
	chatCompletion := &ChatCompletion{}
	err = json.Unmarshal(bodyBytes, chatCompletion)
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}
	return chatCompletion.GetFirstMessageContent(), nil
}
