package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	sh.Devices()
	return c.String(http.StatusOK, "test")
}
