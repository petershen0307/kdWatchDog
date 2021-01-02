package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	tg "gopkg.in/tucnak/telebot.v2"
)

type responseCallbackFunc func(tg.Recipient, interface{}, ...interface{})
type botHandler func(*tg.Message)

// RegisterHandlers register bot handers
func RegisterHandlers(bot *tg.Bot, collection *mongo.Collection) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		switch object := what.(type) {
		case *string:
			bot.Send(to, *object, options...)
		default:
			bot.Send(to, object, options...)
		}
	}
	bot.Handle(getEchoHandler(responseCallback))
	bot.Handle(getAddStockHandler(responseCallback, collection))
	bot.Handle(getListStockHandler(responseCallback, collection))
}
