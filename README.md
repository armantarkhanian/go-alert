### How to use
```golang
package main

import (
	"fmt"
	"log"

	"github.com/armantarkhanian/go-alert"
)

var (
	botToken = "TELEGRAM_BOT_TOKEN"
	chatIDs  = []int64{
		1, // Telegram chatID where to send message
		2, // Telegram chatID where to send message
		3, // Telegram chatID where to send message
	}

	gmailSMTP = "smtp.gmail.com:587"
	emailFrom = "from@gmail.com"
	emailPass = "password"

	emailRcpts = []string{
		"user1@gmail.com",
		"user2@gmail.com",
	}
)

// Stdout is a custom implementation of alert.AlertManager interface
type Stdout struct{}

func (s *Stdout) Alert(message string) error {
	_, err := fmt.Println(message)
	return err
}

func main() {
	alert := alert.New().
		Telegram(botToken, chatIDs).
		Email(gmailSMTP, emailFrom, emailPass, emailRcpts).
		Custom(&Stdout{})

	if err := alert.Send("WARNING! Something bad happend."); err != nil {
		log.Fatalln(err)
	}
}
```
