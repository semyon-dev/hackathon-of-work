package main

import (
	"fmt"
	"hackathon-work/api"
	"hackathon-work/db"
	"hackathon-work/model"
	"strings"
	"unicode"
)

func main() {

	db.Connect()

	restaurant, drivers, store := CreateTypes()

	go api.RunAPI()

	//duties, _ := KeyWords()

	var text = "Обязанности:   Приготовление напитков в баре  Прием и подача заказов по столикам Работа с кассой\n" +
		"Требования:  Желание работать и зарабатывать. Опыт работы на подобной должности будет Вашим преимуществом.\n" +
		"А вообще всему научим и всему обучим, так что не нужно бояться приходить без опыта работы). Наличие медицинской книжки будет Вашим преимуществом.\n" +
		"Условия:  График сменный 2\\2. Будние дни с 9-00 до 21-00. Выходные дни с 9-00 до 21-00. Оформление по ТК. Оплата: 100 руб/час + % + чаевые"

	_ = text

	vacations := db.GetVacations(1, 11)

	for _, vacation := range vacations {
		fmt.Println(vacation.Id)
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

				vacation.Duties = duity
				vacation.Demands = demand

				vacation = DetectType(vacation, restaurant, "restaurant")
				vacation = DetectType(vacation, drivers, "drivers")
				vacation = DetectType(vacation, store, "store")

				fmt.Printf("Type: %s,\nDuity: %s,\nDemand: %s.", vacation.Type, vacation.Duties, vacation.Demands)

			} else {
				fmt.Println("Unsuccessful Demand crop")
			}

		} else {
			fmt.Println("Unsuccessful Duty crop")
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

func CreateTypes() ([]string, []string, []string) {

	restaurant := []string{"официант", "ресторан", "бариста", "фаст-фуд", "повар"}
	drivers := []string{"водитель"}
	store := []string{"склад", "кладовщик", "комплектовщик"}

	return restaurant, drivers, store
}

func DetectType(vacation model.NewVacation, words []string, wantedType string) model.NewVacation {
	for _, word := range words {
		if strings.Contains(strings.ToLower(vacation.Name), word) {

			vacation.Type = wantedType
			db.UpdateNewVacation(vacation)
			break
		}
	}
	return vacation
}
