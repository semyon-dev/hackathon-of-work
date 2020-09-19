package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hackathon-work/NNRequests"
	"io/ioutil"
	"log"
	"net/http"
)

func RunAPI() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/resume", func(c *gin.Context) {

		jsonInput := struct {
			Competences string `json:"competences"`
			Resume      string `json:"resume"`
		}{}

		if err := c.ShouldBindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error: can't bind json",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "ok",
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

	err := r.Run(":5555")
	if err != nil {
		log.Println(err)
	}
}
