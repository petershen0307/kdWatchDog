package bot

import (
	"fmt"

	tg "gopkg.in/tucnak/telebot.v2"
)

func getEchoHandler(responseCallback func(tg.Recipient, interface{}, ...interface{})) (string, func(*tg.Message)) {
	return tg.OnText, func(m *tg.Message) {
		responseCallback(
			m.Sender,
			fmt.Sprintf("Echo %v", m.Text),
		)
	}
}
