package tgsql

import (
	"log"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// registration func without db
func Reg(args string, u tgbotapi.Update) (string) {
	// formatting string into two different strings
	str := strings.Split(args, " ")

	username := str[0]
	password := str[1]
	id := u.SentFrom().UserName
	
	// Adding user into the table
	if strongPassword(password) {
		rows, err := DB.Query("INSERT INTO users (username, password, usertg) VALUES ($1, $2, $3)", username, password, id)
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
func Login(args string) (bool) {
	strings.Split(args, "")
	return true
}

//! REMAKE FUNC WITH USING DB
// func should return username of logged person
func IsLogged() (bool) {
	
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
