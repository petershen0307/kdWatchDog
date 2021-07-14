package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/petershen0307/kdWatchDog/bot"
	"github.com/petershen0307/kdWatchDog/config"
	"github.com/petershen0307/kdWatchDog/handlers"
)

func main() {
	configs := config.Get()
	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(cancel)
	startBot(ctx, configs)
}

func gracefulShutdown(cancel context.CancelFunc) {
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	// wait shutdown event
	sig := <-sigs
	log.Println("Receive shutdown signal:", sig)

}

func startBot(ctx context.Context, configs *config.Config) {
	tgBot := bot.New(*configs)
	mailbox := make(chan handlers.Mail, 10)
	handlers.RegisterTelegramBotHandlers(tgBot, handlers.NewHandler(mailbox, configs))
	go handlers.PostmanDeliver(ctx, tgBot, mailbox)
	go tgBot.Start()
	// wait shutdown event
	<-ctx.Done()
	tgBot.Stop()
}
