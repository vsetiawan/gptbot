package chatgpt

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func setDefaultHTTPHeader(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}

func closeRequestBody(Body io.ReadCloser) {
	if err := Body.Close(); err != nil {
		log.Printf("An error occurred: %s", err.Error())
	}
}
