package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	APP_TOKEN := os.Getenv("APP_TOKEN")
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
}
