package http

import (
	rep "goBot/repository/audiodata"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var pseudoRand = 2

func SendAu(update tgbotapi.Update, b *tgbotapi.BotAPI) {

	if pseudoRand == 0 {
		pseudoRand = 2
	}
	SendEnAudio(update, b, rep.EnAudios[pseudoRand], pseudoRand)
	pseudoRand--
}

func SendEnAudio(update tgbotapi.Update, bot *tgbotapi.BotAPI, file tgbotapi.FilePath, i int) {

	switch i {
	case 1:
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, rep.EnAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := rep.EngAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = NumericKeyboardAudios1
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
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, rep.EnAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := rep.EngAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = NumericKeyboardAudios2
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
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, rep.GerAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := rep.GerAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = NumericKeyboardAudios3
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
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, rep.GerAudios[i])
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		tx := rep.GerAudiosDescriptions[i]
		msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, tx)
		msg2.ReplyMarkup = NumericKeyboardAudios4
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

var NumericKeyboardAudios1 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://en-ege.sdamgia.ru/problem?id=1489"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("actors pretending to be ordinary people", "❌Неправильный ответ"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("people who vote for themselves to win a prize", "❌Неправильный ответ"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("real people preparing dinner parties in their own homes", "✅Правильный ответ!"),
	),
)
var NumericKeyboardAudios2 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://en-ege.sdamgia.ru/problem?id=753"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("original subtitles", "✅Правильный ответ!"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("no subtitles", "❌Неправильный ответ"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("English subtitles", "❌Неправильный ответ"),
	),
)

var NumericKeyboardAudios3 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Русский)", "/sendRuAuText3"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Английский)", "/sendGeAuText3"),
	),
)
var NumericKeyboardAudios4 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Русский)", "/sendRuAuText4"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Английский)", "/sendGeAuText4"),
	),
)
