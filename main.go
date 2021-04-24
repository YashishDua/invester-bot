package main

import (
	"fmt"
	"invester-bot/config"
	"invester-bot/middlewares"
	"net/http"
)

func main() {
	appConfig, err := config.GetConfiguration()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	router := middlewares.GetRouter()

	http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), router)
}
