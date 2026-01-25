package cms

import (
	"fmt"
	"mael/cmd/struct/error"
	"mael/cmd/util/cms"
	responseUtil "mael/cmd/util/response"
	"mael/web/templates/contents/errorAlert"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var animattionActions = map[string]cmsUtil.AnimationPatch{
	"orderUp":      cmsUtil.OrderUp,
	"orderDown":    cmsUtil.OrderDown,
	"modifyDetail": cmsUtil.ModifyDetail,
}

var animattionActionsResBody = map[string]cmsUtil.AnimationPatchResBody{
	"orderUp":      cmsUtil.GetAnimtions,
	"orderDown":    cmsUtil.GetAnimtions,
	"modifyDetail": cmsUtil.GetAnimtionDetail,
}

var animattionActionsResFunc = map[string]cmsUtil.AnimationPatchResFunc{
	"orderUp":      responseUtil.HTMX,
	"orderDown":    responseUtil.HTMX,
	"modifyDetail": responseUtil.HTMXWithSuccess,
}

func GetAnimationRes(c echo.Context) error {
	table, err := cmsUtil.GetAnimtions(c)
	return responseUtil.HTMX(c, table, err)
}

func AddAnimationRes(c echo.Context) error {
	resErr := cmsUtil.AddAnimation(c)
	table, err := cmsUtil.GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func DeleteAnimationRes(c echo.Context) error {
	resErr := cmsUtil.DeleteAnimation(c)
	table, err := cmsUtil.GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func PatchAnimation(c echo.Context) error {
	action := c.FormValue("action")
	actionFunc := animattionActions[action]

	//error msg in case no action
	resErr := resError.New(fmt.Sprintf("Invalid Action : %v", action), "")
	if actionFunc != nil {
		resErr = actionFunc(c)
	}

	actionResBody := animattionActionsResBody[action]
	if actionResBody == nil { //only happens if YOU fucked up
		return responseUtil.HTML(c, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid."))
	}
	resBody, err := actionResBody(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}

	resFunc := animattionActionsResFunc[action]
	if resFunc == nil { //only happens if YOU fucked up
		return responseUtil.HTML(c, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid."))
	}

	return resFunc(c, resBody, resErr)
}

func GetAnimationDetail(c echo.Context) templ.Component {
	detail, err := cmsUtil.GetAnimtionDetail(c)
	if err != nil {
		return errorTemplate.ErrorAlert(err.Title, err.Desc)
	}
	return detail
}
