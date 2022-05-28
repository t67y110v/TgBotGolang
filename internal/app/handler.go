package app

import (
	. "goBot/internal/delivery/http"
	//. "goBot/internal/modules"
	. "goBot/internal/service"
	. "goBot/repository/textdata"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandOpen  = "open"
	commandClose = "close"
)

func (b *Bot) handlerCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStart:
		tx := "Привет! \nTы используешь интерактивного бота для тренировки навыка переводчика!! \nНажми\n/open\nДля открытия панели кнопок\n "
		Sm(message, b.bot, tx)
	case commandOpen:
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ReplyMarkup = NumericKeyboardOpe
		msg.Text = "Панель команд открыта!\n Чтобы закрыть ее напиши /close"
		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case commandClose:
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		msg.Text = "Панель команд закрыта!\nЧтобы открыть ее напиши /open"
		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	if strings.HasPrefix(message.Command(), "en_") {
		tx := message.Text
		Sm(message, b.bot, GetTask(tx))
	} else if strings.HasPrefix(message.Command(), "ger_") {
		tx := message.Text
		Sm(message, b.bot, GetTask(tx))
	} else if strings.HasPrefix(message.Command(), "tr") {
		tx := "отказываться"
		Sm(message, b.bot, tx)
	}

}

func (b *Bot) handlerMessage(message *tgbotapi.Message, update tgbotapi.Update) {
	switch message.Text {
	case "🌍В меню":
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ReplyMarkup = NumericKeyboardOpe
		msg.Text = "Панель команд открыта!\n Чтобы закрыть ее напиши /close"
		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "📝Английский язык":
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ReplyMarkup = NumericKeyboardEng
		msg.Text = "Команды английского языка!\n Чтобы вернуться в меню нажми \n🌍В меню"
		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "📝Немецкие язык":
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		msg.ReplyMarkup = NumericKeyboardGer
		msg.Text = "Команды Немецкого языка!\n Чтобы вернуться в меню нажми \n🌍В меню"
		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	if val, ok := Answ[message.Text]; ok {
		Sm(message, b.bot, val)
	}
	switch message.Text {
	case "📚Информация":
		tx := `Справочная информация по модулям
	/inf_m1 - справочная информация для первого модуля
	/inf_m2 - справочная информация для второго модуля
	/inf_m3 - справочная информация для третьего модуля
	/inf_m4 - справочная информация для четвертого модуля
	/inf_m5 - справочная информация для пятого модуля
	/inf_playlists - Плейлисты с теорией и упраженниями`
		Sm(message, b.bot, tx)
	case "🔎Команды":
		tx := "📝Английский язык - изучение английского языка\n📝Немецкие язык - список модулей немецкого языка\n📚Информация -  список справочной информации необходимой для успешного  решения задач из  модулей \n👥Поиск собеседника - поиск собеседника для практики общения\n📈Статистика - Для просмотра своей статистики\n"
		Sm(message, b.bot, tx)
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
		Sm(message, b.bot, tx)
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
		Sm(message, b.bot, tx)
	case "🎧Английская речь":
		SendAu(update, b.bot)
	}
}
