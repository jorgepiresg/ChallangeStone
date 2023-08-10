package account

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jorgepiresg/ChallangeStone/mocks"
	model_account "github.com/jorgepiresg/ChallangeStone/model/account"
	"github.com/jorgepiresg/ChallangeStone/store"
)

func TestCreate(t *testing.T) {

	type fields struct {
		account *mocks.MockIAccount
	}

	tests := map[string]struct {
		input    model_account.Create
		expected model_account.Account
		err      error
		prepare  func(f *fields)
	}{
		"success": {
			input: model_account.Create{
				Name:     "John Doe",
				CPF:      "111.111.111-11",
				Password: "any",
				Secret:   "secrect",
			},
			prepare: func(f *fields) {
				f.account.EXPECT().GetByCPF(gomock.Any(), "11111111111").Times(1).Return(model_account.Account{}, fmt.Errorf("any"))
				f.account.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(model_account.Account{
					ID: "id",
				}, nil)
			},
			expected: model_account.Account{
				ID: "id",
			},
		},
		"error: name invalid": {
			input:   model_account.Create{},
			prepare: func(f *fields) {},
			err:     fmt.Errorf("name invalid"),
		},

		"error: alredy exist": {
			input: model_account.Create{
				Name:     "John Doe",
				CPF:      "111.111.111-11",
				Password: "any",
				Secret:   "secrect",
			},
			prepare: func(f *fields) {
				f.account.EXPECT().GetByCPF(gomock.Any(), "11111111111").Times(1).Return(model_account.Account{}, nil)
			},
			err: fmt.Errorf("alredy exist"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountMock := mocks.NewMockIAccount(ctrl)

			tt.prepare(&fields{
				account: accountMock,
			})

			a := New(Options{
				Store: store.Store{
					Account: accountMock,
				},
			})

			res, err := a.Create(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if res != tt.expected {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}

		})
	}
}

func TestBalanceByAccountID(t *testing.T) {
	type fields struct {
		account *mocks.MockIAccount
	}

	tests := map[string]struct {
		input    string
		expected int
		err      error
		prepare  func(s *fields)
	}{
		"success": {
			input: "id",
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "id").Times(1).Return(model_account.Account{ID: "id", Balance: 5000}, nil)
			},
			expected: 5000,
		},
		"error": {
			input: "id",
			prepare: func(f *fields) {
				f.account.EXPECT().GetByID(gomock.Any(), "id").Times(1).Return(model_account.Account{}, fmt.Errorf("any"))
			},
			err: fmt.Errorf("any"),
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountMock := mocks.NewMockIAccount(ctrl)

			tt.prepare(&fields{
				account: accountMock,
			})

			a := New(Options{
				Store: store.Store{
					Account: accountMock,
				},
			})

			res, err := a.BalanceByAccountID(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}
		})
	}

}

func TestList(t *testing.T) {
	type fields struct {
		account *mocks.MockIAccount
	}

	tests := map[string]struct {
		expected []model_account.Account
		err      error
		prepare  func(s *fields)
	}{
		"success": {
			prepare: func(f *fields) {
				f.account.EXPECT().Get(gomock.Any()).Times(1).Return([]model_account.Account{{ID: "id"}}, nil)
			},
			expected: []model_account.Account{{ID: "id"}},
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountMock := mocks.NewMockIAccount(ctrl)

			tt.prepare(&fields{
				account: accountMock,
			})

			a := New(Options{
				Store: store.Store{
					Account: accountMock,
				},
			})

			res, err := a.List(context.Background())

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}
		})
	}
}

func TestGetByCPF(t *testing.T) {
	type fields struct {
		account *mocks.MockIAccount
	}

	tests := map[string]struct {
		input    string
		expected model_account.Account
		err      error
		prepare  func(s *fields)
	}{
		"success": {
			input: "11111111111",
			prepare: func(f *fields) {
				f.account.EXPECT().GetByCPF(gomock.Any(), "11111111111").Times(1).Return(model_account.Account{ID: "id"}, nil)
			},
			expected: model_account.Account{ID: "id"},
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			accountMock := mocks.NewMockIAccount(ctrl)

			tt.prepare(&fields{
				account: accountMock,
			})

			a := New(Options{
				Store: store.Store{
					Account: accountMock,
				},
			})

			res, err := a.GetByCPF(context.Background(), tt.input)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(`Expected err: "%s" got "%s"`, tt.err, err)
			}
			if res != tt.expected {
				t.Errorf("Expected result %v got %v", tt.expected, res)
			}
		})
	}
}
