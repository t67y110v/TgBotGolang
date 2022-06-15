package main

import (
	"fmt"
	. "goBot/config"
	"goBot/internal/app"
	bDB "goBot/repository/DB/boltdb"
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

	db, err := bDB.InitBoltDB()
	if err != nil {
		log.Fatal(err)
	}

	informRepos := bDB.NewInformRepos(db)

	telegramBot := app.NewBot(bot, informRepos)
	telegramBot.Start()
}
