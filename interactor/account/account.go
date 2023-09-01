package account

import (
	"context"
	"fmt"

	"github.com/p3ym4n/re"
	"github.com/thanos-go/model"
	"github.com/thanos-go/param"
)

func (s *Service) AddAccount(c context.Context, req *param.AddAccountRequest) (*param.AddAccountResponse, re.Error) {
	const op = re.Op("service.AddAccount")
	meta := re.Meta{"request": *req}

	err := s.validation.AddAccountValidation(req)
	if err != nil {
		return nil, err.ChainWithMeta(op, meta)
	}
	account, err := s.mysqlStore.EmailUniquenessCheck(c, req.Email)
	if err != nil {
		return nil, re.New(op, fmt.Errorf("uniqness check error: %s", err), meta)
	}

	if account == nil {
		account = &model.Account{
			AppVersion: req.AppVersion,
			Email:      req.Email,
		}

		_, err2 := s.mysqlStore.CreateAccount(c, account)
		if err2 != nil {
			return nil, re.New(op, fmt.Errorf("add account error : %s", err2), meta)
		}
	}

	accessToken, err3 := s.authentication.GenerateAccessToken(*account)
	if err3 != nil {
		return nil, re.New(op, fmt.Errorf("add account generate access token error : %s", err3), meta)
	}

	refreshToken, err4 := s.authentication.GenerateRefreshToken(*account)
	if err4 != nil {
		return nil, re.New(op, fmt.Errorf("add account generate refresh token error : %s", err4), meta)
	}

	tokens := param.AccountTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	res := &param.AddAccountResponse{
		AccountID: account.ID,
		Tokens:    tokens,
	}

	return res, nil
}
