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
		tgbotapi.NewInlineKeyboardButtonURL("????????????????????", "https://en-ege.sdamgia.ru/problem?id=1489"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("actors pretending to be ordinary people", "??????????????????????????? ??????????"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("people who vote for themselves to win a prize", "??????????????????????????? ??????????"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("real people preparing dinner parties in their own homes", "??????????????????????? ??????????!"),
	),
)
var NumericKeyboardAudios2 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("????????????????????", "https://en-ege.sdamgia.ru/problem?id=753"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("original subtitles", "??????????????????????? ??????????!"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("no subtitles", "??????????????????????????? ??????????"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("English subtitles", "??????????????????????????? ??????????"),
	),
)

var NumericKeyboardAudios3 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("????????????????????", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("??????????(??????????????)", "/sendRuAuText3"),
		tgbotapi.NewInlineKeyboardButtonData("??????????(????????????????????)", "/sendGeAuText3"),
	),
)
var NumericKeyboardAudios4 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("????????????????????", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("??????????(??????????????)", "/sendRuAuText4"),
		tgbotapi.NewInlineKeyboardButtonData("??????????(????????????????????)", "/sendGeAuText4"),
	),
)
