package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p3ym4n/re"
	"github.com/thanos-go/e"
	"github.com/thanos-go/param"
)

// AddAccount godoc
//
//	@Summary		Create Account
//	@Description	Create Account
//	@Tags			Account
//	@Param			request	body		param.AddAccountRequest	true	"request body"
//	@Success		200		{array}		param.AddAccountResponse
//	@Failure		400		{object}	param.BadRequestHttpError	"bad request"
//	@Failure		401		{object}	param.UnAuthorizedHttpError	"unauthorized"
//	@Failure		422		{object}	param.UnprocessableEntity	"Unprocessable Entity"
//	@Failure		500		{object}	param.InternalError			"internal server error"
//	@Router			/accounts [post]
func (h *Handlers) AddAccount(c echo.Context) error {
	const op = re.Op("handler.AddAccount")

	var req param.AddAccountRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.AccountService.AddAccount(c.Request().Context(), &req)
	if err != nil {
		code, out := e.LogAndDeliver(err.Chain(op))
		return c.JSON(code, out)
	}
	return c.JSON(http.StatusOK, token)

}
