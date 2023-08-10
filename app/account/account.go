package account

import (
	"context"
	"fmt"

	model_account "github.com/jorgepiresg/ChallangeStone/model/account"
	"github.com/jorgepiresg/ChallangeStone/store"
	"github.com/jorgepiresg/ChallangeStone/utils"
	"golang.org/x/crypto/bcrypt"
)

type IAccount interface {
	Create(ctx context.Context, account model_account.Create) (model_account.Account, error)
	BalanceByAccountID(ctx context.Context, ID string) (int, error)
	List(ctx context.Context) ([]model_account.Account, error)
	GetByCPF(ctx context.Context, CPF string) (model_account.Account, error)
}

type Options struct {
	Store store.Store
}

type account struct {
	store store.Store
}

func New(opts Options) IAccount {
	return account{
		store: opts.Store,
	}
}

func (a account) Create(ctx context.Context, account model_account.Create) (model_account.Account, error) {

	var emptyAccount model_account.Account

	account.CPF = utils.CleanCPF(account.CPF)

	if err := account.Valid(); err != nil {
		return emptyAccount, err
	}

	if _, err := a.store.Account.GetByCPF(ctx, account.CPF); err == nil {
		return emptyAccount, fmt.Errorf("alredy exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.MinCost)
	if err != nil {
		return emptyAccount, fmt.Errorf("any error")
	}

	account.Secret = string(hash)

	return a.store.Account.Create(ctx, account)
}

func (a account) BalanceByAccountID(ctx context.Context, ID string) (int, error) {
	account, err := a.store.Account.GetByID(ctx, ID)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (a account) List(ctx context.Context) ([]model_account.Account, error) {
	return a.store.Account.Get(ctx)
}

func (a account) GetByCPF(ctx context.Context, CPF string) (model_account.Account, error) {
	return a.store.Account.GetByCPF(ctx, CPF)
}
