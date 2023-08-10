// Code generated by MockGen. DO NOT EDIT.
// Source: transfer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model_transfer "github.com/jorgepiresg/ChallangeStone/model/transfer"
)

// MockITransfer is a mock of ITransfer interface.
type MockITransfer struct {
	ctrl     *gomock.Controller
	recorder *MockITransferMockRecorder
}

// MockITransferMockRecorder is the mock recorder for MockITransfer.
type MockITransferMockRecorder struct {
	mock *MockITransfer
}

// NewMockITransfer creates a new mock instance.
func NewMockITransfer(ctrl *gomock.Controller) *MockITransfer {
	mock := &MockITransfer{ctrl: ctrl}
	mock.recorder = &MockITransferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransfer) EXPECT() *MockITransferMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockITransfer) Create(ctx context.Context, create model_transfer.DoTransfer) (model_transfer.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, create)
	ret0, _ := ret[0].(model_transfer.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockITransferMockRecorder) Create(ctx, create interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITransfer)(nil).Create), ctx, create)
}

// GetByAccountID mocks base method.
func (m *MockITransfer) GetByAccountID(ctx context.Context, AccountID string) ([]model_transfer.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccountID", ctx, AccountID)
	ret0, _ := ret[0].([]model_transfer.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccountID indicates an expected call of GetByAccountID.
func (mr *MockITransferMockRecorder) GetByAccountID(ctx, AccountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccountID", reflect.TypeOf((*MockITransfer)(nil).GetByAccountID), ctx, AccountID)
}