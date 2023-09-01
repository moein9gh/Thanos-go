package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// HealthCheck godoc
// @Summary      Health check
// @Description  Health check
// @Tags         Health check
// @Success      204
// @Failure      400  {object}  param.BadRequestHttpError  "bad request"
// @Router       /health-check [get]
func (h *Handlers) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
