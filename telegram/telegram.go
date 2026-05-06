package telegram

import (
	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"strconv"
)

type Message struct {
	Text    string `json:"text"`
	GroupID int64  `json:"group_id"`
}

func SendMessage(botAPI string, message string) {
	bot, err := telegramBot.NewBotAPI(botAPI)
	if err != nil {
		panic(err)
	}
	telegramMessage := Message{}
	telegramMessage.Text = message
	groupID := os.Getenv("TELEGRAM_GROUP_ID")
	if groupID == "" {
		panic("TELEGRAM_GROUP_ID não definido")
	}
	telegramMessage.GroupID, err = strconv.ParseInt(groupID, 10, 64)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	alertText := telegramBot.NewMessage(telegramMessage.GroupID, telegramMessage.Text)
	bot.Send(alertText)

}
