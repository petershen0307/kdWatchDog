package bot

import (
	"log"

	"github.com/petershen0307/kdWatchDog/config"
	tg "gopkg.in/tucnak/telebot.v2"
)

type responseMessage struct {
	to      tg.Recipient
	what    interface{}
	options []interface{}
}

// New a webhook server telegram bot
func New(configs config.Config) *tg.Bot {
	webhook := &tg.Webhook{
		Listen:   ":" + configs.HerokuPort,
		Endpoint: &tg.WebhookEndpoint{PublicURL: configs.HerokuURL},
	}

	setting := tg.Settings{
		Token:  configs.TgToken,
		Poller: webhook,
	}

	b, err := tg.NewBot(setting)
	if err != nil {
		log.Panic(err)
	}
	registerHandlers(b)
	return b
}

func registerHandlers(bot *tg.Bot) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		bot.Send(to, what, options...)
	}
	bot.Handle(getEchoHandler(responseCallback))
}
