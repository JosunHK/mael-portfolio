package responseUtil

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"mael/cmd/struct/error"
	i18nUtil "mael/cmd/util/i18n"
	"mael/web/templates/contents/errorAlert"
	"mael/web/templates/contents/successAlert"
)

func HTML(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func HTMX(c echo.Context, cmp templ.Component, err *resError.Error) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if err != nil {
		cmp = errorTemplate.ErrorToastWrap(*err, cmp)
	}
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func HTMXWithSuccess(c echo.Context, cmp templ.Component, err *resError.Error) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if err != nil {
		cmp = errorTemplate.ErrorToastWrap(*err, cmp)
	} else {
		cmp = successTemplate.SuccessToastWrap("Success!", "Save successful!", cmp)
	}
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func overrideAndGetContext(c echo.Context) context.Context {
	context := c.Request().Context()
	context = overrideContextWithLocale(c)
	return context
}

func overrideContextWithLocale(c echo.Context) context.Context {
	var locale string
	cookie, err := c.Cookie(i18nUtil.LOCALE_SETTING_ID)
	if err != nil {
		locale = "en"
	} else {
		locale = cookie.Value
	}

	return context.WithValue(c.Request().Context(), i18nUtil.LOCALE_SETTING_ID, locale)
}
