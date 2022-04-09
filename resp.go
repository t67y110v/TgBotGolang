package main

import (
	//"bolt"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	if botMessage.Text == "/start" {
		fmt.Println("зашел  в старт")
		botMessage.Text = "Привет! \nTы используешь интерактивного бота для тренировки навыка переводчика!! \nСписок доступных команд\n/commands\n\n "
	} else if strings.ToLower(botMessage.Text) == "/commands" {
		fmt.Println("зашел  в список команд  ")
		botMessage.Text = "/ENGtaskList - список модулей английского языка\n/GERtaskList - список модулей немецкого языка\n/information -  список справочной информации необходимой для успешного  решения задач из  модулей \nДля просмотра последних видео с каналов обучающих переводу :\n/yt_1\n/yt_2\n/yt_3\n"
	} else if strings.ToLower(botMessage.Text) == "/information" {
		fmt.Println("зашел  в информацию  ")
		botMessage.Text = "Справочная информация по модулям\n/inf_m1 - справочная информация для первого модуля\n/inf_m2 - справочная информация для второго модуля\n/inf_m3 - справочная информация для третьего модуля\n/inf_m4 - справочная информация для четвертого модуля\n/inf_m5 - справочная информация для пятого модуля\n/inf_playlists - Плейлисты с теорией и упраженниями"
	} else if strings.ToLower(botMessage.Text) == "/gertasklist" {
		fmt.Println("зашел  в информацию  ")
		botMessage.Text = "Немецкий язык\nЗадания первого модуля : подобрать русский эквивалент \n/ger_m1_task1  \n/ger_m1_task2  \n/ger_m1_task3  \nЗадания второго модуля : переведите выделенное слово \n/ger_m2_task1  \n/ger_m2_task2  \n/ger_m2_task3 \n/ger_m2_task4 \nЗадания третьего модуля : перевод с использованием речевой компрессии \n/ger_m3_task1  \n/ger_m3_task2  \n/ger_m3_task3\n Задания четвертого  модуля : заполнить пропуск подходящим словом(словами), поставив их в правильную форму  \n/ger_m4_task1  \n/ger_m4_task2  \n/ger_m4_task3 \nЗадания пятого модуля : переведите выделенное слово \n/ger_m5_task1  \n/ger_m5_task2  \n/ger_m5_task3\nЗадания общего модуля, для тренировки концентрации - Таблица Шульте: \nмаксимально быстро сосчитать по квадратам от 1 до максимально числа \n/m_task1\n/m_task2\n/m_task3\n/m_task4 "
	} else if strings.ToLower(botMessage.Text) == "/engtasklist" || strings.ToLower(botMessage.Text) == "/engtaklist" || strings.ToLower(botMessage.Text) == "/engtaskslist" {
		fmt.Println("зашел  в таксклист")
		botMessage.Text = "Английский язык\nЗадания первого модуля : переведите фразеологизм \n/en_m1_task1  \n/en_m1_task2  \n/en_m1_task3  \nЗадания второго модуля : перевод слова в зависимости от контекста \n/en_m2_task1  \n/en_m2_task2  \n/en_m2_task3 \n/en_m2_task4 \nЗадания третьего модуля : еревод фразы одним словом \n/en_m3_task1  \n/en_m3_task2  \n/en_m3_task3\n Задания четвертого  модуля : вставить подходящее по смыслу слово \n/en_m4_task1  \n/en_m4_task2  \n/en_m4_task3 \nЗадания пятого модуля : перевод фразы не нарушая лексическую сочетаемость \n/en_m5_task1  \n/en_m5_task2  \n/en_m5_task3\nЗадания общего модуля, для тренировки концентрации - Таблица Шульте: \nмаксимально быстро сосчитать по квадратам от 1 до максимально числа \n/m_task1\n/m_task2\n/m_task3\n/m_task4"
	} else if strings.ToLower(botMessage.Text) == "/en_m2_task1" {
		botMessage.Text = "Модуль 2\n Задание 1 \nНапишите перевод  выделенного /***/ слова в заданном контексте \nHer room is /minute/."
	} else if strings.ToLower(botMessage.Text) == "маленькая" || strings.ToLower(botMessage.Text) == "крошечная" || strings.ToLower(botMessage.Text) == "маленький" || strings.ToLower(botMessage.Text) == "крошечный" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m2_task2" {
		botMessage.Text = "Модуль 2 \nЗадание 2 \nНапишите перевод  выделенного /***/ слова в заданном контексте \nI /refuse/ to take out refuse."
	} else if strings.ToLower(botMessage.Text) == "отказываюсь" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m2_task3" {
		botMessage.Text = "Модуль 2 \nЗадание 3 \nНапишите перевод  выделенного /***/ слова в заданном контексте \nI refuse to take out /refuse/."
	} else if strings.ToLower(botMessage.Text) == "мусор" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m3_task1" {
		botMessage.Text = "Модуль 3 \nЗадание 1\nПопробуйте перевести данные слова одним словом\nFried eggs sunny side up"
	} else if strings.ToLower(botMessage.Text) == "глазунья" || strings.ToLower(botMessage.Text) == "яичница" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m3_task2" {
		botMessage.Text = "Модуль 3 \nЗадание 2 \n Попробуйте перевести данные слова одним словом\nBorder patrol agents"
	} else if strings.ToLower(botMessage.Text) == "пограничники" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m3_task3" { //   botMessage.Text[:3]==/en int(botMessage.Text[6])  int(botMessage.Text[12])
		//t:="Тема " + botMessage.Text[6] + "Задане " + botMessage.Text[12] + listss[int(botMessage.Text[6])]
		botMessage.Text = "Модуль 3 \nЗадание 3\nПопробуйте перевести данные слова одним словом\nWhite blood cells"
	} else if strings.ToLower(botMessage.Text) == "лейкоциты" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task1" {
		botMessage.Text = "Модуль 1 \nЗадание 1 \nНапишите перевод данного фразеологизма:\nEas­i­er said then done"
	} else if strings.ToLower(botMessage.Text) == "проще сказать, чем сделать" || strings.ToLower(botMessage.Text) == "легче сказать, чем сделать" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task2" {
		botMessage.Text = "Модуль 1 \nЗадание 2 \nНапишите перевод данного фразеологизма:\nA heart-to-heart talk"
	} else if strings.ToLower(botMessage.Text) == "разговор по душам" || strings.ToLower(botMessage.Text) == "беседа по душам" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m1_task3" {
		botMessage.Text = "Модуль 1 \nЗадание 3\nНапишите перевод данного фразеологизма:\n A hard nut to crack"
	} else if strings.ToLower(botMessage.Text) == "крепкий орешек" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m4_task1" {
		botMessage.Text = "Модуль 4 \nЗадание 1 \nПрочитайте предложения, вместо пропуска добавьте слово или фразу, которое подошло бы по значению:\nMelanie River had just turned ten years old, but she had never been to school. Most children in the third grade ___ to school for two years. But Melanie was not most children"
	} else if strings.ToLower(botMessage.Text) == "had been going" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m4_task2" {
		botMessage.Text = "Модуль 4 \nЗадание 2 \nПрочитайте предложения, вместо пропуска добавьте слово или фразу, которое подошло бы по значению: \nShe and her family had just moved to a smaller town. There were only about 200 students at the local elementary school. Her mother hoped that Melanie ___ able to control her emotions and finally have a few friends"
	} else if strings.ToLower(botMessage.Text) == "would be" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m5_task1" {
		botMessage.Text = "Модуль 5\nЗадание 1\nПереведите фразу так, чтобы не нарушалась лексическая сочетаемость: \nCoffee table"
	} else if strings.ToLower(botMessage.Text) == "журнальный столик" || botMessage.Text == "журнальный стол" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/en_m5_task2" {
		botMessage.Text = "Модуль 5 \nЗадание 2\nПереведите фразу так, чтобы не нарушалась лексическая сочетаемость: \nMedicinal herb"
	} else if strings.ToLower(botMessage.Text) == "лекарственное растение" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/en_m5_task3" {
		botMessage.Text = "Модуль 5 \nЗадание 3 \nПереведите фразу так, чтобы не нарушалась лексическая сочетаемость: Velvet dress"
	} else if strings.ToLower(botMessage.Text) == "бархатное платье" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m1_task1" {
		botMessage.Text = "Тема 1\nЗадание 1\nПереведите фразеолагизм\nÜbung macht den Meister"
	} else if strings.ToLower(botMessage.Text) == "повторение - мать учения" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m1_task2" {
		botMessage.Text = "Тема 1\nЗадание 2\nПереведите фразеолагизм\nLügen haben kurze Beine"
	} else if strings.ToLower(botMessage.Text) == "Всё тайное становится явным" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m1_task3" {
		botMessage.Text = "Тема 1\nЗадание 1\nПереведите фразеолагизм\nAndere Länder, andere Sitten"
	} else if strings.ToLower(botMessage.Text) == "В чужой монастырь со своим уставом не ходят" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m2_task1" {
		botMessage.Text = "Тема 2\nЗадание 1\nПереведите выделенное слово в зависимости от контекста\nDie Schraube und die $$$Mutter$$$ passen nicht zusammen"
	} else if strings.ToLower(botMessage.Text) == "гайка" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m2_task2" {
		botMessage.Text = "Тема 2\nЗадание 2\nПереведите выделенное слово в зависимости от контекста\nDie $$$Mutter$$$ meiner Mutter war Schottin"
	} else if strings.ToLower(botMessage.Text) == "мама" || strings.ToLower(botMessage.Text) == "мать" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m2_task3" {
		botMessage.Text = "Тема 2\n Задание \nПереведите выделенное слово в зависимости от контекста\nBeim Rausgehen brach ihr der $$$Absatz$$$ ab"
	} else if strings.ToLower(botMessage.Text) == "каблук" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m2_task4" {
		botMessage.Text = "Тема21\n Задание 4\nПереведите выделенное слово в зависимости от контекста\nder $$$Absatz$$$ steht auf der ersten Seite des Dokuments"
	} else if strings.ToLower(botMessage.Text) == " абзац" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m3_task1" {
		botMessage.Text = "Тема 3\nЗадание 1\nПереведите данное выражение, применив способ речевой компрессии\nEtwas in Angriff nehmen"
	} else if strings.ToLower(botMessage.Text) == "приступить" || strings.ToLower(botMessage.Text) == "приступать" || strings.ToLower(botMessage.Text) == "начать" || strings.ToLower(botMessage.Text) == "начинать" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m3_task2" {
		botMessage.Text = "Тема 3\nЗадание 2\nПереведите данное выражение, применив способ речевой компрессии\nAuf jdn/etw Einfluss nehmen"
	} else if strings.ToLower(botMessage.Text) == "влиять" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m3_task3" {
		botMessage.Text = "Тема 3\n Задание 3\nПереведите данное выражение, применив способ речевой компрессии\nEtwas in Aussicht stellen"
	} else if strings.ToLower(botMessage.Text) == "обещать" || strings.ToLower(botMessage.Text) == "пообещать" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m4_task1" {
		botMessage.Text = "Тема 4\nЗадание 1\nЗаполните пропуск подходящим по смыслу словом (словами), поставив его  в правильную форму\nIn einer der ältesten und bedeutendsten Totenstädte Ägyptens fanden die Forscher aus Tübingen vergoldete Maske. Sie lag in einem beschädigten Holzsarg in Sakkara und ist mehr als 2.500 __. Sie soll einem altägyptischen Priester aus der 26. Dynastie gehört haben. Auffallend sind die großen Augen."
	} else if strings.ToLower(botMessage.Text) == "jahre alt" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m4_task2" {
		botMessage.Text = "Тема 4\nЗадание 2\nЗаполните пропуск подходящим по смыслу словом, поставив его  в правильную форму\nWasserstoff ist das am häufigsten vorkommende chemische _ im Universum, weshalb es wahrscheinlich auch den ersten Platz im Periodensystem verdient. Es besteht lediglich aus einem Proton und einem Elektron."
	} else if strings.ToLower(botMessage.Text) == "element" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m4_task3" {
		botMessage.Text = "Тема 4\n Задание 3\nЗаполните пропуск подходящим по смыслу словом, поставив его  в правильную форму\nIm Ozean wird Plastik in Mikro- und Nanoplastikpartikel (MNP) zersetzt, also in winzige Kleinstteile. In den nächsten 30 Jahren wird sich der Gehalt dieser Partikel in den Weltmeeren ___ der WWF-Studie verdoppeln."
	} else if strings.ToLower(botMessage.Text) == "laut" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m5_task1" {
		botMessage.Text = "Тема 5\nЗадание 1\nПереведите выделенное слово\nDie /Dose/ Mais"
	} else if strings.ToLower(botMessage.Text) == "банка" || strings.ToLower(botMessage.Text) == "консервная банка" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m5_task2" {
		botMessage.Text = "Тема 5\nЗадание 2\nПереведите выделенное слово\nIch habe einen /Termin/"
	} else if strings.ToLower(botMessage.Text) == "встреча" {
		botMessage.Text = RightAnsw
	} else if botMessage.Text == "/ger_m5_task3" {
		botMessage.Text = "Тема 5\n Задание 3\nПереведите выделенное слово\nDieser /Dom/ wurde noch in der Mitte des 16. Jahrhunderts gebaut"
	} else if strings.ToLower(botMessage.Text) == "собор" {
		botMessage.Text = RightAnsw
	} else if strings.ToLower(botMessage.Text) == "/m_task1" {
		botMessage.Text = "https://4brain.ru/skorochtenie/images/tablicy-shulte/hilevel/7.jpg"
	} else if strings.ToLower(botMessage.Text) == "/m_task2" {
		botMessage.Text = "https://i.pinimg.com/564x/23/9e/1b/239e1bb358ee7f6e7787b53761bb9d9d.jpg"
	} else if strings.ToLower(botMessage.Text) == "/m_task3" {
		botMessage.Text = "https://pandia.ru/text/80/438/images/img4_59.jpg"
	} else if strings.ToLower(botMessage.Text) == "/m_task4" {
		botMessage.Text = "https://pandia.ru/text/80/438/images/img3_72.jpg"
	} else if strings.ToLower(botMessage.Text) == "/inf_m1" {
		botMessage.Text = "При переводе фразеологической единицы задача переводчика заключается не только в том, чтобы верно передать её смысл, но и отразить эмоционально-экспрессивные характеристики, оценочную коннотацию, функционально-стилистические особенности. Также причиной возникновения трудностей при переводе фразеологизма может стать высокая степень его национальной специфичности. В таких случаях задачей переводчика будет его адаптация к культуре и языку целевой аудитории.Рассмотрим ряд приёмов перевода фразеологических единиц и проанализируем их применение на практике:"
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "1.Подбор эквивалента\nPoor lamb, he must be as poor as a church mouse.\nАх ты, ягненочек! Видно, беден, как церковная мышь\nДля перевода данного фразеологизма автор использовал его полный эквивалент в русском языке. Эти два фразеологизма имеют одинаковую семантику, компонентный состав и функционально-стилистические особенности. На наш взгляд, перевод выполнен удачно. Также мы можем предложить ещё один вариант перевода данного фразеологизма: «без гроша за душой»."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "2.Подбор аналога\nWhat he said had a hateful truth in it, and another defect of my character is that I enjoy the company of those, however depraved, who can give me a Roland for my Oliver.\nОн высказал роковую истину. Мне нравятся люди пусть дурные, но которые за словом в карман не лезут\n\nВвиду отсутствия эквивалента данного английского фразеологизма в русском языке, переводчик подобрал ему аналог, построенный на другом образе. Она была вынуждена отказаться от передачи образа фразеологизма оригинала. Его сохранение и буквальный перевод был бы непонятен русскому читателю, поскольку здесь прослеживается национальная специфика: Роланд и Оливер – персонажи французской героической поэмы «Песнь о Роланде», которые сражались друг с другом, но, поскольку их силы были равны, ни один из них так и не победил. Однако семантика русского фразеологизма в тексте перевода отличается от семантики фразеологизма в тексте оригинала. Выражение «за словом в карман не лезет» обычно характеризует человека, который может непринуждённо и легко вести беседу, быть остроумным и быстро находить ответы. В английском же фразеологизме подчёркивается то, что человек может парировать, дать отпор. Предложенный переводчиком вариант несколько неточен, однако может считаться удачным, поскольку он передаёт большую часть информации, заложенной в оригинале. Также этот фразеологизм можно перевести как «дать достойный ответ», «удачно парировать»."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "3. Описательный перевод\nIt had been done when he took silk and it represented him in a wig and gown. Eventheycouldnotmakehimimposing…\nОн тогда только что стал королевским адвокатом и по этому случаю был снят в парике и в мантии, но даже это не придало ему внушительности.\n \nВ оригинале мы видим метафоричность, однако переводчик был вынужден отказаться от неё при переводе. Её сохранение и буквальный перевод фразеологизма как «одеваться в шёлк» были бы непонятны русскому читателю, поскольку здесь явно прослеживается национальная специфика. В Великобритании обычные адвокаты, выступающие в суде, носят суконную мантию, в то время как королевские облачаются в шёлковую, по этой причине их вступление в должность называется «to take silk». На наш взгляд, перевод выполнен удачно."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "4. Лексический перевод\nHe'd be a bit surly sometimes, but when we hadn't had a bite since morning, and we hadn't even got the price of a lie down at the Chink's, he'd be as lively as a cricket.\nИногда он, конечно, хмурился, но, если у нас с утра до вечера маковой росинки во рту не бывало и нечем было заплатить китаезе за ночлег, он только посмеивался. \nВ данном случае мы имеем дело с лексическим переводом, поскольку рассматриваемый нами фразеологизм переведён на русский язык одной лексической единицей. Кроме того, здесь наблюдается его переосмысление автором. Это привело к смысловой неточности: в оригинале, по задумке автора, данный фразеологизм подчёркивает, что, несмотря на непростую ситуацию, герой не унывает и остаётся жизнерадостным, что никак не отражено в переводе. Можно отредактировать перевод следующим образом: «Иногда он, конечно, хмурился, но, даже когда у нас с утра до вечера маковой росинки во рту не бывало и нечем было заплатить китаезе за ночлег, он не унывал/оставался жизнерадостным»."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "5. Контекстуальный перевод\nIt's a smack in the eye of course, it's no good denying that, but the only thing is to grin and bear it.\nКонечно, спорить не приходится, это удар по самолюбию, но что мне остается? Улыбнуться, и все. Как-нибудь переживем.\nДанный фразеологизм метафоричен, относится к разговорной лексике и часто переводится на русский язык как «жестокое разочарование», «удар», «досада», «неприятность». В данном же случае переводчик прибегнул контекстуальному переводу и полной замене образа."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "6. Антонимический перевод\n… and with the possibility that Michael might be killed at any moment - it was true he said he was as safe as a house, he only said that to reassure her, and even generals were killed sometimes - if she was to go on living she must have a child by him.\n… и притом, что Майкла могли в любой момент убить, - конечно, он говорил, что ему абсолютно ничего не грозит, но он просто успокаивал ее, даже генералов и тех убивали, - удержать ее в жизни мог только его ребенок.\nВ данном случае мы видим, что переводчик решил пожертвовать образом, поскольку он остался бы непонятным русскому читателю, и прибегнуть антонимическому переводу. Если не использовать антонимический перевод, данный фразеологизм можно перевести следующим образом: «вне опасности», «в полной безопасности». Однако, принимая во внимание контекст, который здесь немаловажен, вариант, предложенный переводчиком, звучит более уместно. В данном отрывке речь идёт о муже главной героини, который ушёл добровольцем на фронт, а принимая участие в военных действиях, вряд ли возможно оставаться «вне опасности» или «в полной безопасности». Однако в другом контексте данные варианты имеют место быть"
	} else if strings.ToLower(botMessage.Text) == "/inf_m2" {
		botMessage.Text = "мы надеялись сюда никто не тыкнет, пожалуйста, сделайте вид что так и должно быть и посмотрите информацию по другим модулям"
	} else if strings.ToLower(botMessage.Text) == "/inf_m3" {
		botMessage.Text = "Речевая компрессия (3 модуль)\nВыделяют следующие случаи, когда у переводчика появляются основания применить тот или иной способ речевой компрессии:\n1. На языке оригинала есть повторы.\n2. На языке перевода есть ничего не значащие слова.\n3. Оратор говорит слишком быстро.\n4. Предметная ситуация позволяет выразить ту же мысль меньшим количеством слов, что приводит к экономии времени и соответственно возможности более полно концентрировать внимание на поступающих смысловых группах"
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "Разновидности речевых компрессий:\n1. Устранение компонентов текста, содержащих информацию, которая восполняется экстралингвистической ситуацией общения.\n2. Устранение компонентов, содержащих в себе инфомедию, которая не содержится в экстралингвистической ситуации общения.\nВторая группа способов речевой компрессии, основанной на синонимической замене, включает в себя следующие способы:\na.     Сокращение названия\nb.     Преобразование глагольного словосочетания\nc.     Преобразование придаточного предложения\nd.     Замена повторов и параллелизмов\ne.     Замена экстралингвистической информации\nf.      Обобщение целого предложения"

	} else if botMessage.Text == "/inf_m4" {
		botMessage.Text = "Вероятностное прогнозирование (4 модуль)\nС новичками в синхронном переводе часто случается следующая ситуация: «We live in a new technology/technological…», далее молчание. Причин может быть несколько: забыл перевод, потерял нить перевода, не расслышал оригинал.\nКогда в такой ситуации вы «бросаете» предложение, это явно будет заметно для слушателя. Так сложится впечатление, что переводчик не договаривает, и доверие к нему упадет."
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "Выход один – закончить предложение. На помощь приходит прием «вероятностное прогнозирование»: можно закончить ранее озвученной информацией или предположить, что могло бы тут быть по смыслу выступления. Главное здесь – не противоречить всему, что было сказано ранее, не вводить кардинально новую информацию.\nВозьмем пример выше: «We live in a new technology/technological ERA» – era в данном случае введенный элемент на основе вероятностного прогнозирования.\nДля того, чтобы развить в себе способность быстро реагировать в таких ситуациях, можно тренироваться самостоятельно.\nОткрываете ленту новостей, берете любую новость, «выбрасываете» из нее слова или выражения, откладываете до завтра, чтобы забыть, о чем речь и читаете этот текст, вставляя пропущенные слова и выражения, не задерживаясь на этих моментах.\nБудет еще лучше, если ваши родные смогут подобрать вам текст, выпустить слова и дать вам для тренировки, тогда не будет вероятности, что вы запомните что-то из текста и будете вставлять слова по памяти."

	} else if strings.ToLower(botMessage.Text) == "/inf_m5" {
		botMessage.Text = "Лексическая сочетаемость (5 модуль англ.яз.)\nКогда мы говорим на человеческом языке, мы складываем слова во фразы, пользуясь какими-то правилами грамматики, но на самом деле есть очень много тонкостей, которые мы при этом незаметно для себя учитываем. В частности, далеко не любые слова можно соединить в рамках словосочетаний и предложений. "
		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "Ложные друзья переводчика (5 модуль нем.яз.)\nЛожные друзья переводчика — это слова, которые в двух или нескольких языках звучат или даже пишутся одинаково или очень похоже, что может привести к путанице и непониманию. Рассмотрим самые популярные немецкие слова, с которыми у вас могут возникнуть трудности при переводе."
		buf, err = json.Marshal(botMessage)
		if err != nil {
			return err
		}
		_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
		botMessage.Text = "der Zentner — в немецком и русском языках отличается мера веса (центнер). В немецком это масса, равная 50 кг, в русском — 100 кг.\ndas Glas — стакан, рюмка или просто стекло, а не “глаз” — das Auge.\nder Dom – означает в немецком языке собор, а не “дом” - das Haus.\ndie Krawatte — галстук, а не “кровать” — das Bett\n	reklamieren — предъявлять претензии, жалобы, а не “рекламировать” — werben.\ndas Wetter — погода, а не “ветер” — der Wind.\ndie Angel — удочка, a не (ангел) — der Engel.\ndie Dose — банка, штепсельная розетка, а не (доза) — die Dosis.\nder Führer — не только (должность) Гитлера, но также вождь, командир, машинист, вагоновожатый, капитан спортивной команды.\ndas Glück — счастье, благополучие, удача, успех, а не (глюк) в значении (галлюцинация) — die Halluzination/die Wahnvorstellung.\nder Tank — это вовсе и не (танк), а всего лишь (бак) или (цистерна), “танк” по-немецки - der Panzer.\nder Panzer - не только “панцирь”, а также танк\ndas Magazin — (журнал), а не магазин - das Geschäft."

	} else if strings.ToLower(botMessage.Text[:4]) == "/tra" {
		botMessage.Text = translateFunction(botMessage.Text[:4])
	} else if strings.ToLower(botMessage.Text) == "/inf_playlists" {
		botMessage.Text = "Плейлист “Инструменты переводчика”\nhttps://youtube.com/playlist?list=PL7F-BAOq1U91t4u3Uc4b_UZbpJ3Lg5Ypq\nПлейлист “Упражнения для перевода”\nhttps://youtube.com/playlist?list=PL7F-BAOq1U90JZDKkz6MXEZSGzMY1jwaL\nПлейлист “Теория перевода”\nhttps://youtube.com/playlist?list=PL7F-BAOq1U91VGBq318fwrzfTYJsGkGun\nПлейлист “Устный перевод”\nhttps://youtube.com/playlist?list=PL7F-BAOq1U91A4buZYukp4kKOjwf4MLFr\nПлейлист “Письменный перевод”\nhttps://youtube.com/playlist?list=PL7F-BAOq1U910Ozt6zHZuSWoNBwLD5tWZ\n"
	} else if strings.ToLower(botMessage.Text[:4]) == "http" {
		videoUrl, err := GetLast(update.Message.Text)
		if err != nil {
			return err
		}
		botMessage.Text = videoUrl
	} else if strings.ToLower(botMessage.Text) == "/yt_1" {
		videoUrl, err := GetLast(YT_CH_URL1)
		if err != nil {
			return err
		}
		botMessage.Text = videoUrl
	} else if strings.ToLower(botMessage.Text) == "/yt_2" {
		videoUrl, err := GetLast(YT_CH_URL2)
		if err != nil {
			return err
		}
		botMessage.Text = videoUrl
	} else if strings.ToLower(botMessage.Text) == "/yt_3" {
		videoUrl, err := GetLast(YT_CH_URL3)
		if err != nil {
			return err
		}
		botMessage.Text = videoUrl
	} else {
		fmt.Println("написали не  что то не то ")
		botMessage.Text = "ой.. не могу понять твой ответ, проверь правильность написания ответа, либо обратись к списку команд :/commands или к списку заданий :/ENGtaskList"
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
