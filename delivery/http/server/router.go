package server

import (
	apmecho "go.elastic.co/apm/module/apmechov4/v2"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thanos-go/config"
	"github.com/thanos-go/log"

	echomw "github.com/labstack/echo/v4/middleware"
	mw "github.com/thanos-go/delivery/http/middleware"
)

func NewRouter(cfg *config.Config) *echo.Echo {

	router := echo.New()
	router.Use(apmecho.Middleware())
	router.Debug = cfg.IsDebugging()
	router.HidePort = true
	router.HideBanner = true
	router.StdLogger = log.Std()

	// pre-routing middleware
	router.Pre(
		echomw.RemoveTrailingSlash(),
	)

	router.Static("/static", cfg.Static.StaticFilePath)

	router.Use(
		mw.RequestsLogger(),
		echomw.GzipWithConfig(echomw.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Path(), "metrics") || strings.Contains(c.Path(), "subscribe") // Change "metrics" for your own path
			},
		}),
		echomw.RecoverWithConfig(echomw.RecoverConfig{
			DisableStackAll: true,
		}),
		echomw.CORSWithConfig(echomw.CORSConfig{
			Skipper: func(echo.Context) bool {
				return !cfg.App.CorsEnabled
			},
			MaxAge: 31200,
		}),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Skipper: func(c echo.Context) bool {
				return c.IsWebSocket()
			},
			Timeout: cfg.App.RequestTimeout,
		}),
	)

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "object not found"})
	}

	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, echo.Map{"message": "method not allowed"})
	}

	return router
}
