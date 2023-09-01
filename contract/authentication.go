package contract

import (
	"github.com/p3ym4n/re"
	"github.com/thanos-go/model"
	"github.com/thanos-go/pkg/claims"
)

type IAuthInteractor interface {
	GenerateAccessToken(account model.Account) (string, re.Error)
	GenerateRefreshToken(account model.Account) (string, re.Error)
	ParseToken(token string) (*claims.Claims, re.Error)
}
