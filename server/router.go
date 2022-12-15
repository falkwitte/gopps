package server

import (
	"fmt"
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

func admin(c echo.Context) error {
	return c.Render(200, "admin", utils.IPaddress())
}

// always render admin, until the password and the username are correct
func adminboard(c echo.Context) error {

	adminname := c.FormValue("adminname")
	adminpass := c.FormValue("adminpass")

	// change password and username here //
	username := "admin"
	password := "admin"

	fmt.Printf("adminname: ", adminname)
	fmt.Println("\n")

	if !modules.Authvalidator(adminname, adminpass, username, password) {
		return c.Render(401, "admin", nil)
	}
	return c.Render(200, "adminboard", nil)
}

func testwasm(c echo.Context) error {
	return c.Render(200, "testwasm", nil)
}

func Routing(e *echo.Echo) {
	// Routes
	e.GET("/", index)

	e.GET("/admin", admin)

	e.GET("/adminboard", adminboard)
	e.POST("/adminboard", adminboard)

	e.GET("/testwasm", testwasm)

	e.Static("/static/", "static")
}
