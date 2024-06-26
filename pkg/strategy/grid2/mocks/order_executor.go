// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/puper/bbgo/pkg/strategy/grid2 (interfaces: OrderExecutor)
//
// Generated by this command:
//
//	mockgen -destination=mocks/order_executor.go -package=mocks . OrderExecutor
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bbgo "github.com/puper/bbgo/pkg/bbgo"
	fixedpoint "github.com/puper/bbgo/pkg/fixedpoint"
	types "github.com/puper/bbgo/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderExecutor is a mock of OrderExecutor interface.
type MockOrderExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockOrderExecutorMockRecorder
}

// MockOrderExecutorMockRecorder is the mock recorder for MockOrderExecutor.
type MockOrderExecutorMockRecorder struct {
	mock *MockOrderExecutor
}

// NewMockOrderExecutor creates a new mock instance.
func NewMockOrderExecutor(ctrl *gomock.Controller) *MockOrderExecutor {
	mock := &MockOrderExecutor{ctrl: ctrl}
	mock.recorder = &MockOrderExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderExecutor) EXPECT() *MockOrderExecutorMockRecorder {
	return m.recorder
}

// ActiveMakerOrders mocks base method.
func (m *MockOrderExecutor) ActiveMakerOrders() *bbgo.ActiveOrderBook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveMakerOrders")
	ret0, _ := ret[0].(*bbgo.ActiveOrderBook)
	return ret0
}

// ActiveMakerOrders indicates an expected call of ActiveMakerOrders.
func (mr *MockOrderExecutorMockRecorder) ActiveMakerOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveMakerOrders", reflect.TypeOf((*MockOrderExecutor)(nil).ActiveMakerOrders))
}

// ClosePosition mocks base method.
func (m *MockOrderExecutor) ClosePosition(arg0 context.Context, arg1 fixedpoint.Value, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClosePosition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClosePosition indicates an expected call of ClosePosition.
func (mr *MockOrderExecutorMockRecorder) ClosePosition(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePosition", reflect.TypeOf((*MockOrderExecutor)(nil).ClosePosition), varargs...)
}

// GracefulCancel mocks base method.
func (m *MockOrderExecutor) GracefulCancel(arg0 context.Context, arg1 ...types.Order) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GracefulCancel", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GracefulCancel indicates an expected call of GracefulCancel.
func (mr *MockOrderExecutorMockRecorder) GracefulCancel(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GracefulCancel", reflect.TypeOf((*MockOrderExecutor)(nil).GracefulCancel), varargs...)
}

// SubmitOrders mocks base method.
func (m *MockOrderExecutor) SubmitOrders(arg0 context.Context, arg1 ...types.SubmitOrder) (types.OrderSlice, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitOrders", varargs...)
	ret0, _ := ret[0].(types.OrderSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrders indicates an expected call of SubmitOrders.
func (mr *MockOrderExecutorMockRecorder) SubmitOrders(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrders", reflect.TypeOf((*MockOrderExecutor)(nil).SubmitOrders), varargs...)
}
