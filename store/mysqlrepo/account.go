package mysqlrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/p3ym4n/re"

	"github.com/thanos-go/model"
)

func (mr *MysqlRepo) GetAccountByID(c context.Context, id uint) (*model.Account, re.Error) {
	const op = re.Op("store.GetAccountByID")

	stmt := mr.stmt("stmtSelectAccountByID")
	if stmt == nil {
		ps, err := mr.conn.PrepareContext(c, `select accounts.* from accounts WHERE accounts.id = ?`)
		if err != nil {
			return nil, re.New(op, fmt.Errorf("prepare context error : %w", err), re.KindUnexpected)
		}
		mr.setStmt("stmtSelectAccountByID", ps)
		stmt = ps
	}

	account := new(model.Account)
	if err2 := scanAccount(stmt.QueryRowContext(c, id), account); err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil, re.New(op, fmt.Errorf("account not found"), re.KindNotFound)
		}
		return nil, re.New(op, fmt.Errorf("query error : %w", err2), re.KindUnexpected)
	}
	return account, nil
}

func (mr *MysqlRepo) CreateAccount(c context.Context, accountEntity *model.Account) (*model.Account, re.Error) {
	const op = re.Op("store.CreateAccount")

	stmt := mr.stmt("stmtCreateAccount")
	if stmt == nil {
		ps, err := mr.conn.PrepareContext(c, `insert into accounts 
    	(app_version, email, is_deleted) VALUES(?, ?, ?)`)
		if err != nil {
			return nil, re.New(op, fmt.Errorf("prepare context error : %w", err), re.KindUnexpected)
		}
		mr.setStmt("stmtCreateAccount", ps)
		stmt = ps
	}

	res, err2 := stmt.ExecContext(c,
		accountEntity.AppVersion,
		accountEntity.Email,
		accountEntity.IsDeleted,
	)
	if err2 != nil {
		return nil, re.New(op, fmt.Errorf("execute error : %w", err2), re.KindUnexpected)
	}

	id, _ := res.LastInsertId()

	accountEntity.ID = uint(id)
	return accountEntity, nil
}

func (mr *MysqlRepo) EmailUniquenessCheck(ctx context.Context, email string) (*model.Account, re.Error) {
	const op = re.Op("store.SessionAndDeviceIdUniquenessCheck")

	stmt := mr.stmt("stmtSessionAndDeviceIdUniquenessCheck")
	if stmt == nil {
		ps, err := mr.conn.PrepareContext(ctx, `select accounts.* from accounts WHERE accounts.email = ? AND accounts.is_deleted = 0`)
		if err != nil {
			return nil, re.New(op, fmt.Errorf("prepare context error: %w", err), re.KindUnexpected)
		}
		mr.setStmt("stmtSessionAndDeviceIdUniquenessCheck", ps)
		stmt = ps
	}

	account := new(model.Account)
	if err2 := scanAccount(stmt.QueryRowContext(ctx, email), account); err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil, nil
		}
		return nil, re.New(op, fmt.Errorf("query error %w", err2), re.KindUnexpected)
	}

	return account, nil
}

func (mr *MysqlRepo) SoftDeleteAccount(c context.Context, account *model.Account) re.Error {
	const op = re.Op("store.SoftDeleteAccount")
	stmt := mr.stmt("stmtSoftDeleteAccount")
	if stmt == nil {
		ps, err := mr.conn.PrepareContext(c, `
			update accounts set device_id = ?, session_id = ?, platform = ?, app_version = ?, is_deleted = ? where id = ?`)
		if err != nil {
			return re.New(op, fmt.Errorf("prepare context error : %w", err), re.KindUnexpected)
		}
		mr.setStmt("stmtSoftDeleteAccount", ps)
		stmt = ps
	}

	_, err := stmt.ExecContext(c,
		account.AppVersion,
		account.IsDeleted,
		account.ID,
	)
	if err != nil {
		return re.New(op, fmt.Errorf("execute error : %w", err), re.KindUnexpected)
	}
	return nil
}
