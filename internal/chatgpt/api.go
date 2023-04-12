package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	chatCompletionURL = "https://api.openai.com/v1/chat/completions"
)

func (c *Client) chatCompletionAPI(message string) (*ChatCompletionResponse, error) {
	req, err := c.constructHttpRequest(message)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	chatCompletionResponse, err := toChatCompletionResponse(resp)
	if err != nil {
		return nil, err
	}
	return chatCompletionResponse, nil
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

func toChatCompletionResponse(resp *http.Response) (*ChatCompletionResponse, error) {
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

func closeRequestBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}
