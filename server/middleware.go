package server

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func initmiddleware(e *echo.Echo) {
	// always recover
	e.Use(mid.Recover())

	// logging with config
	logFormat := "${remote_ip} - - [${time_custom}] \"${method} ${path} ${protocol}\" ${status} ${bytes_out}\n"
	customTimeFormat := "2/Jan/2006:15:04:05 -0700"

	e.Use(mid.LoggerWithConfig(mid.LoggerConfig{
		Format:           logFormat,
		CustomTimeFormat: customTimeFormat,
	}))

	// request body limit
	e.Use(mid.BodyLimit("5mb"))
}
