package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thanos-go/delivery/http/handler"
	"github.com/thanos-go/pkg/claims"
)

func RegisterRoutes(router *echo.Echo, handler *handler.Handlers) {

	jwtConfig := middleware.JWTConfig{
		Claims:     &claims.Claims{},
		SigningKey: []byte(handler.Config.Authentication.JwtSecret),
		ContextKey: claims.ClaimContextKey,
	}

	v1 := router.Group("/v1")

	v1.POST("/accounts", handler.AddAccount)

	v1.GET("/health-check", handler.HealthCheck, middleware.JWTWithConfig(jwtConfig))
}
