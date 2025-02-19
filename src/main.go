package main

import (
	"log"
	"os"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	log.Println("Bot token is: " + token)

	bot, err := tgbot.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	commands := []tgbot.BotCommand {
		{Command: "start", 	Description: "starting the bot"},
		{Command: "reg", 	Description: "Usage: /reg <username> <password>"},
		{Command: "login", 	Description: "Usage: /login <username> <password>"},
		{Command: "status", Description: "Showing login status"},
	}

	cmdConf := tgbot.NewSetMyCommands(commands...)
	_, err = bot.Request(cmdConf)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			command := update.Message.Command()
			args := update.Message.CommandArguments()

			switch command {
				case "start":
					msg := tgbot.NewMessage(update.Message.Chat.ID, 
						"Reg or login to use other commands\n" +
						"/reg <username> <password>\n" +
						"/login <username> <password>")

					bot.Send(msg)
				case "reg":
					reg(args)
				case "login":
					log.Println("USER LOGGED IN: ", login(args))
				case "status":
					log.Println("LOGGED?: ", isLogged())
			}
		}
	}
}

//! REMAKE FUNC WITH USING DB
// registration func without db
func reg(args string) (string, string) {
	str := strings.Split(args, " ")

	username := str[0]
	password := str[1]
	
	// logs 
	log.Println("New user registered: \n",
		"\t{username}: ", username, "\n",
		"\t{password}: ", password)	

	return username, password
}

//! REMAKE FUNC WITH USING DB
// login func without db
func login(args string) (bool) {
	strings.Split(args, "")
	return true
}

//! REMAKE FUNC WITH USING DB
// func should return username of logged person
func isLogged() (bool) {
	return true
}