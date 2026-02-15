package portfolio

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"mael/cmd/database"
	"mael/db/generated"
	"mael/web/templates/contents/portfolio"
)

func Characters(c echo.Context) templ.Component {
	return portfolioTemplates.Characters()
}

func Animations(c echo.Context) templ.Component {
	queries := sqlc.New(database.DB)
	res, err := queries.GetUploadedAnimations(c.Request().Context())
	if err != nil {
		return portfolioTemplates.Animations([]sqlc.Animation{})
	}

	return portfolioTemplates.Animations(res)
}
