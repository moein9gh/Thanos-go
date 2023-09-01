package authentication

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/p3ym4n/re"
	"github.com/thanos-go/config"
	"github.com/thanos-go/model"

	"github.com/thanos-go/pkg/claims"
)

type Authenticate struct {
	config *config.Authentication
}

func New(config *config.Authentication) *Authenticate {
	return &Authenticate{config: config}
}

func (a Authenticate) GenerateAccessToken(account model.Account) (string, re.Error) {
	const op = re.Op("authenticate.GenerateAccessToken")

	accessExpirationTime := time.Now().Add(time.Duration(a.config.AccessExpirationInMinute) * time.Minute)

	clm := &claims.Claims{
		ID:         account.ID,
		AppVersion: account.AppVersion,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)
	tokenString, err := accessToken.SignedString([]byte(a.config.JwtSecret))
	if err != nil {
		return "", re.New(op, fmt.Errorf("access token signed string error"), re.KindUnexpected)
	}

	return tokenString, nil
}

func (a Authenticate) GenerateRefreshToken(account model.Account) (string, re.Error) {
	const op = re.Op("authenticate.GenerateRefreshToken")

	refreshExpirationTime := time.Now().Add(time.Duration(a.config.RefreshExpirationInMinute) * time.Minute)

	clm := &claims.Claims{
		ID:         account.ID,
		AppVersion: account.AppVersion,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)
	tokenString, err := refreshToken.SignedString([]byte(a.config.JwtSecret))
	if err != nil {
		return "", re.New(op, fmt.Errorf("refresh token signed string error"), re.KindUnexpected)
	}

	return tokenString, nil
}

func (a Authenticate) ParseToken(tokenString string) (*claims.Claims, re.Error) {
	const op = re.Op("authenticate.ParseToken")

	tk, _ := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JwtSecret), nil
	})

	if tk == nil {
		return nil, re.New(op, fmt.Errorf("parse token is nil"), re.KindInvalid)
	}

	if clm, ok := tk.Claims.(*claims.Claims); ok && tk.Valid {
		return clm, nil
	} else {
		return nil, re.New(op, fmt.Errorf("token claim error"), re.KindUnexpected)
	}
}
