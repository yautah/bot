package main

import (
	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	. "github.com/yautah/bot/data"
	"github.com/yautah/bot/message"
	"log"
	"net/http"
	"strconv"
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
	u.PreloadUserInfo = false
	updates := bot.ListenForWebhook(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		cqcode.StrictCommand = true

		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text, update.Message.IsCommand())

		if update.Message.IsCommand() {
			cmd, args := update.Message.Command()
			log.Println(cmd, args)
			switch cmd {
			case "bind":
				if len(args) == 0 {
					message.SendMessage(*bot, update.GroupID, strconv.FormatInt(update.UserID, 10), " 请输入要绑定的玩家TAG，格式：/bind xxxxxxx")
					continue
				}

				u := User{Qq: strconv.FormatInt(update.UserID, 10), Tag: args[0]}
				err := CreateUser(&u)
				if err != nil {
					message.SendMessage(*bot, update.GroupID, strconv.FormatInt(update.UserID, 10), " 绑定失败")
				} else {
					message.SendMessage(*bot, update.GroupID, strconv.FormatInt(update.UserID, 10), " 绑定成功!输入 /chest 就可以快乐的查询宝箱了~~~~")
				}

			case "chest":
				u := FindUserByQq(strconv.FormatInt(update.UserID, 10))
				if u != nil {
					message.SendChestMessage(*bot, update.GroupID, strconv.FormatInt(update.UserID, 10), u.Tag)
				} else {
					message.SendMessage(*bot, update.GroupID, strconv.FormatInt(update.UserID, 10), " 请先绑定你的玩家TAG")
				}
			default:
				message.SendCommandTip(*bot, update.GroupID)
			}

		}

		if update.UserID == 749594 && strings.Contains(update.Message.Text, "战报") {
			// fetchWars()
			msg := message.NewWarMessage(clanTag)
			// bot.SendMessage(2434861, "group", msg)
			bot.SendMessage(13285012, "group", msg)
			// bot.SendMessage(267535, "group", message)
		}

	}
}

func main() {
	startListen()
	// draw.CreateImg()
}
