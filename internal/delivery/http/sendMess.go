package http

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Sm(message *tgbotapi.Message, bot *tgbotapi.BotAPI, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
