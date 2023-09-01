package contract

import (
	"context"

	"github.com/thanos-go/model"

	"github.com/p3ym4n/re"
)

type Store interface {
	AccountStore
}

type AccountStore interface {
	GetAccountByID(ctx context.Context, id uint) (*model.Account, re.Error)
	CreateAccount(c context.Context, accountEntity *model.Account) (*model.Account, re.Error)
	EmailUniquenessCheck(ctx context.Context, email string) (*model.Account, re.Error)
}
