package transfer

import (
	"context"
	"fmt"

	model_transfer "github.com/jorgepiresg/ChallangeStone/model/transfer"
	"github.com/jorgepiresg/ChallangeStone/store"
)

type ITransfer interface {
	GetByAccountID(ctx context.Context, accountID string) ([]model_transfer.Transfer, error)
	Do(ctx context.Context, data model_transfer.DoTransfer) error
}

type Options struct {
	Store store.Store
}

type transfer struct {
	store store.Store
}

func New(opts Options) ITransfer {
	return transfer{
		store: opts.Store,
	}
}

func (t transfer) GetByAccountID(ctx context.Context, accountID string) ([]model_transfer.Transfer, error) {
	return t.store.Transfer.GetByAccountID(ctx, accountID)
}

func (t transfer) Do(ctx context.Context, data model_transfer.DoTransfer) error {

	accountOriginID, err := t.store.Account.GetByID(ctx, data.AccountOriginID)
	if err != nil {
		return fmt.Errorf("account origin not found")
	}

	var validateBalance = func(accountOriginIDBalance, amount int) bool {
		return accountOriginIDBalance >= amount
	}

	if !validateBalance(accountOriginID.Balance, data.Amount) {
		return fmt.Errorf("amount invalid")
	}

	transfer, err := t.store.Transfer.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("fail to do transfer")
	}

	err = t.updateAccount(ctx, transfer.AccountOriginID, -data.Amount)
	if err != nil {
		return fmt.Errorf("fail to update balance account origin")
	}

	err = t.updateAccount(ctx, transfer.AccountDestinationID, data.Amount)
	if err != nil {
		return fmt.Errorf("fail to update balance destination origin")
	}

	return nil
}

func (t transfer) updateAccount(ctx context.Context, AccountID string, amount int) error {

	account, err := t.store.Account.GetByID(ctx, AccountID)
	if err != nil {
		return err
	}

	newBalance := account.Balance + amount

	return t.store.Account.UpdateBalance(ctx, AccountID, newBalance)
}
