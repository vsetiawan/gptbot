package chatgpt

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_setDefaultHTTPHeader(t *testing.T) {
	type args struct {
		req   *http.Request
		token string
	}
	tests := []struct {
		name           string
		args           args
		expectedHeader *http.Request
	}{
		{
			name: "Headers are set correctly",
			args: args{
				req: &http.Request{
					Header: map[string][]string{},
				},
				token: "token",
			},
			expectedHeader: &http.Request{
				Header: map[string][]string{
					"Authorization": {"Bearer token"},
					"Content-Type":  {"application/json"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRequest := tt.args.req
			setDefaultHTTPHeader(tt.args.req, tt.args.token)
			assert.Equal(t, tt.expectedHeader, actualRequest)
		})
	}
}

type bodyStub struct {
	closed bool
}

func (b *bodyStub) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (b *bodyStub) Close() error {
	b.closed = true
	return nil
}

func Test_closeRequestBody(t *testing.T) {
	type args struct {
		Body *bodyStub
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "close request body successfully",
			args: args{
				Body: &bodyStub{
					closed: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualBody := tt.args.Body
			closeRequestBody(tt.args.Body)
			assert.True(t, actualBody.closed)
		})
	}
}
