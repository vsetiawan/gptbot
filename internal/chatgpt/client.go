package chatgpt

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) Answer(message string) (string, error) {
	resp, err := c.chatCompletionAPI(message)
	if err != nil {
		return "", err
	}
	return resp.getAnswer(), nil
}
