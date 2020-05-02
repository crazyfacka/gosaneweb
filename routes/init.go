package routes

import (
	"strconv"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/crazyfacka/gosaneweb/repository"
	"github.com/labstack/echo"
)

var sh repository.Scan

// Start loads all the routes and starts the webserver
func Start(scanHandler repository.Scan) {
	e := echo.New()

	sh = scanHandler

	if domain.Confs.Debug {
		e.GET("/test", test)
	}

	e.GET("/devices", devices)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(domain.Confs.Port)))
}
