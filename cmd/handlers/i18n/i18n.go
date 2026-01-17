package i18n

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	i18nUtil "mael/cmd/util/i18n"
	i18nTemplates "mael/web/templates/contents/i18n"
)

func Table(c echo.Context) templ.Component {
	locale := c.Param("locale")
	return i18nTemplates.I18n(locale)
}

func SetLocale(c echo.Context) error {
	locale := c.Param("locale")
	cookie := new(http.Cookie)
	cookie.Name = i18nUtil.LOCALE_SETTING_ID
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Path = "/"
	cookie.Secure = true
	cookie.Value = locale
	cookie.Expires = time.Now().Add(240 * time.Hour) //10 days
	c.SetCookie(cookie)

	return nil
}
