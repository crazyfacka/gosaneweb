package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	fmt.Printf("%+v\n", sh.Devices())
	return c.String(http.StatusOK, "test")
}
