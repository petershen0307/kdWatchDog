package handlers

import (
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	tg "gopkg.in/tucnak/telebot.v2"
)

type responseCallbackFunc func(*post)

type post struct {
	to      tg.Recipient
	what    interface{}
	options []interface{}
}

// RegisterHandlers register bot handers
func RegisterHandlers(bot *tg.Bot, configs *config.Config) {
	responseCallback := func(p *post) {
		switch p.what.(type) {
		case string:
			bot.Send(p.to, p.what, p.options...)
		case tg.Photo:
			bot.SendAlbum(p.to, tg.Album{p.what.(*tg.Photo)})
		}
	}
	userColl := db.GetCollection(configs.MongoDBURI, configs.DBName, "users")
	stockColl := db.GetCollection(configs.MongoDBURI, configs.DBName, "stocks")
	bot.Handle(getEchoHandler(responseCallback))
	bot.Handle(getAddStockHandler(responseCallback, userColl))
	bot.Handle(getListStockHandler(responseCallback, userColl))
	bot.Handle(getDelStockHandler(responseCallback, userColl))
	bot.Handle(getQueryStockHandler(responseCallback, userColl, stockColl))
}
