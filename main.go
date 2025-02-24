package main

import (
	"log"

	"goBot/tg"
	"goBot/tgsql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func main() {
	// dataBase + bot initialization
	tgsql.DBInit()
	tg.BotInitTg()

	tgsql.DB.Stats()


	//? main loop
	for update := range tg.Updates {
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

					tg.Bot.Send(msg)
				case "reg":
					err := tgsql.Reg(args, update)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err)
					tg.Bot.Send(msg)
				case "login":
					log.Println("USER LOGGED IN: ", tgsql.Login(args))
				case "status":
					log.Println("LOGGED?: ", tgsql.IsLogged()) 
			}
		}
	}
}
