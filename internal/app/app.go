package app

import (
	"fmt"
	"log"
	"strings"

	. "goBot/config"
	. "goBot/internal/delivery/http"
	. "goBot/internal/modules"
	. "goBot/internal/service"
	. "goBot/repository/textdata"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func Run() {

	u1 := new(User)
	u2 := new(User)
	var flag = false
	var countOfUsers = false
	fmt.Println("Bot is up!")
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
			if update.Message.Text == "🙅‍♂️Прекратить общение" {
				flag = false
			}
			if flag {
				if update.Message.From.ID == u1.UserID {
					update.Message.Chat.ID = u2.UserID
					update.Message.From.FirstName = u2.UserName
					Sm(update, bot, update.Message.Text)
					continue
				} else if update.Message.From.ID == u2.UserID {
					update.Message.Chat.ID = u1.UserID
					update.Message.From.FirstName = u1.UserName
					Sm(update, bot, update.Message.Text)
					continue
				}
			}
			switch update.Message.Text {
			case "👥Поиск собеседника":
				if !countOfUsers {
					countOfUsers = true
					tx := "ожидание собеседника.."
					Sm(update, bot, tx)
					u1.UserID = update.Message.From.ID
					u1.UserName = update.Message.From.FirstName
				} else if countOfUsers {
					countOfUsers = false
					flag = true
					tx := "Собеседник найден!"
					Sm(update, bot, tx)
					u2.UserID = update.Message.From.ID
					u2.UserName = update.Message.From.FirstName
				}
			case "/open":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = NumericKeyboardOpe
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
			case "🌍В меню":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = NumericKeyboardOpe
				msg.Text = "Панель команд открыта!\n Чтобы закрыть ее напиши /close"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "📝Английский язык":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = NumericKeyboardEng
				msg.Text = "Команды английского языка!\n Чтобы вернуться в меню нажми \n🌍В меню"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			case "📝Немецкие язык":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyMarkup = NumericKeyboardGer
				msg.Text = "Команды Немецкого языка!\n Чтобы вернуться в меню нажми \n🌍В меню"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
			pseudoRand--
			test := update.Message.Text
			if strings.HasPrefix(test, "/en_") {
				tx := update.Message.Text
				Sm(update, bot, GetTask(tx))
			} else if strings.HasPrefix(test, "/ger_") {
				tx := update.Message.Text
				Sm(update, bot, GetTask(tx))
			} else if strings.HasPrefix(test, "/tr") {
				tx := "отказываться"
				Sm(update, bot, tx)
			}
			if val, ok := Answ[test]; ok {
				Sm(update, bot, val)
			}
			switch update.Message.Text {
			case "/start":
				tx := "Привет! \nTы используешь интерактивного бота для тренировки навыка переводчика!! \nНажми\n/open\nДля открытия панели кнопок\n "
				Sm(update, bot, tx)
			case "📚Информация":
				tx := `Справочная информация по модулям
			/inf_m1 - справочная информация для первого модуля
			/inf_m2 - справочная информация для второго модуля
			/inf_m3 - справочная информация для третьего модуля
			/inf_m4 - справочная информация для четвертого модуля
			/inf_m5 - справочная информация для пятого модуля
			/inf_playlists - Плейлисты с теорией и упраженниями`
				Sm(update, bot, tx)
			case "🔎Команды":
				tx := "📝Английский язык - изучение английского языка\n📝Немецкие язык - список модулей немецкого языка\n📚Информация -  список справочной информации необходимой для успешного  решения задач из  модулей \n👥Поиск собеседника - поиск собеседника для практики общения\n📈Статистика - Для просмотра своей статистики\n"
				Sm(update, bot, tx)
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
				Sm(update, bot, tx)
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
				Sm(update, bot, tx)
			case "🎧Английская речь":
				SendEnAudio(update, bot, EnAudios[pseudoRand], pseudoRand)
			}
		}
	}
}
