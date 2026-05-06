package main

import (
	"alertmanager/email"
	"alertmanager/slack"
	"alertmanager/sms"
	"alertmanager/telegram"
	"os"
)

func main() {
	// MENSAGEM POR EMAIL
	// subject := "Alerta de Servidor down"
	// server := "Google"
	// errorMessage := "Erro ao conectar o servidor."
	// time := "29/04/2026"
	// emailTemplate := "./email/template.html"

	// email.SendMail([]string{"clever.nvs@gmail.com"}, subject, server, errorMessage, time, emailTemplate)

	// MENSAGEM POR SLACK
	// message := "Alerta de Servidor down"
	// slack.SendMessage(message)

	// message := "Alerta de Servidor down"
	// phone := "55991999999"
	// sms.SendMessage(message, phone)

	// MENSAGEM POR TELEGRAM
	botAPI := os.Getenv("TELEGRAM_BOT_API")
	telegram.SendMessage(botAPI, "Alerta de Servidor down")
}
