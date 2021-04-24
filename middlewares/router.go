package middlewares

import (
	"invester-bot/apis"

	"github.com/go-chi/chi"
)

func GetRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(ResponseMiddleware)

	router.Get("/knockknock", apis.GetHealthHandler)
	router.Get("/trades", apis.GetTradesHandler)

	return router
}
