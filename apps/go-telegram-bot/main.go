package main

import (
	"fmt"
	"go-telegram-bot/services"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("BOT_TOKEN")
	fmt.Println(token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized  account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, " kamu mengirim chat: "+update.Message.Text)

			if update.Message.IsCommand() {
				handleCommand(bot, update.Message)
			}
			// optional jika menamah bagian ini maka akan mereply message
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "weather":
		weatherCommand(bot, message)
	default:
		weatherCommand(bot, message)
	}
}

func weatherCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	city := "Denpasar"
	weater, err := services.GetWeather(city)
	if err != nil {
		log.Fatal(err)
	}

	stringMessage := fmt.Sprintf("cuaca kota %s adalah %v celcius", weater.Name, weater.Main.Temp)
	msg := tgbotapi.NewMessage(message.Chat.ID, stringMessage)
	bot.Send(msg)
}
