package routes

import (
	"net/http"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func scan(c echo.Context) error {
	device := new(domain.Device)
	if err := c.Bind(device); err != nil {
		log.Info().Err(err).Msg("Bad scan request")
		return c.NoContent(http.StatusBadRequest)
	}

	log.Debug().Interface("data", device).Msg("Scan request")

	file, err := sh.Scan(*device)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Blob(http.StatusOK, "image/png", file)
}
