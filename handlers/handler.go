package handlers

import (
	"context"

	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"go.mongodb.org/mongo-driver/mongo"
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
	userID   int
	fromUser interface{}
	toUser   interface{}
	fromMsg  string
	toMsg    string
}

// Handler is the handler structure, communicate with postman
type Handler struct {
	mailbox   chan Mail
	userColl  *mongo.Collection
	stockColl *mongo.Collection
}

func getHandlerMap(handler *Handler) map[string]func(*Mail) {
	funcMap := map[string]func(*Mail){}
	funcMap[echoCommand] = handler.echo
	funcMap[addCommand] = handler.AddStock
	return funcMap
}

func tgHandlerCallback(commandFunc func(*Mail), m *tg.Message) {
	fromMail := &Mail{
		platform: TelegramBot,
		userID:   m.Sender.ID,
		fromUser: m.Sender,
		toUser:   m.Sender,
		fromMsg:  m.Text,
	}
	commandFunc(fromMail)
}

// NewHandler create a handler
func NewHandler(mailbox chan Mail, configs *config.Config) *Handler {
	// for unit test without DB
	if configs == nil {
		return &Handler{
			mailbox: mailbox,
		}
	}
	return &Handler{
		mailbox:   mailbox,
		userColl:  db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameUsers),
		stockColl: db.GetCollection(configs.MongoDBURI, configs.DBName, db.CollectionNameStocks),
	}
}

// RegisterTelegramBotHandlers register bot handers
func RegisterTelegramBotHandlers(bot *tg.Bot, handler *Handler) {
	funcMap := getHandlerMap(handler)
	for k, v := range funcMap {
		command := k
		if k == echoCommand {
			command = tg.OnText
		}
		bot.Handle(command, func(m *tg.Message) {
			tgHandlerCallback(v, m)
		})
	}
}

// PostmanDeliver will response message to bot client
func PostmanDeliver(ctx context.Context, bot *tg.Bot, mailbox chan Mail) {
	for {
		select {
		case <-ctx.Done():
			return
		case mail := <-mailbox:
			if mail.platform == TelegramBot {
				bot.Send(mail.toUser.(*tg.User), mail.toMsg)
			}
		}
	}
}
