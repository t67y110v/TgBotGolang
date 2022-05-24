package textdata

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var GerAudios = map[int]tgbotapi.FilePath{
	1: tgbotapi.FilePath("goBot/repository/audioData/audio3.mp3"),
	2: tgbotapi.FilePath("goBot/repository/audioData/audio4.mp3"),
}
var EnAudios = map[int]tgbotapi.FilePath{
	1: tgbotapi.FilePath("repository/audioData/audio1.mp3"),
	2: tgbotapi.FilePath("repository/audioData/audio2.mp3"),
}
var GerAudiosDescriptions = map[int]string{
	1: "",
	2: "",
}
var EngAudiosDescriptions = map[int]string{
	1: "The TV programme is designed to feature",
	2: "In his course Steven Roberts uses original films with",
}
