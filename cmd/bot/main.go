package main

import (
	"fmt"
	. "goBot/config"
	"goBot/internal/app"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Println("Bot is up!")
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	telegramBot := app.NewBot(bot)
	telegramBot.Start()
}
