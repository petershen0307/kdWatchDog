package main

import (
	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/handlers"
)

func main() {
	configs := config.Get()

	tgBot := bot.New(*configs)
	handlers.RegisterHandlers(tgBot)
	tgBot.Start()
}
