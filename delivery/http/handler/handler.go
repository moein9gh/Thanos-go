package handler

import (
	"github.com/thanos-go/config"
	"github.com/thanos-go/contract"
)

type Handlers struct {
	AccountService        contract.Account
	AuthenticationService contract.IAuthInteractor
	Config                config.Config
	Repo                  contract.Store
}

func NewHandlers(accountService contract.Account, authenticationService contract.IAuthInteractor, cfg config.Config, repo contract.Store) *Handlers {

	return &Handlers{
		AuthenticationService: authenticationService,
		Config:                cfg,
		Repo:                  repo,
		AccountService:        accountService,
	}
}
