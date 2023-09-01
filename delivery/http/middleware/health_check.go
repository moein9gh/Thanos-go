package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanos-go/log"

	"github.com/thanos-go/store/mysqlrepo"
)

func HealthCheck(mysqlRepo *mysqlrepo.MysqlRepo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			isDown := false
			if err := mysqlRepo.Ping(); err != nil {
				log.Named("http health check middleware").Error("can't ping mysql %v", err)
				isDown = true
			}
			if isDown {
				return c.NoContent(http.StatusServiceUnavailable)
			}
			return next(c)
		}
	}
}
