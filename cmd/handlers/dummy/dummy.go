package dummy

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	dummyTemplates "mael/web/templates/contents/dummy"
)

func Dummy(c echo.Context) templ.Component {
	return dummyTemplates.Dummy()
}
