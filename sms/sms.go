package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func SendMessage(message string, phone string) {

	endpoint := "https://rest.nexmo.com/sms/json"
	data := url.Values{}
	apiKey := os.Getenv("VONAGE_API_KEY")
	if apiKey == "" {
		panic("API Key do Vonage não configurada")
	}

	apiSecret := os.Getenv("VONAGE_API_SECRET")
	if apiSecret == "" {
		panic("API Secret do Vonage não configurada")
	}
	data.Set("api_key", apiKey)
	data.Set("api_secret", apiSecret)
	data.Set("from", "Sistema de Monitoramento")
	data.Set("text", message)
	data.Set("to", phone)
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	log.Printf("Mensagem enviada para %s com status %s", phone, resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}
