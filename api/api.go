package api

import (
	"github.com/gin-gonic/gin"
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

	err := r.Run(":5555")
	if err != nil {
		log.Println(err)
	}
}
