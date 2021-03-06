// Code generated by MockGen. DO NOT EDIT.
// Source: conn_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgx "github.com/jackc/pgx/v4"
)

// MockFakePoolInterface is a mock of FakePoolInterface interface.
type MockFakePoolInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFakePoolInterfaceMockRecorder
}

// MockFakePoolInterfaceMockRecorder is the mock recorder for MockFakePoolInterface.
type MockFakePoolInterfaceMockRecorder struct {
	mock *MockFakePoolInterface
}

// NewMockFakePoolInterface creates a new mock instance.
func NewMockFakePoolInterface(ctrl *gomock.Controller) *MockFakePoolInterface {
	mock := &MockFakePoolInterface{ctrl: ctrl}
	mock.recorder = &MockFakePoolInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFakePoolInterface) EXPECT() *MockFakePoolInterfaceMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockFakePoolInterface) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockFakePoolInterfaceMockRecorder) BeginTx(ctx, txOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockFakePoolInterface)(nil).BeginTx), ctx, txOptions)
}
