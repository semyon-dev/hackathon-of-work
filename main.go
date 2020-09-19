package main

import (
	"hackathon-work/api"
	"hackathon-work/db"
	"hackathon-work/mongo"
)

func main() {
	db.Connect()
	mongo.Connect()
	api.RunAPI()
}
