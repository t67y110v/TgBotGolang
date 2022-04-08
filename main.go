package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
	YOUTUBE_API_TOKEN  = "YT token"
	YOUTUBE_VIDEO_URL  = "https://www.youtube.com/watch?v="
	botToken           = "TG token"
	botApi             = "https://api.telegram.org/bot"
	RightAnsw          = "Да! Все верно!\nДля продолжения выполенения заданий используй комманду /tasksList"
	WrongAnsw          = "Ответ неправльный, попробуй еще"
)

//точка входа программы
func main() {
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			log.Println("Smth went wrong", err.Error())
		}
		fmt.Println(updates)
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}
	}
}

// Запрос обновлений
func GetUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getupdates" + "?offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

//отвечать на обновления
func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	if botMessage.Text == "/start" {
		fmt.Println("зашел  в старт")
		botMessage.Text = "Привет! \n Tы используешь интерактивного бота для тренировки навыка переводчика!! \n Список доступных команд /ENGtasksList - список модулей, и заданий содержащихся в них\n /ENGmoduleInformation -  список справочной информации необходимой для успешного  решения задач из  модуля \n "
	} else if strings.ToLower(botMessage.Text) == "/commands" {
		fmt.Println("зашел  в список команд  ")
		botMessage.Text = "/ENGtasksList - список модулей, и заданий содержащихся в них\n /ENGmoduleInformation -  список справочной информации необходимой для успешного  решения задач из  модуля \n "
	} else if strings.ToLower(botMessage.Text) == "/ENGtasksList" || strings.ToLower(botMessage.Text) == "/engtasklist" {
		fmt.Println("зашел  в таксклист")
		botMessage.Text = "Задания первого модуля для английского языка: \n  /en_m1_task1  /en_m1_task2  /en_m1_task3  \n Задания второго модуля \n /en_m2_task1  /en_m2_task2  /en_m2_task3 /en_m2_task4 \n Задания третьего модуля /en_m3_task1  /en_m3_task2  /en_m3_task3\n Задания четвертого  модуля \n /en_m4_task1  /en_m4_task2  /en_m4_task3 \n Задания пятого модуля /en_m5_task1  /en_m5_task2  /en_m5_task3\n "
	} else if strings.ToLower(botMessage.Text) == "/en_m2_task1" {
		botMessage.Text = "Модуль 2\n Задание 1 \n Напишите перевод  выделенного /***/ слова в заданном контексте \n 1)We put up our tents and made a /fire/."
	} else if strings.ToLower(botMessage.Text) == "Костер" || strings.ToLower(botMessage.Text) == "костер" || strings.ToLower(botMessage.Text) == "Костёр" || strings.ToLower(botMessage.Text) == "костёр" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m2_task2" {
		botMessage.Text = "Модуль 2 \n Задание 2 \nНапишите перевод  выделенного /***/ слова в заданном контексте \n А secretary was fired after she was caught stealing from the company"
	} else if strings.ToLower(botMessage.Text) == "сожгли" || strings.ToLower(botMessage.Text) == "Cожгли" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m2_task3" {
		botMessage.Text = "Модуль 2 \n Задание 3 \n Напишите перевод  выделенного /***/ слова в заданном контексте \n  Nick is too /mean/ to buy her a ring."
	} else if strings.ToLower(botMessage.Text) == "скуп" || strings.ToLower(botMessage.Text) == "Скуп" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m2_task4" {
		botMessage.Text = "Модуль 2 \n Задание 4 \n Напишите перевод  выделенного /***/ слова в заданном контексте  \n She is a /mean/ piano player."
	} else if strings.ToLower(botMessage.Text) == "подлая" || strings.ToLower(botMessage.Text) == "Подлая" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m3_task1" {
		botMessage.Text = "Модуль 3 \n Задание 1\nПопробуйте перевести данные слова одним словом\nFried eggs sunny side up"
	} else if strings.ToLower(botMessage.Text) == "глазунья" || strings.ToLower(botMessage.Text) == "яичница" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m3_task2" {
		botMessage.Text = "Модуль 3 \n Задание 2 \n Попробуйте перевести данные слова одним словом\nBorder patrol agents"
	} else if strings.ToLower(botMessage.Text) == "пограничники" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m3_task3" {
		botMessage.Text = "Модуль 3 \n Задание 3\nПопробуйте перевести данные слова одним словом\nWhite blood cells"
	} else if strings.ToLower(botMessage.Text) == "лейкоциты" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task1" {
		botMessage.Text = "Модуль 1 \n Задание 1 \n Напишите перевод данного предложения: \nWe were sure that Jack would win the competition he had been preparing for"
	} else if strings.ToLower(botMessage.Text) == "мы были уверены, что джек выиграет эти соревнования и что он станет новым чемпионом" || strings.ToLower(botMessage.Text) == "мы были уверены, что джек выиграет этот конкурс и что он станет новым чемпионом" || strings.ToLower(botMessage.Text) == "мы были уверены, что джек выиграет эти соревнования и станет новым чемпионом" || strings.ToLower(botMessage.Text) == "Мы были уверены, что Джек выиграет соревнования, к которым он готовился" || strings.ToLower(botMessage.Text) == "мы были уверены, что жек выиграет соревнование, к которому он готовился" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task2" {
		botMessage.Text = "Модуль 1 \n Задание 2 \n Напишите перевод данного предложения:\n Mr. Willson, who stood next to us, suggested going outside"
	} else if strings.ToLower(botMessage.Text) == "мистер уилсон, который стоял рядом с нами, предложил выйти на улицу" || strings.ToLower(botMessage.Text) == "мистер уилсон, который стоял около нас, предложил выйти на улицу" || strings.ToLower(botMessage.Text) == "мистер уилсон, стоявший рядом с нами, предложил выйти на улицу" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task3" {
		botMessage.Text = "Модуль 1 \n Задание 3\n Напишите перевод данного предложения: \n I don’t like that kind of situations when you have to talk to someone but don’t remember his name"
	} else if strings.ToLower(botMessage.Text) == "я не люблю такие ситуации, когда ты должен поговорить с кем то, но не помнишь его имя" || botMessage.Text == "я не люблю такие ситуации, когда ты должен поговорить с кем то, но ты не помнишь его имя" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m4_task1" {
		botMessage.Text = "Модуль 4 \n Задание 1 \n Прочитайте предложения, вместо пропуска добавьте слово или фразу, которое подошло бы по значению:\n Melanie River had just turned ten years old, but she had never been to school. Most children in the third grade ___ to school for two years. But Melanie was not most children"
	} else if strings.ToLower(botMessage.Text) == "had been going" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m4_task3" {
		botMessage.Text = "Модуль 4 \n Задание 3 \n Прочитайте предложения, вместо пропуска добавьте слово или фразу, которое подошло бы по значению: \n She and her family had just moved to a smaller town. There were only about 200 students at the local elementary school. Her mother hoped that Melanie ___ able to control her emotions and finally have a few friends"
	} else if strings.ToLower(botMessage.Text) == "would be" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m5_task1" {
		botMessage.Text = "Модуль 5 \n  Задание 1\n Переведите фразу так, чтобы не нарушалась лексическая сочетаемость: \n Coffee table"
	} else if strings.ToLower(botMessage.Text) == "журнальный столик" || botMessage.Text == "журнальный стол" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m5_task2" {
		botMessage.Text = "Модуль 5 \n Задание 2\n Переведите фразу так, чтобы не нарушалась лексическая сочетаемость: \n Medicinal herb"
	} else if strings.ToLower(botMessage.Text) == "лекарственное растение" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m5_task3" {
		botMessage.Text = "Модуль 5 \n Задание 3 \n Переведите фразу так, чтобы не нарушалась лексическая сочетаемость: Velvet dress"
	} else if strings.ToLower(botMessage.Text) == "бархатное платье" {
		botMessage.Text = RightAnsw
	} else {
		fmt.Println("написали не  что то не то ")
		botMessage.Text = "ой.. не могу понять твой ответ, проверь правильность написания ответа, либо обратись к списку команд :/commands или к списку заданий :/ENGtasksList"
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
