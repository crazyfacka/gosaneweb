package routes

import (
	"net/http"

	"github.com/crazyfacka/gosaneweb/domain"

	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	devices := sh.Devices()

	dev := devices[0]
	dev.Ft[domain.MODE].ToUse = "Color"
	dev.Ft[domain.RESOLUTION].ToUse = "100"

	file, err := sh.Scan(dev)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Blob(http.StatusOK, "image/png", file)
}
