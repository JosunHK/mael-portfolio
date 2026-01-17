package dummy

import (
	"mael/cmd/layout"
	"mael/cmd/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//dummy routes
	e.GET("/", middleware.Pages(layout.Layout, Dummy))
}
