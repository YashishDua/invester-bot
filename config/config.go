package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Port    int `json:"port"`
	CoinAPI struct {
		BaseUrl string `json:"base_url"`
		Token   string `envconfig:"COIN_TOKEN", json:"token"`
	} `json:"coinmarketcap_api"`
	Trade struct {
		MediumRisk struct {
			Bucket       int `json:"bucket"`
			PerChange1h  int `json:"percent_change_1h"`
			PerChange24h int `json:"percent_change_24h"`
			Value        int `json:"value"`
		} `json:"medium_risk"`
	} `json:"trade"`
}

// TODO: Test and production variants to be introduced
func GetConfiguration() (Configuration, error) {
	configuration := Configuration{}

	configFile, err := ioutil.ReadFile("./config/development.json")

	if err != nil {
		return configuration, err
	}

	// Merging statis json into config
	if err = json.Unmarshal(configFile, &configuration); err != nil {
		return configuration, err
	}

	// Merging env variables into config
	if err = envconfig.Process("", &configuration); err != nil {
		return configuration, err
	}

	return configuration, nil
}
