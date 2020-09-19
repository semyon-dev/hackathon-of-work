package main

import (
	"hackathon-work/config"
	"hackathon-work/db"
	"hackathon-work/mongo"
)

func main() {
	config.Load()
	db.Connect()
	mongo.Connect()
	//api.RunAPI()

	db.GetStoreAnswers()
}
