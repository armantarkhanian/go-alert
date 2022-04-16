package alert

import (
	"fmt"
	"net/http"
)

type TelegramAlerter struct {
	chatIDs  []int64
	botToken string
}

var _ AlertManager = &TelegramAlerter{}

func (t *TelegramAlerter) Alert(message string) error {
	for _, chatID := range t.chatIDs {
		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", t.botToken, chatID, message)

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		_ = resp.Body.Close()
	}

	return nil
}

func Telegram(botToken string, chatIDs ...int64) *TelegramAlerter {
	return &TelegramAlerter{
		chatIDs:  chatIDs,
		botToken: botToken,
	}
}
