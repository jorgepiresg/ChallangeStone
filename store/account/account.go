package account

import (
	"context"

	"github.com/jmoiron/sqlx"
	model_account "github.com/jorgepiresg/ChallangeStone/model/account"
)

//go:generate mockgen -source=$GOFILE -destination=../../mocks/account_mock.go -package=mocks
type IAccount interface {
	Get(ctx context.Context) ([]model_account.Account, error)
	GetByID(ctx context.Context, ID string) (model_account.Account, error)
	Create(ctx context.Context, account model_account.Create) (model_account.Account, error)
	GetByCPF(ctx context.Context, cpf string) (model_account.Account, error)
	UpdateBalance(ctx context.Context, ID string, amount int) error
}

type account struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) IAccount {
	return account{
		db: db,
	}
}

func (a account) Get(ctx context.Context) ([]model_account.Account, error) {
	accounts := make([]model_account.Account, 0, 0)

	err := a.db.SelectContext(ctx, &accounts, `SELECT id, name, cpf, secret, balance, created_at FROM account`)
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (a account) GetByID(ctx context.Context, ID string) (model_account.Account, error) {

	var account model_account.Account

	err := a.db.GetContext(ctx, &account, `SELECT id, name, cpf, secret, balance, created_at FROM account where id = $1`, ID)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (a account) Create(ctx context.Context, create model_account.Create) (model_account.Account, error) {

	var account model_account.Account

	rows, err := a.db.NamedQueryContext(ctx, `INSERT INTO account (name, cpf, secret, balance) VALUES (:name, :cpf, :secret, 0) RETURNING *`, create)
	if err != nil {
		return account, err
	}

	for rows.Next() {
		err = rows.StructScan(&account)
		if err != nil {
			return account, err
		}
	}

	return account, nil
}

func (a account) GetByCPF(ctx context.Context, cpf string) (model_account.Account, error) {
	var account model_account.Account

	err := a.db.GetContext(ctx, &account, `SELECT id, name, cpf, secret, balance, created_at FROM account where cpf = $1`, cpf)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (a account) UpdateBalance(ctx context.Context, ID string, amount int) error {

	_, err := a.db.ExecContext(ctx, `UPDATE account SET balance = $1 where id = $2;`, amount, ID)
	if err != nil {
		return err
	}

	return nil
}
