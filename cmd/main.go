package main

import (
	"log"
	"moon-telegrambot/internal/telegram"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Telegram Bot Start ::" + port)
	telegram.RunServer(":" + port)
}
