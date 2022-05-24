package service

import (
	"fmt"
	"strconv"
	"strings"
)

var task_description = map[string]string{
	"en$1":  "⚫Переведите фразеолагизм\n",
	"en$2":  "⚫Переведите выделенное слово в зависимости от контекста\n",
	"en$3":  "⚫Переведите фразу на русский одним словом\n",
	"en$4":  "⚫Вставьте подходящее по смыслу слово\n",
	"en$5":  "⚫Переведите фразу так, чтобы не нарушалась лексическая сочетаемость\n",
	"ger$1": "⚪Переведите фразеолагизм\n",
	"ger$2": "⚪Переведите фразеолагизм\n",
	"ger$3": "⚪Переведите фразу на русский одним словом\n",
	"ger$4": "⚪Вставьте подходящее по смыслу слово\n",
	"ger$5": "⚪Переведите фразу так, чтобы не нарушалась лексическая сочетаемость\n"}

var tasks = map[string]string{
	"en$1$1":  "Easier said then done",
	"en$1$2":  "A heart-to-heart talk",
	"en$1$3":  "A hard nut to crack",
	"en$2$1":  "Her room is $minute$",
	"en$2$2":  "I $refuse$ to take out refuse",
	"en$2$3":  "I refuse to take out $refuse$",
	"en$3$1":  "Fried eggs sunny side up",
	"en$3$2":  "Border patrol agents",
	"en$3$3":  "White blood cells",
	"en$4$1":  "Melanie River had just turned ten years old, but she had never been to school. Most children in the third grade ___ to school for two years. But Melanie was not most children",
	"en$4$2":  "She and her family had just moved to a smaller town. There were only about 200 students at the local elementary school. Her mother hoped that Melanie ___ able to control her emotions and finally have a few friends",
	"en$5$1":  "Coffee table",
	"en$5$2":  "Medicinal herb",
	"en$5$3":  "Velvet dress",
	"ger$1$1": "Übung macht den Meister",
	"ger$1$2": "Lügen haben kurze Beine",
	"ger$1$3": "Andere Länder, andere Sitten",
	"ger$2$1": "Die Schraube und die $Mutter$ passen nicht zusammen",
	"ger$2$2": "Die $Mutter$ meiner Mutter war Schottin",
	"ger$2$3": "Beim Rausgehen brach ihr der $Absatz$ ab",
	"ger$2$4": "der $Absatz$ steht auf der ersten Seite des Dokuments",
	"ger$3$1": "Etwas in Angriff nehmen",
	"ger$3$2": "Auf jdn/etw Einfluss nehmen",
	"ger$3$3": "Etwas in Aussicht stellen",
	"ger$4$1": "In einer der ältesten und bedeutendsten Totenstädte Ägyptens fanden die Forscher aus Tübingen vergoldete Maske. Sie lag in einem beschädigten Holzsarg in Sakkara und ist mehr als 2.500 ______. Sie soll einem altägyptischen Priester aus der 26. Dynastie gehört haben. Auffallend sind die großen Augen.",
	"ger$4$2": "Wasserstoff ist das am häufigsten vorkommende chemische _____ im Universum, weshalb es wahrscheinlich auch den ersten Platz im Periodensystem verdient. Es besteht lediglich aus einem Proton und einem Elektron.",
	"ger$4$3": "Im Ozean wird Plastik in Mikro- und Nanoplastikpartikel (MNP) zersetzt, also in winzige Kleinstteile. In den nächsten 30 Jahren wird sich der Gehalt dieser Partikel in den Weltmeeren ___ der WWF-Studie verdoppeln.",
	"ger$5$1": "Die Dose Mais",
	"ger$5$2": "Ich habe einen Termin",
	"ger$5$3": "Dieser Dom wurde noch in der Mitte des 16. Jahrhunderts gebaut"}

func GetTask(mes string) string {

	module, task := Convert(mes)
	modul, tas := strconv.Itoa(module), strconv.Itoa(task)
	language := Getlanguage(mes)
	fmt.Println(modul, tas)
	readyMes := "Тема " + modul + "\nЗадание " + tas + "\n" + task_description[language+"$"+modul] + tasks[language+"$"+modul+"$"+tas]

	return readyMes
}
func Getlanguage(s string) string {
	if strings.HasPrefix(s, "/en_m") {
		return "en"
	} else if strings.HasPrefix(s, "/ger_m") {
		return "ger"
	} else {
		return "en"
	}
}
