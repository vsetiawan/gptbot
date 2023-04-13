package chatgpt

import (
	"fmt"
	"net/http"
)

func setDefaultHTTPHeader(req *http.Request, token string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}
