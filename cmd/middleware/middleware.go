package middleware

import (
	"context"
	"fmt"
	"net/http"

	i18nUtil "mael/cmd/util/i18n"
	responseUtil "mael/cmd/util/response"
	"mael/cmd/util/secure"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type nothingHandler func() error
type requestHandler func(echo.Context) error
type redirectHandler func(echo.Context) (string, error)
type staticPageHandler func(echo.Context, templ.Component) error
type errorHandler func(echo.Context, templ.Component, error) error
type serviceHandler func(echo.Context) (err error, statusCode int, resObj any)
type pageHandler func(echo.Context) templ.Component
type FileHandler func(echo.Context) (err error, fileType string, resObj []byte)

func WithCSP(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nonce := secure.GenerateNonce(16)
		newContext := templ.WithNonce(c.Request().Context(), nonce)
		cspHeader := fmt.Sprintf(`default-src 'self';
            script-src 'nonce-%s';
            style-src 'self' 'nonce-%s' https://fonts.gstatic.com;
			img-src 'self';
            style-src-elem 'nonce-%s';
            font-src 'self' 'nonce-%s' https://fonts.gstatic.com;`, nonce, nonce, nonce, nonce)
		c.Response().Header().Set("Content-Security-Policy", cspHeader)
		c.SetRequest(c.Request().WithContext(newContext))

		if err := next(c); err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
}

// set locale to the context, so that we can retrive it within templ elements
// because in templ we are only exposed to context.Context, but echo uses echo.Context,
// in which, direct call of withValue on echo.Context does not update the underlying echo.Context
func WithLocale(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var locale string
		cookie, err := c.Cookie(i18nUtil.LOCALE_SETTING_ID)
		if err != nil {
			locale = "zh"
		} else {
			locale = cookie.Value
		}
		newContext := context.WithValue(c.Request().Context(), i18nUtil.LOCALE_SETTING_ID, locale)

		c.SetRequest(c.Request().WithContext(newContext))

		if err := next(c); err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Error(err)
		}
		return err
	}
}

func StaticPages(next staticPageHandler, content templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c, content)
	}
}

func Pages(next staticPageHandler, p pageHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c, p(c))
	}
}

func StaticHTMX(component templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		return responseUtil.HTML(c, component)
	}
}

func HTMX(p pageHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return responseUtil.HTML(c, p(c))
	}
}

func HTML(next requestHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func Redirect(next redirectHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		url, err := next(c)
		if err != nil {
			return err
		}

		c.Response().Header().Add("hx-redirect", url)
		return c.NoContent(http.StatusCreated)
	}
}

// this is for a htmx extension that handle error status code,
// thats why we use status 200 dispite it being a nocontent response
func NoContent(next requestHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Error(err)
			return c.String(500, err.Error())
		}

		return c.NoContent(200)
	}
}

func Nothing(next nothingHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next()
		if err != nil {
			log.Error(err)
			return c.String(500, err.Error())
		}

		return c.NoContent(204)
	}
}

func JSON(handler serviceHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, statusCode, resObj := handler(c)
		if err != nil {
			log.Error(err)
			return c.String(statusCode, err.Error())
		}

		return c.JSON(statusCode, resObj)
	}
}

func Image(handler FileHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, fileType, resObj := handler(c)
		if err != nil {
			log.Error(err)
			return c.Blob(http.StatusInternalServerError, fileType, nil)
		}
		return c.Blob(http.StatusOK, fileType, resObj)
	}
}
