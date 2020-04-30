package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
