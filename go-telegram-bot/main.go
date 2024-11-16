package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Debug = false
	Token = ""
)

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	updates := bot.GetUpdatesChan(u)

loop:
	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			if !update.Message.IsCommand() {
				continue
			}
			log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case <-quitCh:
			log.Println("shutting down gracefully")
			close(quitCh)
			break loop
		}
	}
}
