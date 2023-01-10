package server

import (
	"github.com/labstack/echo/v4"
	modules "gopps/modules/auth"
	"gopps/modules/utils"
)

// if you need to add more data, just add more fields to this struct
type Data struct {
	Ip string
}

func index(c echo.Context) error {
	data := &Data{
		Ip: utils.IPaddress(),
	}
	return c.Render(200, "index", data)
}

// always render adminlogin, until the password and the username are correct
func adminboard(c echo.Context) error {

	adminname := c.FormValue("adminname")
	adminpass := c.FormValue("adminpass")

	// change password and username here //
	username := "admin"
	password := "admin"

	if !modules.Authvalidator(adminname, adminpass, username, password) {
		return c.Render(401, "adminlogin", nil)
	}
	return c.Render(200, "adminboard", nil)
}

func Routing(e *echo.Echo) {
	// Routes
	e.GET("/", index)

	e.GET("/admin", adminboard)
	e.POST("/admin", adminboard)

	e.Static("/static/", "static")
}
