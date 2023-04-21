package chatgpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatCompletionResponse_firstChoiceMessageContent(t *testing.T) {
	type fields struct {
		ID      string
		Object  string
		Created int64
		Model   string
		Usage   struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		}
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
			FinishReason string `json:"finish_reason"`
			Index        int    `json:"index"`
		}
	}
	tests := []struct {
		name   string
		fields *fields
		want   string
	}{
		{
			name: "",
			fields: func() *fields {
				f := &fields{}
				f.Choices[0].Message.Content = "content"
				return f
			}(),
			want: "content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatCompletionResponse{
				ID:      tt.fields.ID,
				Object:  tt.fields.Object,
				Created: tt.fields.Created,
				Model:   tt.fields.Model,
				Usage:   tt.fields.Usage,
				Choices: tt.fields.Choices,
			}
			assert.Equalf(t, tt.want, c.firstChoiceMessageContent(), "firstChoiceMessageContent()")
		})
	}
}
