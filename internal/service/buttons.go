package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NumericKeyboardOpe = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔎Команды"),
		tgbotapi.NewKeyboardButton("📚Информация"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📈Статистика"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📝Английский язык"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📝Немецкие язык"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👥Поиск собеседника"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🙅‍♂️Прекратить общение"),
	),
)
var NumericKeyboardEng = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📝Английские задания"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🎧Английская речь"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🌍В меню"),
	),
)
var NumericKeyboardGer = tgbotapi.NewReplyKeyboard(

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📝Немецкие задания"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🎧Немецкая речь"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🌍В меню"),
	),
)
