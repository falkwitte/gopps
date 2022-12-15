package server

import (
	"github.com/labstack/echo/v4"
	"gopps/modules/utils"
	"html/template"
)

func Startserver() {
	// initializing echo
	e := echo.New()
	t := template.New("")

	// rendering
	renderer := &TemplateRenderer{
		templates: template.Must(t.ParseGlob("templates/*.gohtml")),
	}
	e.Renderer = renderer

	Routing(e)
	initmiddleware(e)

	e.Logger.Fatal(e.Start(utils.IPaddress()))
}
