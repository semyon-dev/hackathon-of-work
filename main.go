package main

import (
	"fmt"
	"hackathon-work/db"
	"strings"
)

func main() {

	db.Connect()

	duties, demands := KeyWords()

	var text = "Обязанности:   Приготовление напитков в баре  Прием и подача заказов по столикам Работа с кассой\n" +
		"Требования:  Желание работать и зарабатывать. Опыт работы на подобной должности будет Вашим преимуществом.\n" +
		"А вообще всему научим и всему обучим, так что не нужно бояться приходить без опыта работы). Наличие медицинской книжки будет Вашим преимуществом.\n" +
		"Условия:  График сменный 2\\2. Будние дни с 9-00 до 21-00. Выходные дни с 9-00 до 21-00. Оформление по ТК. Оплата: 100 руб/час + % + чаевые"

	_ = text

	resumes := db.GetResumes(100, 0)
	for _, resume := range resumes {
		for _, word := range duties {
			if strings.Contains(strings.ToLower(resume.Description), word) {
				index := strings.Index(strings.ToLower(resume.Description), word)
				temp := resume.Description[index:]

			}
		}
	}
}

func KeyWords() ([]string, []string)  {

	duties := []string{"обязанност"}
	demands := []string{"требован", "приглашаем"}
	return duties, demands
}
