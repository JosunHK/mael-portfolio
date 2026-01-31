package cms

import (
	"fmt"
	"mael/cmd/struct/error"
	responseUtil "mael/cmd/util/response"
	"mael/web/templates/contents/errorAlert"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var animationActions = map[string]AnimationPatch{
	"orderUp":      OrderUp,
	"orderDown":    OrderDown,
	"modifyDetail": ModifyDetail,
}

var animationActionsResBody = map[string]AnimationPatchResBody{
	"orderUp":      GetAnimtions,
	"orderDown":    GetAnimtions,
	"modifyDetail": GetAnimtionDetail,
}

var animationActionsResFunc = map[string]AnimationPatchResFunc{
	"orderUp":      responseUtil.HTMX,
	"orderDown":    responseUtil.HTMX,
	"modifyDetail": responseUtil.HTMXWithSuccess,
}

func GetAnimationRes(c echo.Context) error {
	table, err := GetAnimtions(c)
	return responseUtil.HTMX(c, table, err)
}

func AddAnimationRes(c echo.Context) error {
	resErr := AddAnimation(c)
	table, err := GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func DeleteAnimationRes(c echo.Context) error {
	resErr := DeleteAnimation(c)
	table, err := GetAnimtions(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func PatchAnimation(c echo.Context) error {
	action := c.FormValue("action")
	actionFunc := animationActions[action]

	//error msg in case no action
	resErr := resError.New(fmt.Sprintf("Invalid Action : %v", action), "")
	if actionFunc != nil {
		resErr = actionFunc(c)
	}

	actionResBody := animationActionsResBody[action]
	if actionResBody == nil { //only happens if YOU fucked up
		return responseUtil.HTML(c, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid."))
	}

	resBody, err := actionResBody(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}

	resFunc := animationActionsResFunc[action]
	if resFunc == nil { //only happens if YOU fucked up
		return responseUtil.HTML(c, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid."))
	}

	return resFunc(c, resBody, resErr)
}

func GetAnimationDetail(c echo.Context) templ.Component {
	detail, err := GetAnimtionDetail(c)
	if err != nil {
		return errorTemplate.ErrorAlert(err.Title, err.Desc)
	}
	return detail
}
