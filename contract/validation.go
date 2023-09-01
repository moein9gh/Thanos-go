package contract

import (
	"github.com/p3ym4n/re"
	"github.com/thanos-go/param"
)

type Validation interface {
	AccountValidation
}

type AccountValidation interface {
	AddAccountValidation(req *param.AddAccountRequest) re.Error
}
