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
		case *tg.Photo:
			bot.SendAlbum(p.to, tg.Album{p.what.(*tg.Photo)})
		}
	}
	userColl := db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameUsers)
	stockColl := db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameStocks)
	bot.Handle(getEchoHandler(responseCallback))
	bot.Handle(getAddStockHandler(responseCallback, userColl))
	bot.Handle(getListStockHandler(responseCallback, userColl))
	bot.Handle(getDelStockHandler(responseCallback, userColl))
	bot.Handle(getQueryStockHandler(responseCallback, userColl, stockColl, configs.ImgurClientID))
}

// defined bot platform
type BotPlatform int

// TelegramBot is a bot platform
const TelegramBot BotPlatform = 1

// Mail is the structure for handler and bot
type Mail struct {
	platform BotPlatform
	fromID   interface{}
	toID     interface{}
	fromMsg  string
	toMsg    string
}

// Handler is the handler structure, communicate with postman
type Handler struct {
	mailBox chan Mail
}

type handlerBroker struct {
	handlerMap map[string]func(*Mail)
}

func newHandlerBroker(handler *Handler) *handlerBroker {
	broker := handlerBroker{}
	broker.handlerMap[echoCommand] = handler.echo
	return &broker
}

func tgHandlerCallback(broker *handlerBroker, command string, m *tg.Message) {
	fromMail := &Mail{
		platform: TelegramBot,
		fromID:   m.Sender,
		toID:     m.Sender,
		fromMsg:  m.Text}
	if command == tg.OnText {
		// echo
		broker.handlerMap[echoCommand](fromMail)
	}
}
