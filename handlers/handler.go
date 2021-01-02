package handlers

import (
	tg "gopkg.in/tucnak/telebot.v2"
)

// RegisterHandlers register bot handers
func RegisterHandlers(bot *tg.Bot) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		bot.Send(to, what, options...)
	}
	bot.Handle(getEchoHandler(responseCallback))
}
