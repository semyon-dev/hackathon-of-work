package main

import (
	"fmt"
	"hackathon-work/db"
	"strings"
	"unicode"
)

func main() {

	db.Connect()

	//duties, _ := KeyWords()

	var text = "Обязанности:   Приготовление напитков в баре  Прием и подача заказов по столикам Работа с кассой\n" +
		"Требования:  Желание работать и зарабатывать. Опыт работы на подобной должности будет Вашим преимуществом.\n" +
		"А вообще всему научим и всему обучим, так что не нужно бояться приходить без опыта работы). Наличие медицинской книжки будет Вашим преимуществом.\n" +
		"Условия:  График сменный 2\\2. Будние дни с 9-00 до 21-00. Выходные дни с 9-00 до 21-00. Оформление по ТК. Оплата: 100 руб/час + % + чаевые"

	_ = text

	vacations := db.GetVacations(1, 11)

	for _, vacation := range vacations {
		//fmt.Println("++++++",vacation.Description, "=========")
		//for _, word := range duties {
		//	if strings.Contains(strings.ToLower(vacation.Description), word) {
		//		index := strings.Index(strings.ToLower(vacation.Description), word)
		//		fmt.Println(index)
		//		temp := []rune(vacation.Description)
		//		t := temp[index:]
		//		t = t[len([]rune(duties[0])):]
		//		t = DeleteStartSpaces(t)
		//		fmt.Println("Parsed Duties: ", string(t))
		//	}
		//}
		pieces := strings.Split(strings.ToLower(vacation.Description), "обязанности: ")
		for _, val := range pieces {
			fmt.Println("/////", val)
		}

		var duity, demand string

		if len(pieces) > 1 {
			temp := pieces[1]
			temp = DeleteStartSpaces(temp)
			if strings.Contains(temp, "требования:") {
				i := strings.Index(temp, "требования:")
				duity = temp[:i]
			}

			temps := strings.Split(strings.ToLower(temp), "требования: ")
			if len(temps) > 1 {
				t := temps[1]
				if strings.Contains(t, "условия:") {
					i := strings.Index(t, "условия:")
					demand = t[:i]
				}
				fmt.Println("Duity: ", duity, "\nDemand: ", demand)
			}

		} else {
			fmt.Println("Unsuccessful crop")
		}

	}
}

func DeleteStartSpaces(raw string) (formatted string) {
	c := 0
	for i := 0; !unicode.IsLetter(rune(raw[i])); i++ {
		c++
	}
	formatted = raw[c:]
	return formatted
}

func KeyWords() ([]string, []string) {

	duties := []string{"обязанности:", "обязанност"}
	demands := []string{"требования:", "приглашаем"}
	return duties, demands
}

func ResumeTypes() ([]string, []string, []string) {

	restaurant := []string{"официант", "ресторан", "бариста", "фаст-фуд", "повар"}
	drivers := []string{"водитель", "курьер"}
	store := []string{"склад", "кладовщик"}

	return restaurant, drivers, store
}
