package main

import (
	"github.com/semyon-dev/HackatonWork/db"
	"github.com/semyon-dev/HackatonWork/model"
)

func main() {

	db.Connect()

	//duties := [...]string{"обязанност"}
	//demands := [...]string{"требован"}
	//scheduleTypes := [...]string{"смен", "2/2", "5/2", "пятидневка"}

	var text = "Обязанности:   Приготовление напитков в баре  Прием и подача заказов по столикам Работа с кассой\n" +
		"Требования:  Желание работать и зарабатывать. Опыт работы на подобной должности будет Вашим преимуществом.\n" +
		"А вообще всему научим и всему обучим, так что не нужно бояться приходить без опыта работы). Наличие медицинской книжки будет Вашим преимуществом.\n" +
		"Условия:  График сменный 2\\2. Будние дни с 9-00 до 21-00. Выходные дни с 9-00 до 21-00. Оформление по ТК. Оплата: 100 руб/час + % + чаевые"

	_ = text
}

func ParseData(raw string) model.Vacation {
	return model.Vacation{
		Id: "",
		Name:      "",
		AreaName:  "",
		City:      "",
		CompanyID: "",
		Company:   "",
		Duties:    "",
		Demands:   "",
		Terms:     model.Terms{},
	}
}
