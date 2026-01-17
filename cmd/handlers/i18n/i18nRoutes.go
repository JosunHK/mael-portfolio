package i18n

import (
	"mael/cmd/layout"
	"mael/cmd/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//main endpoints
	e.POST("/i18n/set/:locale", middleware.NoContent(SetLocale))
	e.GET("/i18n/:locale", middleware.Pages(layout.Layout, Table))

	//api endpoints
	e.GET("/i18n/items/:locale", middleware.HTML(GetItems))
	e.POST("/i18n/items/:locale", middleware.HTML(AddItems))

	//I'll add this when I want
	//e.DELETE("/i18n/items/:locale", middleware.HTML(i18nAPI.DeleteItems))
}
