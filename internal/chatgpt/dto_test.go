package chatgpt

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatCompletionResponse_firstChoiceMessageContent(t *testing.T) {
	tests := []struct {
		name   string
		fields *ChatCompletionResponse
		want   string
	}{
		{
			name: "",
			fields: func() *ChatCompletionResponse {
				f := &ChatCompletionResponse{}
				actualContentValue := "content"
				fString := fmt.Sprintf(`{"choices":[{"message":{"content": "%v"}}]}`, actualContentValue)
				_ = json.Unmarshal([]byte(fString), f)
				return f
			}(),
			want: "content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields
			assert.Equalf(t, tt.want, c.firstChoiceMessageContent(), "firstChoiceMessageContent()")
		})
	}
}
