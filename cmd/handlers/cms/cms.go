package cms

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mael/cmd/layout"
	"mael/cmd/middleware"
	"mael/web/templates/contents/cms"
)

const CMS_ROUTE_PREFIX = "cms"

func RegisterRoutes(e *echo.Echo) {
	e.GET(fmt.Sprintf("/%s", CMS_ROUTE_PREFIX), middleware.StaticPages(layout.CMSLayout, cmsTemplates.Categories()))
}
