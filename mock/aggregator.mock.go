// Code generated by MockGen. DO NOT EDIT.
// Source: aggregator.go
//
// Generated by this command:
//
//	mockgen -source=aggregator.go -destination=../mock/aggregator.mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	options "go.mongodb.org/mongo-driver/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockIAggregator is a mock of IAggregator interface.
type MockIAggregator[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIAggregatorMockRecorder[T]
}

// MockIAggregatorMockRecorder is the mock recorder for MockIAggregator.
type MockIAggregatorMockRecorder[T any] struct {
	mock *MockIAggregator[T]
}

// NewMockIAggregator creates a new mock instance.
func NewMockIAggregator[T any](ctrl *gomock.Controller) *MockIAggregator[T] {
	mock := &MockIAggregator[T]{ctrl: ctrl}
	mock.recorder = &MockIAggregatorMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAggregator[T]) EXPECT() *MockIAggregatorMockRecorder[T] {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockIAggregator[T]) Aggregate(ctx context.Context, opts ...*options.AggregateOptions) ([]*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Aggregate", varargs...)
	ret0, _ := ret[0].([]*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockIAggregatorMockRecorder[T]) Aggregate(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockIAggregator[T])(nil).Aggregate), varargs...)
}

// AggregateWithParse mocks base method.
func (m *MockIAggregator[T]) AggregateWithParse(ctx context.Context, result any, opts ...*options.AggregateOptions) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, result}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AggregateWithParse", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AggregateWithParse indicates an expected call of AggregateWithParse.
func (mr *MockIAggregatorMockRecorder[T]) AggregateWithParse(ctx, result any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, result}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateWithParse", reflect.TypeOf((*MockIAggregator[T])(nil).AggregateWithParse), varargs...)
}
