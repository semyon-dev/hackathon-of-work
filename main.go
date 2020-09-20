package main

import (
	"fmt"
	"hackathon-work/ML"
	"hackathon-work/api"
	"hackathon-work/config"
	"hackathon-work/db"
	"hackathon-work/mongo"
	"strconv"
)

func main() {
	config.Load()
	db.Connect()
	mongo.Connect()
	api.RunAPI()

	//done0 := make(chan struct{})
	//done1 := make(chan struct{})
	//done2 := make(chan struct{})
	//
	//go db.UpdateRestaurantsAnswers(done0)
	//go db.UpdateDriversAnswers(done1)
	//go db.UpdateStoreAnswers(done2)
	//
	//<-done0
	//<-done1
	//<-done2
	//Restaruns()
	//Store()
	//ML.ReadFile()
}

func Restaruns() {
	restaurants := mongo.GetAllRestaurants()
	var ch int
	for _, v := range restaurants {
		if v.Description == "" {
			continue
		}
		fmt.Println(ch)
		ch++
		if ch == 5000 {
			ch = 0
			ML.FINISH()
		}
		if v.Q3 == "" {
			// профессиональное образовательное учреждение, где вы учились, и специальность
			ML.CreateAnswer(v.Description, "профессиональное образовательное учреждение, где вы учились и специальность?", strconv.Itoa(int(v.Id))+"_3")
		}
		if len(v.Q6) == 0 {
			ML.CreateAnswer(v.Description, "Рестораны где работал?", strconv.Itoa(int(v.Id))+"_6")
		}
		if len(v.Q7) == 0 {
			ML.CreateAnswer(v.Description, "дополнительные профессиональные знания и навыки?", strconv.Itoa(int(v.Id))+"_7")
		}
	}
}

func Store() {
	restaurants := mongo.GetAllStores()
	var ch int
	for _, v := range restaurants {
		if v.Description == "" {
			continue
		}
		fmt.Println(ch)
		ch++
		if ch == 5000 {
			ch = 0
			ML.FINISH()
		}
		// Отрасль
		if len(v.Q4) == 0 {
			// Отрасль
			ML.CreateAnswer(v.Description, "Отрасль?", strconv.Itoa(int(v.Id))+"_4")
		}
		if len(v.Q6) == 0 {
			ML.CreateAnswer(v.Description, "С какой системой хранения Вы имели опыт работы?", strconv.Itoa(int(v.Id))+"_6")
		}
		if len(v.Q7) == 0 {
			ML.CreateAnswer(v.Description, "Опыт работы с программами?", strconv.Itoa(int(v.Id))+"_7")
		}
		if len(v.Q8) == 0 {
			ML.CreateAnswer(v.Description, "Опыт с инструментарием?", strconv.Itoa(int(v.Id))+"_8")
		}
		if len(v.Q9) == 0 {
			ML.CreateAnswer(v.Description, "Типы работ?", strconv.Itoa(int(v.Id))+"_9")
		}
		if v.Q11 == 0 {
			ML.CreateAnswer(v.Description, "Уровень ЗП (желаемый)?", strconv.Itoa(int(v.Id))+"_11")
		}
	}
}
