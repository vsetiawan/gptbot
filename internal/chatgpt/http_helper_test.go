package chatgpt

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
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
	closed    bool
	returnErr bool
}

func (b *bodyStub) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (b *bodyStub) Close() error {
	if b.returnErr {
		return errors.New("error")
	}
	b.closed = true
	return nil
}

func Test_closeRequestBody(t *testing.T) {
	type args struct {
		Body *bodyStub
	}
	origOutput := log.Writer()
	defer log.SetOutput(origOutput)
	tests := []struct {
		name         string
		args         args
		buf          *bytes.Buffer
		expectClosed bool
		expectLog    bool
	}{
		{
			name: "close request body successfully",
			args: args{
				Body: &bodyStub{
					returnErr: false,
				},
			},
			expectClosed: true,
			expectLog:    false,
		},
		{
			name: "close request body error",
			args: args{
				Body: &bodyStub{
					returnErr: true,
				},
			},
			buf:          &bytes.Buffer{},
			expectClosed: false,
			expectLog:    true,
		},
	}
	for _, tt := range tests {
		if tt.expectLog {
			log.SetOutput(tt.buf)
		}
		t.Run(tt.name, func(t *testing.T) {
			actualBody := tt.args.Body
			closeRequestBody(tt.args.Body)
			assert.Equal(t, tt.expectClosed, actualBody.closed)
			if tt.expectLog {
				got := tt.buf.String()
				assert.NotEqual(t, "<nil>", got)
			}
		})
	}
}
