package routes

import (
	"net/http"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/labstack/echo"
)

type valsDefault struct {
	Values  []string
	Default string
}

type rangedDefault struct {
	Min     string
	Max     string
	Default string
}

type renderData struct {
	Name       string
	Modes      valsDefault
	Resolution valsDefault
	Brightness rangedDefault
	Contrast   rangedDefault
	X          rangedDefault
	Y          rangedDefault
	Top        rangedDefault
	Left       rangedDefault
}

func index(c echo.Context) error {
	device := sh.Devices()[0]

	data := &renderData{
		Name: device.Name,
		Modes: valsDefault{
			Values:  device.Ft[domain.MODE].Values,
			Default: device.Ft[domain.MODE].Default,
		},
		Resolution: valsDefault{
			Values:  device.Ft[domain.RESOLUTION].Values,
			Default: device.Ft[domain.RESOLUTION].Default,
		},
		Brightness: rangedDefault{
			Min:     device.Ft[domain.BRIGHTNESS].Values[0],
			Max:     device.Ft[domain.BRIGHTNESS].Values[1],
			Default: device.Ft[domain.BRIGHTNESS].Default,
		},
		Contrast: rangedDefault{
			Min:     device.Ft[domain.CONTRAST].Values[0],
			Max:     device.Ft[domain.CONTRAST].Values[1],
			Default: device.Ft[domain.CONTRAST].Default,
		},
		X: rangedDefault{
			Min:     device.Ft[domain.X].Values[0],
			Max:     device.Ft[domain.X].Values[1],
			Default: device.Ft[domain.X].Default,
		},
		Y: rangedDefault{
			Min:     device.Ft[domain.Y].Values[0],
			Max:     device.Ft[domain.Y].Values[1],
			Default: device.Ft[domain.Y].Default,
		},
		Top: rangedDefault{
			Min:     device.Ft[domain.T].Values[0],
			Max:     device.Ft[domain.T].Values[1],
			Default: device.Ft[domain.T].Default,
		},
		Left: rangedDefault{
			Min:     device.Ft[domain.L].Values[0],
			Max:     device.Ft[domain.L].Values[1],
			Default: device.Ft[domain.L].Default,
		},
	}

	return c.Render(http.StatusOK, "index.html", data)
}
