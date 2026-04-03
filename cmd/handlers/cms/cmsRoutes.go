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
	e.GET("/cms/animation/sub/table/:mainId", GetSubAnimationRes)
	e.GET("/cms/animation/sub/:mainId", middleware.Pages(layout.CMSLayout, GetSubAnimationWrapper))
	e.POST("/cms/animation", AddAnimationRes)
	e.POST("/cms/animation/sub/:mainId", AddSubAnimationRes)
	e.DELETE("/cms/animation/:id", DeleteAnimationRes)
	e.DELETE("/cms/animation/sub/:id", DeleteSubAnimationRes)
	e.PATCH("/cms/animation/:id", PatchAnimation)
	e.PATCH("/cms/animation/sub/:id", PatchSubAnimation)

	//Detail
	e.GET("/cms/animation/:id", middleware.Pages(layout.CMSLayout, GetAnimationDetail))
	e.GET("/cms/animation/sub/detail/:id", middleware.Pages(layout.CMSLayout, GetSubAnimationDetail))
}
