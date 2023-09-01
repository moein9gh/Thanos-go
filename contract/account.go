package contract

import (
	"context"

	"github.com/thanos-go/param"

	"github.com/p3ym4n/re"
)

type Account interface {
	AddAccount(c context.Context, req *param.AddAccountRequest) (*param.AddAccountResponse, re.Error)
}
