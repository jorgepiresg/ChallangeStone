package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/jorgepiresg/ChallangeStone/store/account"
	"github.com/jorgepiresg/ChallangeStone/store/transfer"
)

type Store struct {
	Account  account.IAccount
	Transfer transfer.ITransfer
}

type Options struct {
	DB *sqlx.DB
}

func New(opts Options) Store {
	return Store{
		Account:  account.New(opts.DB),
		Transfer: transfer.New(opts.DB),
	}
}
