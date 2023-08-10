package model_transfer

import "time"

type Transfer struct {
	ID                   string    `db:"id"`
	AccountOriginID      string    `db:"account_origin_id"`
	AccountDestinationID string    `db:"account_destination_id"`
	Amount               int       `db:"amount"`
	CreatedAt            time.Time `db:"created_at"`
}

type DoTransfer struct {
	AccountOriginID      string `json:"account_origin_id" db:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id" db:"account_destination_id"`
	Amount               int    `json:"amount" db:"amount"`
}

type ResTransfer struct {
	Origin      Account
	Destination Account
}

type Account struct {
	ID      string
	Balance int
}

type DoTransferRequest struct {
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}
