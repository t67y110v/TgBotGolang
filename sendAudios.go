package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendEnAudio(update tgbotapi.Update, bot *tgbotapi.BotAPI, file tgbotapi.FilePath, i int) {

	switch i {
	case 1:
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, EnAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := EngAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = numericKeyboardAudios1
		if _, err := bot.Send(msg2); err != nil {
			panic(err)
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg2 := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg2); err != nil {
				panic(err)
			}
		}

	case 2:
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, EnAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := EngAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = numericKeyboardAudios2
		if _, err := bot.Send(msg2); err != nil {
			panic(err)
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg2 := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg2); err != nil {
				panic(err)
			}
		}
	}
}
func SendGerAudio(update tgbotapi.Update, bot *tgbotapi.BotAPI, file tgbotapi.FilePath, i int) {

	switch i {
	case 1:
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, GerAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := GerAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = numericKeyboardAudios3
		if _, err := bot.Send(msg2); err != nil {
			panic(err)
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg2 := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg2); err != nil {
				panic(err)
			}
		}

	case 2:
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, GerAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := GerAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = numericKeyboardAudios4
		if _, err := bot.Send(msg2); err != nil {
			panic(err)
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg2 := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg2); err != nil {
				panic(err)
			}
		}
	}
}
