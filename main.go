package main

import (
	"os"

	bot "burry.trade/discord_bot/Bot"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot.BotToken = os.Getenv("TOKEN")
	bot.Run() // call the run function of bot/bot.go
}
