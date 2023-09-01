package claims

import (
	"github.com/golang-jwt/jwt/v4"
)

const ClaimContextKey = "account"

type Claims struct {
	ID         uint   `json:"id"`
	AppVersion string `json:"app_version"`
	Email      string `json:"email"`
	jwt.StandardClaims
}
