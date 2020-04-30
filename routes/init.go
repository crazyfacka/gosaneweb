package routes

import (
	"strconv"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/labstack/echo"
)

// Start loads all the routes and starts the webserver
func Start() {
	e := echo.New()

	if domain.Confs.Debug {
		e.GET("/test", test)
	}

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(domain.Confs.Port)))
}
