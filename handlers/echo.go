package handlers

import (
	"fmt"

	tg "gopkg.in/tucnak/telebot.v2"
)

const echoCommand = "/echo"

func getEchoHandler(responseCallback responseCallbackFunc) (string, func(*tg.Message)) {
	return tg.OnText, func(m *tg.Message) {
		responseCallback(
			&post{
				to:   m.Sender,
				what: fmt.Sprintf("Echo %v", m.Text),
			})
	}
}

func (self *Handler) echo(mail *Mail) {
	mail.toMsg = fmt.Sprintf("Echo %v", mail.fromMsg)
	self.mailBox <- *mail
}
