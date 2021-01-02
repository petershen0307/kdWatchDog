package bot

import (
	"fmt"

	tg "gopkg.in/tucnak/telebot.v2"
)

func getEchoHandler(bot *tg.Bot) (string, func(*tg.Message)) {
	return tg.OnText, func(m *tg.Message) {
		bot.Send(m.Sender, echo(m.Text))
	}
}

func echo(msg string) string {
	return fmt.Sprintf("Echo %v", msg)
}
