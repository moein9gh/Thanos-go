package account

import (
	"github.com/thanos-go/config"
	"github.com/thanos-go/contract"
)

type Service struct {
	config         config.Config
	mysqlStore     contract.Store
	authentication contract.IAuthInteractor
	validation     contract.Validation
}

func New(config config.Config, mysqlStore contract.Store, authentication contract.IAuthInteractor,
	validation contract.Validation) Service {
	return Service{
		config:         config,
		mysqlStore:     mysqlStore,
		authentication: authentication,
		validation:     validation,
	}
}
