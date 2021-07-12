package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
)

func main() {
	configs := config.Get()

	tgBot := bot.New(*configs)

	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(cancel)
	go tgBot.Start()
	// wait shutdown event
	<-ctx.Done()
	tgBot.Stop()
}

func gracefulShutdown(cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-sigs:
		log.Println("Receive shutdown signal:", sig)
		cancel()
	}
}
