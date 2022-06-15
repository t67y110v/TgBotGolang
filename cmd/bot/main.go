package main

import (
	"fmt"
	. "goBot/config"
	"goBot/internal/app"
	bDB "goBot/repository/DB"
	"goBot/repository/DB/boltdb"
	"log"

	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Println("Bot is up!")
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	informRepos := boltdb.NewInformRepos(db)

	telegramBot := app.NewBot(bot, informRepos)
	telegramBot.Start()
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bDB.Information))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
