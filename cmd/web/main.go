package main

import (
	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/db"
	"github.com/petershen0307/kdWatchDog/handlers"
)

func main() {
	configs := config.Get()

	tgBot := bot.New(*configs)

	handlers.RegisterHandlers(tgBot, db.GetCollection(configs.MongoDBURI, configs.DBName, "users"))
	tgBot.Start()
}
