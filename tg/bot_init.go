package tg

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	Updates tgbotapi.UpdatesChannel
	Bot 	*tgbotapi.BotAPI
)

func BotInitTg() {	
	// loading env + starting bot
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	log.Println("Bot token is: " + token)

	var err error
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	Bot.Debug = true

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	Updates = Bot.GetUpdatesChan(u)
	
	// adding commands
	commands := []tgbotapi.BotCommand {
		{Command: "start", 		Description: "starting the bot"},
		{Command: "reg", 		Description: "Usage: /reg <username> <password>"},
		{Command: "login", 		Description: "Usage: /login <username> <password>"},
		{Command: "status", 	Description: "Showing login status"},
		{Command: "download", 	Description: "Download by link"},
	}
	
	cmdConf := tgbotapi.NewSetMyCommands(commands...)
	_, err = Bot.Request(cmdConf)
	if err != nil {
		log.Panic(err)
	}
}