package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var MongoDBPassword string

// загрузка конфигов
func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("× Ошибка загрузки .env файла: " + err.Error())
	} else {
		MongoDBPassword = os.Getenv("MONGODB_PASSWORD")
	}
}
