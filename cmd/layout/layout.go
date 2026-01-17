package layout

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	responseUtil "mael/cmd/util/response"
	layoutTemplates "mael/web/templates/layout"
)

func Layout(c echo.Context, content templ.Component) error {
	return responseUtil.HTML(c, layoutTemplates.Layout(content))
}

func ErrorPage(c echo.Context, content templ.Component) error {
	return responseUtil.HTML(c, layoutTemplates.ErrorPage(content))
}

func Component(c echo.Context, content templ.Component) error {
	return responseUtil.HTML(c, content)
}
