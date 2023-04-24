package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_newTelegramUpdateConfig(t *testing.T) {
	tests := []struct {
		name string
		want tgbotapi.UpdateConfig
	}{
		{
			name: "created update config successfully",
			want: tgbotapi.UpdateConfig{
				Offset:         telegramUpdateOffset,
				Limit:          0,
				Timeout:        telegramUpdateTimeout,
				AllowedUpdates: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTelegramUpdateConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTelegramUpdateConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeTgUpdatesChan(t *testing.T) {
	type args struct {
		botAPI BotAPI
	}
	tests := []struct {
		name string
		args args
		want <-chan tgbotapi.Update
	}{
		{
			name: "Successfully created TG Updates Channel",
			args: args{
				botAPI: &botAPIStub{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeTgUpdatesChan(tt.args.botAPI)
			assert.NotNil(t, got)
		})
	}
}
