package handlers

import (
	"fmt"

	tg "gopkg.in/tucnak/telebot.v2"
)

func getEchoHandler(responseCallback responseCallbackFunc) (string, func(*tg.Message)) {
	return tg.OnText, func(m *tg.Message) {
		responseCallback(
			&post{
				to:   m.Sender,
				what: fmt.Sprintf("Echo %v", m.Text),
			})
	}
}
