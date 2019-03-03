package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(YOUR_KEY_API)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Command() {
		case "time":
			resp, err := http.Get("http://worldtimeapi.org/api/ip")
			if nil != err {
				fmt.Println("error connection", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			if nil != err {
				fmt.Println("errorination happened reading the body", err)
				return
			}
			fmt.Println(string(body[:]))
			msg.Text = string(body)
		case "status":
			msg.Text = "I'm ok."
		case "name":
			msg.Text = update.Message.From.UserName
		default:
			msg.Text = "Hello, lets play with me \n i have command /time \n/name \n/status"
		}
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
