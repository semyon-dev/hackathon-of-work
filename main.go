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

	//db.GetStoreAnswers()
	//done := make(chan struct{})
	//go db.GetRestaurantsAnswers(done)
	//<-done
	//db.GetDriversAnswers()
}
