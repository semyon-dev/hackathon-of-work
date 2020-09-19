package main

import (
	"fmt"
	"hackathon-work/db"
	"math/big"
	"strings"
	"unicode"
)

func main() {

	db.Connect()

	duties, _ := KeyWords()

	var text = "Обязанности:   Приготовление напитков в баре  Прием и подача заказов по столикам Работа с кассой\n" +
		"Требования:  Желание работать и зарабатывать. Опыт работы на подобной должности будет Вашим преимуществом.\n" +
		"А вообще всему научим и всему обучим, так что не нужно бояться приходить без опыта работы). Наличие медицинской книжки будет Вашим преимуществом.\n" +
		"Условия:  График сменный 2\\2. Будние дни с 9-00 до 21-00. Выходные дни с 9-00 до 21-00. Оформление по ТК. Оплата: 100 руб/час + % + чаевые"

	_ = text

	vacations := db.GetVacations(3, 0)
	fmt.Println(vacations)
	for _, vacation := range vacations {
		for _, word := range duties {
			if strings.Contains(strings.ToLower(vacation.Description), word) {
				index := strings.Index(strings.ToLower(vacation.Description), word)
				fmt.Println(index)
				temp := []rune(vacation.Description)
				t := temp[index+12:]
				c := 0
				for i:=0; !unicode.IsLetter(t[i]); i++ {
					c++
				}
				fmt.Println("Parsed Duties: ",string(t[c:]))
			}
		}

		for _, word := range  {

		}
	}
}

func KeyWords() ([]string, []string) {

	duties := []string{"обязанности:", "обязанност"}
	demands := []string{"требования:", "приглашаем"}
	return duties, demands
}

func ResumeTypes() ([]string, []string, []string) {

	restaurant := []string{"официант"}

}
