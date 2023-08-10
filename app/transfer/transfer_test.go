package transfer

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jorgepiresg/ChallangeStone/mocks"
	model_account "github.com/jorgepiresg/ChallangeStone/model/account"
	model_transfer "github.com/jorgepiresg/ChallangeStone/model/transfer"
	"github.com/jorgepiresg/ChallangeStone/store"
)

func TestGetByAccountID(t *testing.T) {

	type fields struct {
		transfer *mocks.MockITransfer
	}

	tests := map[string]struct {
		input    string
		expected []model_transfer.Transfer
		err      error
		prepare  func(f *fields)
	}{
		"success": {
			input: "id",
			prepare: func(f *fields) {
				f.transfer.EXPECT().GetByAccountID(gomock.Any(), "id").Times(1).Return([]model_transfer.Transfer{
					{
						ID:                   "id",
						AccountOriginID:      "account_origin_id",
						AccountDestinationID: "account_destination_id",
						Amount:               5000,
					}}, nil)
			},
			expected: []model_transfer.Transfer{
				{
					ID:                   "id",
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
					Amount:               5000,
				}},
		},
		"error": {
			input: "id",
			prepare: func(f *fields) {
				f.transfer.EXPECT().GetByAccountID(gomock.Any(), "id").Times(1).Return(nil, fmt.Errorf("any"))
			},
			err: fmt.Errorf("any"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			transferMock := mocks.NewMockITransfer(ctrl)

			tt.prepare(&fields{
				transfer: transferMock,
			})

			a := New(Options{
				Store: store.Store{
					Transfer: transferMock,
				},
			})

			res, err := a.GetByAccountID(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}

		})
	}
}

func TestDo(t *testing.T) {

	type fields struct {
		transfer *mocks.MockITransfer
		account  *mocks.MockIAccount
	}

	tests := map[string]struct {
		input   model_transfer.DoTransfer
		err     error
		prepare func(f *fields)
	}{
		"success": {
			input: model_transfer.DoTransfer{
				AccountOriginID:      "account_origin_id",
				AccountDestinationID: "account_destination_id",
				Amount:               5000,
			},
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "account_origin_id").Times(1).Return(model_account.Account{
					Balance: 10000,
				}, nil)

				f.transfer.EXPECT().Create(gomock.Any(), model_transfer.DoTransfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
					Amount:               5000,
				}).Times(1).Return(model_transfer.Transfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
				}, nil)

				f.account.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(2).Return(model_account.Account{}, nil)

				f.account.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Times(2).Return(nil)
			},
		},
		"error: account origin not found": {
			input: model_transfer.DoTransfer{},
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "").Times(1).Return(model_account.Account{}, fmt.Errorf("account origin not found"))
			},
			err: fmt.Errorf("account origin not found"),
		},
		"error: amount invalid": {
			input: model_transfer.DoTransfer{
				AccountOriginID:      "account_origin_id",
				AccountDestinationID: "account_destination_id",
				Amount:               5000,
			},
			err: fmt.Errorf("amount invalid"),
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "account_origin_id").Times(1).Return(model_account.Account{
					Balance: 4999,
				}, nil)
			},
		},
		"error: fail to do transfer": {
			input: model_transfer.DoTransfer{
				AccountOriginID:      "account_origin_id",
				AccountDestinationID: "account_destination_id",
				Amount:               5000,
			},
			err: fmt.Errorf("fail to do transfer"),
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "account_origin_id").Times(1).Return(model_account.Account{
					Balance: 10000,
				}, nil)

				f.transfer.EXPECT().Create(gomock.Any(), model_transfer.DoTransfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
					Amount:               5000,
				}).Times(1).Return(model_transfer.Transfer{}, fmt.Errorf("error any"))
			},
		},
		"error: fail to update balance account origin": {
			input: model_transfer.DoTransfer{
				AccountOriginID:      "account_origin_id",
				AccountDestinationID: "account_destination_id",
				Amount:               5000,
			},
			err: fmt.Errorf("fail to update balance account origin"),
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "account_origin_id").Times(1).Return(model_account.Account{
					Balance: 10000,
				}, nil)

				f.transfer.EXPECT().Create(gomock.Any(), model_transfer.DoTransfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
					Amount:               5000,
				}).Times(1).Return(model_transfer.Transfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
				}, nil)

				f.account.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(1).Return(model_account.Account{}, fmt.Errorf("any error"))
			},
		},
		"error: fail to update balance destination origin": {
			input: model_transfer.DoTransfer{
				AccountOriginID:      "account_origin_id",
				AccountDestinationID: "account_destination_id",
				Amount:               5000,
			},
			err: fmt.Errorf("fail to update balance destination origin"),
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "account_origin_id").Times(1).Return(model_account.Account{
					Balance: 10000,
				}, nil)

				f.transfer.EXPECT().Create(gomock.Any(), model_transfer.DoTransfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
					Amount:               5000,
				}).Times(1).Return(model_transfer.Transfer{
					AccountOriginID:      "account_origin_id",
					AccountDestinationID: "account_destination_id",
				}, nil)

				f.account.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(1).Return(model_account.Account{}, nil)

				f.account.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)

				f.account.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(1).Return(model_account.Account{}, fmt.Errorf("any error"))
			},
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountMock := mocks.NewMockIAccount(ctrl)
			transferMock := mocks.NewMockITransfer(ctrl)

			tt.prepare(&fields{
				account:  accountMock,
				transfer: transferMock,
			})

			a := New(Options{
				Store: store.Store{
					Account:  accountMock,
					Transfer: transferMock,
				},
			})

			err := a.Do(context.Background(), tt.input)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}

		})
	}
}
