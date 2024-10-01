package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	dotenv "github.com/joho/godotenv"
)

const token = "TGBOT_TOKEN"

func init() {
	dotenv.Load()
}

func main() {
	secretToken, ok := os.LookupEnv(token)
	if !ok {
		log.Fatalf("environment variable [%s] required", token)
	}

	bot, err := tgbotapi.NewBotAPI(secretToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConf := tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(updateConf)
	if err != nil {
		log.Panic(err)
	}

	for u := range updates {
		if u.Message == nil {
			continue
		}

		log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)

		msg := tgbotapi.NewMessage(u.Message.Chat.ID, "received text: "+u.Message.Text)
		// msg.ReplyToMessageID = u.Message.MessageID

		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
