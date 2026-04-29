package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func SendMail(to []string, subject string, serverName string, errorMessage string, time string, templatePath string) {
	from := os.Getenv("EMAIL_CLIENT")
	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		panic("Senha do e-mail não configurada")
		os.Exit(1)
	}

	smtpHost := os.Getenv("SMTP_CLIENT")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles(templatePath)

	var emailBody bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset-\"UTF-8\";\n\n"
	emailBody.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	t.Execute(&emailBody, struct {
		Server string
		Error  string
		Time   string
	}{
		Server: serverName,
		Error:  errorMessage,
		Time:   time,
	})

	err := smtp.SendMail(smtpHost+""+smtpPort, auth, from, to, emailBody.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar o email: %s", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")
}
