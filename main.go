package main

import (
	"alertmanager/email"
)

func main() {
	subject := "Alerta de Servidor down"
	server := "Google"
	errorMessage := "Erro ao conectar o servidor."
	time := "29/04/2026"
	emailTemplate := "./email/template.html"

	email.SendMail([]string{"clever.nvs@gmail.com"}, subject, server, errorMessage, time, emailTemplate)
}
