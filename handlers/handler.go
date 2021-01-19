package handlers

import (
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	tg "gopkg.in/tucnak/telebot.v2"
)

type responseCallbackFunc func(tg.Recipient, interface{}, ...interface{})

// RegisterHandlers register bot handers
func RegisterHandlers(bot *tg.Bot, configs *config.Config) {
	responseCallback := func(to tg.Recipient, what interface{}, options ...interface{}) {
		switch object := what.(type) {
		case *string:
			bot.Send(to, *object, options...)
		default:
			bot.Send(to, object, options...)
		}
	}
	userColl := db.GetCollection(configs.MongoDBURI, configs.DBName, "users")
	bot.Handle(getEchoHandler(responseCallback))
	bot.Handle(getAddStockHandler(responseCallback, userColl))
	bot.Handle(getListStockHandler(responseCallback, userColl))
	bot.Handle(getDelStockHandler(responseCallback, userColl))
}
