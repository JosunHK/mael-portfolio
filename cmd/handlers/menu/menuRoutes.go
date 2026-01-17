package menu

import (
	"github.com/labstack/echo/v4"
	"mael/cmd/middleware"
	"mael/cmd/util/menuProvider"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/purge", middleware.Nothing(menuProvider.PurgeCache))
}
