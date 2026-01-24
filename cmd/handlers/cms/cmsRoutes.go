package cms

import (
	"github.com/labstack/echo/v4"
	"mael/cmd/layout"
	"mael/cmd/middleware"
	"mael/web/templates/contents/cms"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/cms", middleware.StaticPages(layout.CMSLayout, cmsTemplates.Animations()))
	e.GET("/cms/animation", GetAnimationRes)
	e.POST("/cms/animation", AddAnimationRes)
	e.DELETE("/cms/animation/:id", DeleteAnimationRes)
}
