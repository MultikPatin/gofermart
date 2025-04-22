package main

import (
	"main/internal/adapters"
	"main/internal/app"
	"main/internal/config"
	"net/http"
)

func main() {
	logger := adapters.GetLogger()
	defer adapters.SyncLogger()

	c := config.Parse(logger)

	shorterApp, err := app.NewApp(c)
	if err != nil {
		logger.Fatalw(
			"Starting server",
			"error", err.Error(),
		)
	}

	defer func(shorterApp *app.App) {
		err := shorterApp.Close()
		if err != nil {
			logger.Infow(
				"Closing server",
				"error", err.Error(),
			)
		}
	}(shorterApp)

	logger.Infow(
		"Starting server",
		"addr", c.Addr,
	)

	if err := http.ListenAndServe(shorterApp.Addr, shorterApp.Router); err != nil {
		logger.Fatalw(err.Error(), "event", "start server")
	}
}
