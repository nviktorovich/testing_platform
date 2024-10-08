// Code generated by MockGen. DO NOT EDIT.
// Source: ./resolve_strategy.go

// Package testdata is a generated GoMock package.
package testdata

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResolveStrategy is a mock of ResolveStrategy interface.
type MockResolveStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockResolveStrategyMockRecorder
}

// MockResolveStrategyMockRecorder is the mock recorder for MockResolveStrategy.
type MockResolveStrategyMockRecorder struct {
	mock *MockResolveStrategy
}

// NewMockResolveStrategy creates a new mock instance.
func NewMockResolveStrategy(ctrl *gomock.Controller) *MockResolveStrategy {
	mock := &MockResolveStrategy{ctrl: ctrl}
	mock.recorder = &MockResolveStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolveStrategy) EXPECT() *MockResolveStrategyMockRecorder {
	return m.recorder
}

// Resolve mocks base method.
func (m *MockResolveStrategy) Resolve(correct, current interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", correct, current)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve.
func (mr *MockResolveStrategyMockRecorder) Resolve(correct, current interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockResolveStrategy)(nil).Resolve), correct, current)
}
