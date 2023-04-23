package chatgpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "client created successfully",
			args: args{
				token: "token",
			},
			want: &Client{
				Token: "token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewClient(tt.args.token), "NewClient(%v)", tt.args.token)
		})
	}
}
