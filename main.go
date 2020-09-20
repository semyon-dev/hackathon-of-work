package main

import (
	"github.com/hacktrud/hackathon-work/api"
)

func main() {
	//config.Load()
	//db.Connect()
	//mongo.Connect()
	api.RunAPI()

	//db.GetStoreAnswers()
	//done := make(chan struct{})
	//go db.GetRestaurantsAnswers(done)
	//<-done
	//db.GetDriversAnswers()
}
