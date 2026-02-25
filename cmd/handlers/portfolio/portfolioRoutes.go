package portfolio

import (
	"mael/cmd/layout"
	"mael/cmd/middleware"
	portfolioTemplates "mael/web/templates/contents/portfolio"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", middleware.Pages(layout.Layout, Animations))
	e.GET("/animation", middleware.Pages(layout.Layout, Animations))
	e.GET("/animation/body", AnimationBody)
	e.GET("/characters", middleware.Pages(layout.Layout, Characters))
	e.GET("/characters/body", middleware.HTMX(portfolioTemplates.Characters()))

}
