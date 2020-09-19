package NNRequests

import (
	"io/ioutil"
	"log"
	"github.com/imroc/req"
)

type QA struct{
	Contexts []string    `json:"context_raw"`
	Questions []string   `json:"question_raw"`
}

func GetMapAnswers(obj map[string]map[string]string, context string)map[string]string{
	var answers = make(map[string]string)
	for key,val:=range obj{
		answers[key]= getAnswer(QA{Contexts:[]string{context},Questions:[]string{val["question"]}})
	}
	return answers
}


func getAnswer(qa QA)string{
	header := req.Header{
		"Accept":        "application/json",
	}
	param := req.BodyJSON(&qa)

	r, err := req.Post("https://7005.lnsigo.mipt.ru/model", header, param)
	if err != nil {
		log.Fatal(err)
	}

	var s,err2 = ioutil.ReadAll(r.Response().Body)
	if err!=nil{
		panic(err2)
	}

	return string(s)
}
