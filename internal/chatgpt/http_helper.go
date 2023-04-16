package chatgpt

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func setDefaultHTTPHeader(req *http.Request, token string) {
	log.Printf("token:%v", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}

func closeRequestBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}
