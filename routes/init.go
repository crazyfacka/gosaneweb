package routes

import (
	"html/template"
	"io"
	"strconv"

	"github.com/crazyfacka/gosaneweb/domain"
	"github.com/crazyfacka/gosaneweb/repository"
	"github.com/labstack/echo"
)

var sh repository.Scan

// Template the main struct defining the usage of templates
type Template struct {
	templates *template.Template
}

// Render creates the final HTML of the template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Start loads all the routes and starts the webserver
func Start(scanHandler repository.Scan) {
	sh = scanHandler

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.GET("/", index)
	e.GET("/devices", devices)
	e.POST("/scan", scan)

	if domain.Confs.Debug {
		e.Debug = true
		e.GET("/test", test)
	}

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(domain.Confs.Port)))
}
