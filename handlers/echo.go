package handlers

import (
	"fmt"
)

const echoCommand = "/echo"

func (handle *Handler) echo(mail *Mail) {
	mail.toMsg = fmt.Sprintf("Echo %v", mail.fromMsg)
	handle.mailbox <- *mail
}
