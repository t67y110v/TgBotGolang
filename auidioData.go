package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboardAudios1 = tgbotapi.NewInlineKeyboardMarkup(
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
var numericKeyboardAudios2 = tgbotapi.NewInlineKeyboardMarkup(
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

var numericKeyboardAudios3 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Русский)", "/sendRuAuText3"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Английский)", "/sendGeAuText3"),
	),
)
var numericKeyboardAudios4 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Информация", "https://edu.tatar.ru/chistopol/dou20_mozaika/page1130.htm"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Русский)", "/sendRuAuText4"),
		tgbotapi.NewInlineKeyboardButtonData("Текст(Английский)", "/sendGeAuText4"),
	),
)
var GerAudios = map[int]tgbotapi.FilePath{
	1: tgbotapi.FilePath("pkg/audio3.mp3"),
	2: tgbotapi.FilePath("pkg/audio4.mp3"),
}
var EnAudios = map[int]tgbotapi.FilePath{
	1: tgbotapi.FilePath("pkg/audio1.mp3"),
	2: tgbotapi.FilePath("pkg/audio2.mp3"),
}
var GerAudiosDescriptions = map[int]string{
	1: "",
	2: "",
}
var EngAudiosDescriptions = map[int]string{
	1: "The TV programme is designed to feature",
	2: "In his course Steven Roberts uses original films with",
}
