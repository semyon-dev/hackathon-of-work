package main

import (
	"fmt"
	"hackathon-work/ML"
	"hackathon-work/config"
	"hackathon-work/db"
	"hackathon-work/mongo"
	"strconv"
)

func main() {
	config.Load()
	db.Connect()
	mongo.Connect()
	//api.RunAPI()

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
	//t()
	ML.ReadFile()
}

func t() {
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
