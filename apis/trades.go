package apis

import (
	"encoding/json"
	"fmt"
	"invester-bot/config"
	"net/http"
)

// TODO: To be injected globally or per request level
var AppConfig, _ = config.GetConfiguration()

type CoinAPIResponse struct {
	Status struct {
	} `json:"status"`
	Data []CoinAPIData `json:"data"`
}

type CoinAPIData struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  struct {
		INR struct {
			Price        float32 `json:"price"`
			Vol24h       float32 `json:"volume_24h"`
			PerChange1h  float32 `json:"percent_change_1h"`
			PerChange24h float32 `json:"percent_change_24h"`
			PerChange7d  float32 `json:"percent_change_7d"`
		} `json:"INR"`
	} `json:"quote"`
}

type Trade struct {
	Name   string `json:"name"`
	Action string `json:"action"`
	Value  int    `'json:"value"`
}

// TODO: Refactor! If you have reached this place, do not judge me.
func getMediumTrades() ([]Trade, error) {
	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?start=1&limit=%d&convert=INR&sort=volume_24h",
		AppConfig.CoinAPI.BaseUrl, AppConfig.Trade.MediumRisk.Bucket)

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-CMC_PRO_API_KEY", AppConfig.CoinAPI.Token)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &CoinAPIResponse{}

	json.NewDecoder(res.Body).Decode(response)

	var mediumRiskTrades []Trade

	for i := 0; i < len(response.Data); i++ {
		data := response.Data[i]

		if data.Quote.INR.PerChange1h >= float32(AppConfig.Trade.MediumRisk.PerChange1h) {
			mediumRiskTrades = append(mediumRiskTrades, Trade{
				Name:   data.Name,
				Action: "buy",
				Value:  AppConfig.Trade.MediumRisk.Value,
			})
		}
	}

	return mediumRiskTrades, nil
}

func GetTradesHandler(w http.ResponseWriter, r *http.Request) {
	trades, err := getMediumTrades()

	if err != nil {
		HandleError(w, InternalServerError(err))

		return
	}

	HandleResponse(w, trades)
}
