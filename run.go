package main

import (
	"github.com/catsworld/qq-bot-api"
	// "github.com/catsworld/qq-bot-api/cqcode"
	"github.com/yautah/bot/message"
	"log"
	"net/http"
	"strings"
)

const (
	clanTag = "cc8gr0q"
	tag     = ""
)

func startListen() {
	bot, err := qqbotapi.NewBotAPI("abcdefghijklmn", "http://192.168.1.69:5700", "abcdefghijklmn")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true
	updates := bot.ListenForWebhook(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		log.Println("wtffffff", update.Message)
		if update.Message == nil {
			continue
		}

		if update.UserID == 749594 && strings.Contains(update.Message.Text, "战报") {
			// fetchWars()
			msg := message.NewWarMessage(clanTag)
			// bot.SendMessage(2434861, "group", msg)
			bot.SendMessage(13285012, "group", msg)
			// bot.SendMessage(267535, "group", message)
		}

		if update.UserID == 749594 && strings.Contains(update.Message.Text, "查宝箱") {
			message.SendChestMessage("gvqygqyr", *bot)
		}
	}
}

func main() {
	startListen()
	// draw.CreateImg()
}
