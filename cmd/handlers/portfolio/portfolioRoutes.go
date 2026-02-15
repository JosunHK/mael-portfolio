package portfolio

import (
	"github.com/labstack/echo/v4"
	"mael/cmd/layout"
	"mael/cmd/middleware"
	"mael/web/templates/contents/portfolio"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", middleware.Pages(layout.Layout, Animations))
	e.GET("/animation", middleware.Pages(layout.Layout, Animations))
	e.GET("/animation/body", middleware.HTMX(Animations))
	e.GET("/characters", middleware.Pages(layout.Layout, Characters))
	e.GET("/characters/body", middleware.StaticHTMX(portfolioTemplates.Characters()))
}
