package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "World")
}
