package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	
	// adding commands
	commands := []tgbotapi.BotCommand {
		{Command: "start", 	Description: "starting the bot"},
		{Command: "reg", 	Description: "Usage: /reg <username> <password>"},
		{Command: "login", 	Description: "Usage: /login <username> <password>"},
		{Command: "status", Description: "Showing login status"},
	}
	
	cmdConf := tgbotapi.NewSetMyCommands(commands...)
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

			//? test
			//? log.Printf("\nTYPE %T\n TGBOT USER_ID: %s", update, update.SentFrom().UserName)
			
			switch command {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, 
						"Reg or login to use other commands\n" +
						"/reg <username> <password>\n" +
						"/login <username> <password>")

					bot.Send(msg)
				case "reg":
					err := reg(args, db, update)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err)
					bot.Send(msg)
				case "login":
					log.Println("USER LOGGED IN: ", login(args))
				case "status":
					log.Println("LOGGED?: ", isLogged()) 
			}
		}
	}
}

// registration func without db
func reg(args string, db *sql.DB, u tgbotapi.Update) (string) {
	// formatting string into two different strings
	str := strings.Split(args, " ")

	username := str[0]
	password := str[1]
	id := u.SentFrom().UserName
	
	// Adding user into the table
	if strongPassword(password) {
		rows, err := db.Query("INSERT INTO users (username, password, usertg) VALUES ($1, $2, $3)", username, password, id)
		if err != nil {
			log.Fatal(err)
		}
		
		// logs 
		log.Println(rows)
		log.Println("New user registered: \n",
			"\t{username}: ", username, "\n",
			"\t{password}: ", password)	
		
		log.Println("User registred")
		return "User registred"
	}

	log.Println("password isn't strong enough")
	return "password isn't strong enough"
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

// strong password checker
func strongPassword(password string) (bool){
	re := regexp.MustCompile("[0-9]+")
	if len(password) > 8 && re.FindAllString(password, -1) != nil {
		return true
	}

	log.Println("Password isn't strong enought")
	return false
}
