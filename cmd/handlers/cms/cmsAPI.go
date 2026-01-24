package cms

import (
	"github.com/labstack/echo/v4"
	"mael/cmd/util/cms"
	responseUtil "mael/cmd/util/response"
)

func GetAnimationRes(c echo.Context) error {
	table, err := cmsUtil.GetAnimtions(c)
	return responseUtil.HTMX(c, table, err)
}

func AddAnimationRes(c echo.Context) error {
	resErr := cmsUtil.AddAnimation(c)
	table, err := cmsUtil.GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of adding
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func DeleteAnimationRes(c echo.Context) error {
	resErr := cmsUtil.DeleteAnimation(c)
	table, err := cmsUtil.GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of adding
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}
