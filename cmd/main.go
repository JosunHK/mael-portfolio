package main

import (
	"fmt"
	"net/http"
	"os"

	"mael/cmd/database"
	"mael/cmd/handlers/cms"
	"mael/cmd/handlers/dummy"
	"mael/cmd/handlers/i18n"
	"mael/cmd/handlers/menu"
	"mael/cmd/layout"
	"mael/cmd/middleware"
	i18nUtil "mael/cmd/util/i18n"
	errorTemplate "mael/web/templates/contents/errorAlert"
	playgroundTemplates "mael/web/templates/contents/playground"

	twmerge "github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
	eMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const MYSQL_PARAMS = "?parseTime=true&loc=Local"

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.SetOutput(file)
		log.SetLevel(log.DebugLevel)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	//merger for tailwind
	config := twmerge.MakeDefaultConfig()
	_ = twmerge.CreateTwMerge(config, nil) // config, cache (if nil default will be used)
}

func main() {
	PORT := os.Getenv("PORT")
	PORT = ":" + PORT
	if err := database.InitDB(os.Getenv("DB_CREDENTIALS") + MYSQL_PARAMS); err != nil {
		log.Error(err)
		return
	}

	defer database.DB.Close()

	if err := i18nUtil.InitI18n(); err != nil {
		log.Error(err)
		return
	}

	e := echo.New() //http client
	e.Use(eMiddleware.Secure())
	e.Use(eMiddleware.Recover())
	e.Use(eMiddleware.RequestLoggerWithConfig(eMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values eMiddleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
	e.Use(middleware.WithLocale)
	e.Use(middleware.WithCSP)
	e.Use(eMiddleware.Gzip())
	e.Pre(eMiddleware.RemoveTrailingSlashWithConfig(
		eMiddleware.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		},
	))

	//static files

	if os.Getenv("env") == "production" {
		e.Static("/assets", "railwayAssets")
	} else {
		e.Static("/assets", "assets")
	}
	e.Static("/static", "web/static")
	//e.File("/favicon.ico", "web/static/favicon.ico")

	e.GET("/playground", middleware.StaticPages(layout.Layout, playgroundTemplates.Playground()))

	menu.RegisterRoutes(e)
	i18n.RegisterRoutes(e)
	dummy.RegisterRoutes(e)
	cms.RegisterRoutes(e)

	e.RouteNotFound("", middleware.StaticPages(layout.ErrorPage, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid.")))
	e.RouteNotFound("/*", middleware.StaticPages(layout.ErrorPage, errorTemplate.ErrorAlert("404 Page Not Found", "The path you requested is invalid.")))

	//exit ->
	e.Logger.Fatal(e.Start(PORT))
}
