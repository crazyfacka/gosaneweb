package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func devices(c echo.Context) error {
	devices := sh.Devices()
	return c.JSON(http.StatusOK, devices)
}
