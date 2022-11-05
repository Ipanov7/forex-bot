package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"lorisocchipinti.com/gbp-rates/logger"
)

// Only for demo purposes. Never hardcode sensitive stuff!
const (
	tgApiKey = "<YOUR-TELEGRAM-BOT-APIKEY>"
	tgChatID = "<YOUR-TELEGRAM-CHAT-ID>"
)

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(message string) {
	url := "https://api.telegram.org/bot" + tgApiKey + "/sendMessage"
	payload, _ := json.Marshal(TelegramMessage{ChatID: tgChatID, Text: message})

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	logger.Log(fmt.Sprintf("Message sent with status %s", res.Status))
}
