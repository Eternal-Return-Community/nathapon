package utils

import (
	"log"
	"nathapon/src/models"
	"os"

	"github.com/joho/godotenv"
)

var Env models.Env

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[Error] - .env file not found.")
	}

	Env = models.Env{
		Auth:     checkENV("AUTH", "\n-> OAuth is missing in .env!"),
		Token:    checkENV("TOKEN", "\n-> Token is missing in .env!"),
		Client:   checkENV("CLIENT", "\n-> Client-ID is missing in .env!"),
		Database: checkENV("DATABASE", "\n-> Database URL is missing in .env!"),
		Webhook:  checkENV("WEBHOOK", "\n-> Webhook URL is missing in .env!"),
	}
}

func checkENV(name string, message string) string {
	if env := os.Getenv(name); env != "" {
		return env
	}

	log.Fatalln(message)
	return ""
}
