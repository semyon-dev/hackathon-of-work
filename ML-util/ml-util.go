package ML_util

import "hackathon-work/model"

// файл с функциями для генерации вопросов к ML в нужном формате

import (
	"encoding/json"
	"fmt"
	"hackathon-work/mongo"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var res model.Res
var a int

func CreateAnswer(context, question, id string) {
	var data model.Data
	data.Title = ""
	var a = []model.Answer{}
	qas := model.Qas{
		Answers:  a,
		Question: question,
		Id:       id,
	}
	var Qas []model.Qas
	var Paragraphs []model.Paragraph
	Qas = append(Qas, qas)
	Paragraph := model.Paragraph{
		Context: context,
		Qas:     Qas,
	}
	Paragraphs = append(Paragraphs, Paragraph)
	data.Paragraphs = Paragraphs
	res.Data = append(res.Data, data)
}

func FINISH() {
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	// write to JSON file
	a++
	jsonFile, err := os.Create("questions" + strconv.Itoa(a) + ".json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	_, err = jsonFile.Write(resJson)
	if err != nil {
		log.Println(err)
	}
	err = jsonFile.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("JSON data written to ", jsonFile.Name())
	res.Data = nil
}

type Key struct {
	Text        string  `json:"text"`
	Probability float64 `json:"probability"`
	StartLogit  float64 `json:"start_logit"`
	EndLogit    float64 `json:"end_logit"`
}

type Reply map[string][]Key //json  key = id + q

func ReadFile() {
	var plan, err = ioutil.ReadFile("pred.json")
	if err != nil {
		log.Println(err)
	}
	var data Reply
	var kkk map[string][]Key
	data = kkk
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Println(err)
	}
	var a int
	var restaurant model.CandidateRestaurant
	for key, v := range data {
		for l := 0; l < len(v); l++ {
			if v[l].Probability >= 0.70 {
				a++
				mas := strings.Split(key, "_")
				in, err := strconv.Atoi(mas[0])
				if err != nil {
					log.Println(err)
				}
				restaurant.Id = int64(in)
				fmt.Println(restaurant.Id)
				mongo.UpdateRestaurantsAnswers(restaurant, "q"+mas[1], v[l].Text)
			}
		}
	}
	fmt.Println(a)
}

func ReadFileStores() {
	var plan, err = ioutil.ReadFile("pred1-5.json")
	if err != nil {
		log.Println(err)
	}
	var data Reply
	var kkk map[string][]Key
	data = kkk
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Println(err)
	}
	var a int
	var restaurant model.CandidateStore
	for key, v := range data {
		for l := 0; l < len(v); l++ {
			if v[l].Probability >= 0.70 {
				a++
				mas := strings.Split(key, "_")
				in, err := strconv.Atoi(mas[0])
				if err != nil {
					log.Println(err)
				}
				restaurant.Id = int64(in)
				fmt.Println(restaurant.Id)
				mongo.UpdateStoreAnswers(restaurant, "q"+mas[1], v[l].Text)
			}
		}
	}
	fmt.Println(a)
}
