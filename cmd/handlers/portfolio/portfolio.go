package portfolio

import (
	"mael/cmd/database"
	sqlc "mael/db/generated"
	portfolioTemplates "mael/web/templates/contents/portfolio"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)


func Characters(c echo.Context) templ.Component {
	return portfolioTemplates.Characters()
}

func Animations(c echo.Context) templ.Component {
	queries := sqlc.New(database.DB)
	deskId, errD := queries.GetThumbDesktop(c.Request().Context())
	mobilId, errM := queries.GetThumbMobile(c.Request().Context())
	
	if errD != nil || errM != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{})
	}
  
  res, err := queries.GetUploadedAnimations(c.Request().Context())
	if err != nil {
		return portfolioTemplates.Animations([]sqlc.Animation{})
	}
  
	resD, err := queries.GetAnimationById(c.Request().Context(), deskId.Int64)
	if err != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{})
	}
  
	resM, err := queries.GetAnimationById(c.Request().Context(), mobilId.Int64)
	if err != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{})
	}
  
	return portfolioTemplates.Animations(res, resD, resM)
}

