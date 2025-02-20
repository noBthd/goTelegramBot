package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// db connetion
	connStr := "user=postgres password=1234 dbname=tgBot sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n")

	// loading env + starting bot
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

	// adding commands
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

	//? main loop
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
					reg(args, db)
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
func reg(args string, db *sql.DB) {
	// formatting string into two different strings
	str := strings.Split(args, " ")

	username := str[0]
	password := str[1]
	
	// Adding user into the table
	rows, err := db.Query("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)

	// logs 
	log.Println("New user registered: \n",
		"\t{username}: ", username, "\n",
		"\t{password}: ", password)	
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

func strongPassword(password string) (bool){
	if len(password) > 8 {
		return true
	}

	log.Println("Password isn't strong enought")
	return false
}
