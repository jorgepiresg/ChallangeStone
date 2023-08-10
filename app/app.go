package app

import (
	"github.com/jorgepiresg/ChallangeStone/app/account"
	"github.com/jorgepiresg/ChallangeStone/app/transfer"
	"github.com/jorgepiresg/ChallangeStone/store"
)

type App struct {
	Account  account.IAccount
	Transfer transfer.ITransfer
}

type Options struct {
	Store store.Store
}

func New(opts Options) App {
	return App{
		Account:  account.New(account.Options{Store: opts.Store}),
		Transfer: transfer.New(transfer.Options{Store: opts.Store}),
	}
}
