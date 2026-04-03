package cms

import (
	"fmt"
	resError "mael/cmd/struct/error"
	responseUtil "mael/cmd/util/response"
	cmsTemplates "mael/web/templates/contents/cms"
	errorTemplate "mael/web/templates/contents/errorAlert"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var animationActions = map[string]AnimationPatch{
	"orderUp":                  OrderUp,
	"orderDown":                OrderDown,
	"modifyDetail":             ModifyDetail,
	"modifyThumbMobileDetail":  ModifyThumbMobile,
	"modifyThumbDesktopDetail": ModifyThumbDesktop,
	"modifyThumbMobileTable":   ModifyThumbMobile,
	"modifyThumbDesktopTable":  ModifyThumbDesktop,
}

var animationActionsResBody = map[string]AnimationPatchResBody{
	"orderUp":                  GetAnimations,
	"orderDown":                GetAnimations,
	"modifyDetail":             GetAnimtionDetail,
	"modifyThumbMobileDetail":  GetAnimtionDetail,
	"modifyThumbDesktopDetail": GetAnimtionDetail,
	"modifyThumbMobileTable":   GetAnimations,
	"modifyThumbDesktopTable":  GetAnimations,
}

var animationActionsResFunc = map[string]AnimationPatchResFunc{
	"orderUp":                  responseUtil.HTMX,
	"orderDown":                responseUtil.HTMX,
	"modifyDetail":             responseUtil.HTMXWithSuccess,
	"modifyThumbMobileDetail":  responseUtil.HTMX,
	"modifyThumbDesktopDetail": responseUtil.HTMX,
	"modifyThumbMobileTable":   responseUtil.HTMX,
	"modifyThumbDesktopTable":  responseUtil.HTMX,
}

var subAnimationActions = map[string]AnimationPatch{
	"orderUp":      SubOrderUp,
	"orderDown":    SubOrderDown,
	"modifyDetail": SubModifyDetail,
}

var subAnimationActionsResBody = map[string]AnimationPatchResBody{
	"orderUp":      GetSubAnimations,
	"orderDown":    GetSubAnimations,
	"modifyDetail": GetSubAnimtionDetail,
}

var subAnimationActionsResFunc = map[string]AnimationPatchResFunc{
	"orderUp":      responseUtil.HTMX,
	"orderDown":    responseUtil.HTMX,
	"modifyDetail": responseUtil.HTMXWithSuccess,
}

func GetAnimationRes(c echo.Context) error {
	table, err := GetAnimations(c)
	return responseUtil.HTMX(c, table, err)
}

func GetSubAnimationRes(c echo.Context) error {
	table, err := GetSubAnimations(c)
	return responseUtil.HTMX(c, table, err)
}

func GetSubAnimationWrapper(c echo.Context) templ.Component {
	id, err := strconv.ParseInt(c.Param("mainId"), 10, 64)
	if err != nil {
		return errorTemplate.SimpleError("invalid Id")
	}

	return cmsTemplates.SubAnimations(id)
}

func AddAnimationRes(c echo.Context) error {
	resErr := AddAnimation(c)
	table, err := GetAnimations(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func AddSubAnimationRes(c echo.Context) error {
	resErr := AddSubAnimation(c)
	table, err := GetSubAnimations(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}
	return responseUtil.HTMX(c, table, resErr)
}

func DeleteAnimationRes(c echo.Context) error {
	resErr := DeleteAnimation(c)
	table, err := GetAnimations(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}

	return responseUtil.HTMX(c, table, resErr)
}

func DeleteSubAnimationRes(c echo.Context) error {
	resErr := DeleteSubAnimation(c)
	table, err := GetSubAnimations(c)
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

func PatchSubAnimation(c echo.Context) error {
	action := c.FormValue("action")
	actionFunc := subAnimationActions[action]

	//error msg in case no action
	resErr := resError.New(fmt.Sprintf("Invalid Action : %v", action), "")
	if actionFunc != nil {
		resErr = actionFunc(c)
	}

	actionResBody := subAnimationActionsResBody[action]
	if actionResBody == nil { //only happens if YOU fucked up
		return responseUtil.HTML(c, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid."))
	}

	resBody, err := actionResBody(c)
	if err != nil && resErr == nil { //we pioritize the error of action
		resErr = err
	}

	resFunc := subAnimationActionsResFunc[action]
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

func GetSubAnimationDetail(c echo.Context) templ.Component {
	detail, err := GetSubAnimtionDetail(c)
	if err != nil {
		return errorTemplate.ErrorAlert(err.Title, err.Desc)
	}
	return detail
}
