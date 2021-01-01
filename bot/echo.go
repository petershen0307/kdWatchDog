package bot

import tg "gopkg.in/tucnak/telebot.v2"

func getEchoHandler(bot *tg.Bot) func(m *tg.Message) {
	return func(m *tg.Message) {
		bot.Send(m.Sender, "Echo "+m.Text)
	}
}
