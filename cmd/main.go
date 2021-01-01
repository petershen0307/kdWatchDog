package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		herokuPort = os.Getenv("PORT")
		herokuURL  = os.Getenv("HEROKU_URL")
		tgToken    = os.Getenv("TG_TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + herokuPort,
		Endpoint: &tb.WebhookEndpoint{PublicURL: herokuURL},
	}

	pref := tb.Settings{
		Token:  tgToken,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "Echo "+m.Text)
	})

	b.Start()
}
