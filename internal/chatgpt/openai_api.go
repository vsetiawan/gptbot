package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	chatCompletionURL = "https://api.openai.com/v1/chat/completions"
)

type chatCompletionAPI struct {
	Token string
}

type chatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Client) chatCompletionAPI(message string) (*ChatCompletionResponse, error) {
	chatCompletionAPI := &chatCompletionAPI{Token: c.Token}
	req, err := chatCompletionAPI.constructHttpRequest(message)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if resp.Status != "200" {
		log.Printf("chat completion API Status Code: %v", resp.Status)
		return nil, errors.New("error calling chat completion API")
	}

	chatCompletionResponse, err := chatCompletionAPI.toChatCompletionResponse(resp)
	if err != nil {
		return nil, err
	}
	return chatCompletionResponse, nil
}

func (c *chatCompletionAPI) constructHttpRequest(message string) (*http.Request, error) {
	payload, err := c.constructPayload(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", chatCompletionURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return nil, err
	}

	setDefaultHTTPHeader(req, c.Token)
	return req, nil
}

func (c *chatCompletionAPI) constructPayload(message string) ([]byte, error) {
	data := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []chatCompletionMessage{
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

func (c *chatCompletionAPI) toChatCompletionResponse(resp *http.Response) (*ChatCompletionResponse, error) {
	defer closeRequestBody(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
		return nil, err
	}
	chatCompletion := &ChatCompletionResponse{}
	err = json.Unmarshal(bodyBytes, chatCompletion)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return chatCompletion, nil
}
