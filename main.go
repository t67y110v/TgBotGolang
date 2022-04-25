package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Println("Hello world!")
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		fmt.Println("Error with Bot Api")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	var pseudoRand = len(GerAudios) + 1
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if pseudoRand == 1 || pseudoRand < 1 {
			pseudoRand = len(GerAudios) + 1
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "/open":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = numericKeyboardOpe
				msg.Text = "Панель команд открыта!\n Чтобы закрыть ее напиши /close"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "/close":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				msg.Text = "Панель команд закрыта!\nЧтобы открыть ее напиши /open"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "123":
				tx := "123"
				msg.Text = tx
			}
			pseudoRand--
			test := update.Message.Text
			if test[:4] == "/en_" {
				tx := update.Message.Text
				sm(update, bot, getTask(tx))
			} else if test[:5] == "/ger_" {
				tx := update.Message.Text
				sm(update, bot, getTask(tx))
			}
			if val, ok := answ[test]; ok {
				sm(update, bot, val)
			}
			switch update.Message.Text {
			case "/start":
				tx := "Привет! \nTы используешь интерактивного бота для тренировки навыка переводчика!! \nНажми\n/open\nДля открытия панели кнопок\n "
				sm(update, bot, tx)
			case "📚Информация":
				tx := `Справочная информация по модулям
			/inf_m1 - справочная информация для первого модуля
			/inf_m2 - справочная информация для второго модуля
			/inf_m3 - справочная информация для третьего модуля
			/inf_m4 - справочная информация для четвертого модуля
			/inf_m5 - справочная информация для пятого модуля
			/inf_playlists - Плейлисты с теорией и упраженниями`
				sm(update, bot, tx)
			case "🔎Команды":
				tx := `"/engtaskList - список модулей английского языка
				/gertasklist - список модулей немецкого языка
				/information -  список справочной информации необходимой для успешного  решения задач из  модулей 
				Для просмотра последних видео с каналов обучающих переводу :
				/yt"`
				sm(update, bot, tx)
			case "📝Немецкие задания":
				tx := `Немецкий язык
				Задания первого модуля : подобрать русский эквивалент 
				/ger_m1_task1  
				/ger_m1_task2  
				/ger_m1_task3  
				Задания второго модуля : переведите выделенное слово 
				/ger_m2_task1  
				/ger_m2_task2  
				/ger_m2_task3 
				/ger_m2_task4 
				Задания третьего модуля : перевод с использованием речевой компрессии 
				/ger_m3_task1  
				/ger_m3_task2  
				/ger_m3_task3
				Задания четвертого  модуля : заполнить пропуск подходящим словом(словами), поставив их в правильную форму  
				ger_m4_task1  
				/ger_m4_task2  
				/ger_m4_task3 
				Задания пятого модуля : переведите выделенное слово 
				/ger_m5_task1  
				/ger_m5_task2  
				/ger_m5_task3
				Задания общего модуля, для тренировки концентрации - Таблица Шульте: 
				максимально быстро сосчитать по квадратам от 1 до максимально числа 
				/m_task1
				/m_task2
				/m_task3
				/m_task4 `
				sm(update, bot, tx)
			case "📝Английские задания":
				tx := `Английский язык
				Задания первого модуля : переведите фразеологизм 
				/en_m1_task1  
				/en_m1_task2 
				/en_m1_task3  
				Задания второго модуля : перевод слова в зависимости от контекста 
				/en_m2_task1  
				/en_m2_task2  
				/en_m2_task3 
				/en_m2_task4 
				Задания третьего модуля : еревод фразы одним словом 
				/en_m3_task1  
				/en_m3_task2  
				/en_m3_task3
				Задания четвертого  модуля : вставить подходящее по смыслу слово 
				/en_m4_task1  
				/en_m4_task2  
				/en_m4_task3 
				Задания пятого модуля : перевод фразы не нарушая лексическую сочетаемость 
				/en_m5_task1  
				/en_m5_task2  
				/en_m5_task3
				Задания общего модуля, для тренировки концентрации - Таблица Шульте: 
			    максимально быстро сосчитать по квадратам от 1 до максимально числа 
				/m_task1
				/m_task2
				/m_task3
				/m_task4`
				sm(update, bot, tx)
			case "🎧Английская речь":
				SendEnAudio(update, bot, EnAudios[pseudoRand], pseudoRand)
			case "/yt":
				videoUrl, err := GetLast(ytChannels[1])
				if err != nil {
					videoUrl = "oops"
				}
				sm(update, bot, videoUrl)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			pseudoRand--
			pseudoRand--
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

		}

	}
}
