package portfolio

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"mael/web/templates/contents/portfolio"
)

func Animations(c echo.Context) templ.Component {
	return portfolioTemplates.Animations()
}

func Characters(c echo.Context) templ.Component {
	return portfolioTemplates.Characters()
}
