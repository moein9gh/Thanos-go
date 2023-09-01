package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/thanos-go/log"

	jsoniter "github.com/json-iterator/go"
	"github.com/thanos-go/pkg/baser"
	"github.com/thanos-go/pkg/claims"
)

func BearerAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Param("token")
			claim, err := parseJWT(authHeader)
			if err != nil {
				log.Debug("jwt err:", err)
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "the authorization token is not valid",
				})
			}

			id := claim.ID

			accountID, err := strconv.Atoi(fmt.Sprintf("%v", id))
			if err != nil {
				log.Debug("strconv Atoi err:", err)
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "the authorization token is not valid",
				})
			}

			c.Set("account_id", uint(accountID))

			return next(c)
		}
	}
}

func parseJWT(rawHeader string) (claims.Claims, error) {

	//claim := make(map[string]interface{})
	var claim claims.Claims
	parts := strings.Split(strings.Replace(rawHeader, "Bearer ", "", 1), ".")
	if len(parts) != 3 {
		return claim, errors.New("given token is not a valid JWT")
	}

	decoded, err := baser.DecodeString(parts[1])
	if err != nil {
		return claim, fmt.Errorf("base64 decode: %w", err)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(decoded, &claim); err != nil {
		return claim, fmt.Errorf("json decode: %w", err)
	}

	return claim, nil
}
