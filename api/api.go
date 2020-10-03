package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hackathon-work/NNRequests"

	"io/ioutil"
	"log"
	"net/http"
)

func RunAPI() {

	var jsonModelsPath = "hackathon-work/data/jsonModels/"

	fileNames, err := ioutil.ReadDir(jsonModelsPath)
	for i := range fileNames {
		fmt.Println(fileNames[i].Name())
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.Static("static", "./hackathon-work/static")
	r.Static("js", "./hackathon-work/static/dist/js")
	r.Static("css", "./hackathon-work/static/dist/css")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/resume", func(c *gin.Context) {

		fmt.Println(" in resume")

		jsonInput := struct {
			Competences string `json:"competention"`
			Resume      string `json:"context"`
		}{}

		err = c.ShouldBindJSON(&jsonInput)
		if err != nil {
			return
		}

		//fmt.Println("------")
		//fmt.Println(jsonInput.Competences, jsonInput.Resume)

		jsonStructData, err := ioutil.ReadFile(jsonModelsPath + jsonInput.Competences + ".json")

		dictQuestions := make(map[string]map[string]string)

		err = json.Unmarshal(jsonStructData, &dictQuestions)
		if err != nil {
			log.Println(err)
		}

		var answers = NNRequests.GetMapAnswers(dictQuestions, jsonInput.Resume)
		returnedJson, err := json.Marshal(answers)

		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
			"data":    string(returnedJson),
		})
	})

	r.POST("/answers", func(ctx *gin.Context) {
		body, err := ctx.Request.GetBody()
		if err != nil {
			log.Println(err)
		}
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			log.Println(err)
		}
		dictQuestions := make(map[string]map[string]string)
		jsonData := make(map[string]interface{})
		err = json.Unmarshal(bytes, &jsonData)
		dictQuestions = jsonData["questions"].(map[string]map[string]string)

		answers := NNRequests.GetMapAnswers(dictQuestions, jsonData["context"].(string))

		jsonStr, err := json.Marshal(answers)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, gin.H{
			"data": string(jsonStr),
		})

	})

	err = r.Run(":5000")
	if err != nil {
		log.Println(err)
	}
}
