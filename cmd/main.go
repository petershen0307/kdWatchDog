package main

import (
	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
)

func main() {
	configs := config.Get()

	tgBot := bot.New(*configs)
	tgBot.Start()
}
