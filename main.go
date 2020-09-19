package main

import (
	"hackathon-work/api"
	"hackathon-work/db"
)

func main() {
	db.Connect()
	api.RunAPI()
}
