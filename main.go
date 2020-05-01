package main

import (
	"os"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/crazyfacka/gosaneweb/repository"
	"github.com/crazyfacka/gosaneweb/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := domain.LoadConfiguration(); err != nil {
		log.Error().Err(err).Msg("Error reading config file")
		os.Exit(-1)
	}

	scanHandler, err := repository.InitScanImage()
	if err != nil {
		log.Error().Err(err).Msg("Error getting scan handler")
		os.Exit(-1)
	}

	routes.Start(scanHandler)
}
