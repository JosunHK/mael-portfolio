package portfolio

import (
	"mael/cmd/database"
	resError "mael/cmd/struct/error"
	responseUtil "mael/cmd/util/response"
	sqlc "mael/db/generated"
	errorTemplate "mael/web/templates/contents/errorAlert"
	portfolioTemplates "mael/web/templates/contents/portfolio"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)



func AnimationBody(c echo.Context) error {
	res, err := GetAnimations(c)
	
	return responseUtil.HTMX(c, res, err)
}

func Characters(c echo.Context) templ.Component {
	return portfolioTemplates.Characters()
}

func GetAnimations(c echo.Context) (templ.Component, *resError.Error) {
	queries := sqlc.New(database.DB)
	deskId, errD := queries.GetThumbDesktop(c.Request().Context())
	mobilId, errM := queries.GetThumbMobile(c.Request().Context())
	
	if errD != nil || errM != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{}), resError.New("Failed to retrive Data ", errD.Error()+" "+errM.Error())
	}
	resD, err := queries.GetAnimationById(c.Request().Context(), deskId.Int64)
	if err != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{}), resError.New("Failed to retrive Data ", err.Error())
	}
	resM, err := queries.GetAnimationById(c.Request().Context(), mobilId.Int64)
	if err != nil {
		return portfolioTemplates.Animations(sqlc.Animation{}, sqlc.Animation{}), resError.New("Failed to retrive Data ", err.Error())
	}
	return portfolioTemplates.Animations(resD, resM), nil
}

func Animations(c echo.Context) templ.Component {
	res, err := GetAnimations(c)
	if err != nil{
		return errorTemplate.ErrorAlert(err.Title, err.Desc)
	}
	return res
}