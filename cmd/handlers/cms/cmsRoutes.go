package cms

import (
	"mael/cmd/layout"
	"mael/cmd/middleware"
	cmsTemplates "mael/web/templates/contents/cms"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	RegisterAnimationRoutes(e)
}

func RegisterAnimationRoutes(e *echo.Echo) {
	//profile
	e.GET("/cms/animation", middleware.StaticPages(layout.CMSLayout, cmsTemplates.Animations()))
	e.GET("/cms/animation/table", GetAnimationRes)
	e.POST("/cms/animation", AddAnimationRes)
	e.DELETE("/cms/animation/:id", DeleteAnimationRes)
	e.PATCH("/cms/animation/:id", PatchAnimation)

	//Detail
	e.GET("/cms/animation/:id", middleware.Pages(layout.CMSLayout, GetAnimationDetail))
}
