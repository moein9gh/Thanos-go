package mysqlrepo

import (
	"github.com/thanos-go/model"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

func scanAccount(row Scanner, u *model.Account) error {
	return row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.AppVersion,
		&u.Email,
		&u.IsDeleted,
	)
}
