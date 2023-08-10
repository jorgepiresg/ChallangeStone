package transfer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	model_transfer "github.com/jorgepiresg/ChallangeStone/model/transfer"
)

//go:generate mockgen -source=$GOFILE -destination=../../mocks/transfer_mock.go -package=mocks
type ITransfer interface {
	Create(ctx context.Context, create model_transfer.DoTransfer) (model_transfer.Transfer, error)
	GetByAccountID(ctx context.Context, AccountID string) ([]model_transfer.Transfer, error)
}

type transfer struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) ITransfer {
	return transfer{
		db: db,
	}
}

func (t transfer) Create(ctx context.Context, create model_transfer.DoTransfer) (model_transfer.Transfer, error) {

	var transfer model_transfer.Transfer

	rows, err := t.db.NamedQueryContext(ctx, `INSERT INTO transfer (account_origin_id, account_destination_id, amount) VALUES (:account_origin_id, :account_destination_id, :amount) RETURNING *`, create)
	if err != nil {
		fmt.Println(err.Error())
		return transfer, err
	}

	for rows.Next() {
		err = rows.StructScan(&transfer)
		if err != nil {
			return transfer, err
		}
	}

	return transfer, nil
}

func (t transfer) GetByAccountID(ctx context.Context, AccountID string) ([]model_transfer.Transfer, error) {

	var tranfers []model_transfer.Transfer

	err := t.db.SelectContext(ctx, &tranfers, `SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfer`)
	if err != nil {
		return tranfers, err
	}

	return tranfers, nil
}
