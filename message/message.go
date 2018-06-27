package message

import (
	"fmt"
	"github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	. "github.com/yautah/bot/draw"
	. "github.com/yautah/bot/network"
	// "log"
)

const (
	warHeader = "东风吹，战鼓擂，皇室究竟谁怕谁？ \r\r"
	warWin    = "恭喜 【%s】 在战斗日中战胜了 【%s】, 结果 %d : %d，请收下我的膝盖！\r"
	warLose   = "很遗憾... 【%s】 在战斗日中败给了 【%s】, 结果 %d : %d，下次再接再厉哟！\r"
	warDraw   = "很遗憾.... 【%s】 在战斗日中和 【%s】 打成了平手, 结果 %d : %d，下次再接再厉哟！\r"

	chestHeader = "小蜜掐指一算，你接下来的箱子是：\r"
)

func NewWarMessage(tag string) cqcode.Message {

	results := FetchWars(tag)

	message := make(cqcode.Message, 0)

	for _, v := range results {
		message.Append(&cqcode.Text{Text: warHeader})

		switch v.Winner {
		case 1:
			message.Append(
				&cqcode.Text{Text: fmt.Sprintf(warWin, v.Team[0].Name, v.Opponent[0].Name, v.TeamCrowns, v.OpponentCrowns)})
		case -1:
			message.Append(&cqcode.Text{Text: fmt.Sprintf(warLose, v.Team[0].Name, v.Opponent[0].Name, v.TeamCrowns, v.OpponentCrowns)})
		case 0:
			message.Append(&cqcode.Text{Text: fmt.Sprintf(warDraw, v.Team[0].Name, v.Opponent[0].Name, v.TeamCrowns, v.OpponentCrowns)})
		}
		message.Append(&cqcode.Text{Text: fmt.Sprintf("\r出战卡组\r")})

		b := CreateDeckImg(v.Team[0].Deck, v.Opponent[0].Deck)
		image, _ := qqbotapi.NewImageBase64(b)
		message.Append(image)
		message.Append(&cqcode.Text{Text: "\r"})

		message.Append(&cqcode.Text{Text: fmt.Sprintf("\r卡组链接: \r")})
		message.Append(&cqcode.Text{Text: v.Team[0].DeckLink})

	}
	return message
}

func SendChestMessage(tag string, bot qqbotapi.BotAPI) {

	msg := make(cqcode.Message, 0)
	msg.Append(&cqcode.Text{Text: "天灵灵~~~~~~~~~地灵灵~~~~~~~~~"})

	bot.SendMessage(2434861, "group", msg)

	chest := FetchChest(tag)

	message := make(cqcode.Message, 0)
	message.Append(&cqcode.Text{Text: chestHeader})

	b := CreateChestImg(chest)
	image, _ := qqbotapi.NewImageBase64(b)
	message.Append(image)
	bot.SendMessage(2434861, "group", message)

}

func NewTextMessage(text string) cqcode.Message {

	message := make(cqcode.Message, 0)
	message.Append(&cqcode.Text{Text: text})

	return message
}
