package main

import (
	"stakeway/config"
	"stakeway/internal/app"
	"stakeway/pkg/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.New()
	if err != nil {

	}

	a := app.New(cfg, log)
	if err = a.Run(); err != nil {
		log.Errorf("error running app: %s", err)
	}
}
