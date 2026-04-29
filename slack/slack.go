package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendMessage(message string) {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("Token do Slack não configurado")
		os.Exit(1)
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("ID do canal do Slack não configurado")
		os.Exit(1)
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   danger,
		Pretext: "Alerta de Servidor down",
		Text:    message,
	}
	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensagem enviada com sucesso %s as %s", channelID, timestamp)

}
