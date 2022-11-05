package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"lorisocchipinti.com/gbp-rates/logger"
	"lorisocchipinti.com/gbp-rates/telegram"
)

// Only for demo purposes. Never hardcode sensitive stuff!
const (
	apikey = "<YOUR-API-KEY>"
)

type Quote struct {
	Result float64
}

type RateRequest struct {
	From        string  `json:"from"`
	To          string  `json:"to"`
	AverageRate float64 `json:"avg_rate"`
}

func main() {
	lambda.Start(CheckRate)
}

func CheckRate(ctx context.Context, request RateRequest) {
	if time.Now().Hour() == 9 {
		logger.Log("Sending daily wake-up message...")
		telegram.SendMessage("Good morning, bee ready to make money today! 🐝 💸 🐝")
	}
	logger.Log("Fetching rate...")
	currentRate := fetchRate(request.From, request.To)
	pips := computePips(currentRate, request.AverageRate)
	logger.Log(fmt.Sprintf("Current rate is %.4f. Possible gain/loss is %d PIPs", currentRate, pips))
	logger.Log("Sending alert to user...")
	if pips > 0 {
		telegram.SendMessage(fmt.Sprintf("🟢 Good news! Price is up %d PIPs with the current %s%s rate %.4f 🤑 🤑 🤑", pips, request.From, request.To, currentRate))
	} else {
		telegram.SendMessage(fmt.Sprintf("🔴 No luck, price is down %d PIPs with the current %s%s rate %.4f 😰", pips, request.From, request.To, currentRate))
	}
}

func fetchRate(from string, to string) float64 {
	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%s&from=%s&amount=1", to, from)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err)
	}

	req.Header.Set("apikey", apikey)
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
	}
	quote, err := parseQuote(body)
	if err != nil {
		logger.Error(err)
	}
	return quote.Result
}

func parseQuote(data []byte) (Quote, error) {
	var quote Quote
	err := json.Unmarshal(data, &quote)
	if err != nil {
		return Quote{}, err
	}
	return quote, nil
}

func computePips(currentRate float64, averageRate float64) int {
	return int((currentRate - averageRate) * 10000)
}
