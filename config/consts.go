package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	BotToken          = ""
	YOUTUBE_API_TOKEN = ""
	image1            = tgbotapi.FilePath("")
	image2            = tgbotapi.FilePath("")
	image3            = tgbotapi.FilePath("")
)

var ytChannels = map[int]string{
	1: "https://www.youtube.com/channel/UCAzSfJ8mJ0A8v1oOK2lstPw",
	2: "https://www.youtube.com/channel/UCMTcsanYhBtOb096XegDZQA",
	3: "https://www.youtube.com/channel/UCGywxE07aUcWNAjSjZ1rwhw",
}
