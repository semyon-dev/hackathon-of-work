package ML

import (
	"encoding/json"
	"fmt"
	"hackathon-work/model"
	"log"
)

func CreateAnswer() {
	var res model.Res
	var data model.Data
	data.Title = "title"
	var a []model.Answer

	Answer := model.Answer{
		AnswerStart: 0,
		Text:        "text",
	}
	a = append(a, Answer)
	qas := model.Qas{
		Answers:  a,
		Question: "Сколько вы уже работаете?",
		Id:       "2",
	}
	var Qas []model.Qas
	var Paragraphs []model.Paragraph
	Qas = append(Qas, qas)
	Paragraph := model.Paragraph{
		Context: "я работаю уже 15 лет, меня зовут Дима, hello",
		Qas:     Qas,
	}
	Paragraphs = append(Paragraphs, Paragraph)
	data.Paragraphs = Paragraphs
	res = append(res, data)
	resM, err := json.Marshal(res)
	log.Println(err)
	fmt.Println(string(resM))
}
