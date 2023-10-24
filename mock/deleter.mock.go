// Code generated by MockGen. DO NOT EDIT.
// Source: deleter.go
//
// Generated by this command:
//
//	mockgen -source=deleter.go -destination=../mock/deleter.mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockiDeleter is a mock of iDeleter interface.
type MockiDeleter[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockiDeleterMockRecorder[T]
}

// MockiDeleterMockRecorder is the mock recorder for MockiDeleter.
type MockiDeleterMockRecorder[T any] struct {
	mock *MockiDeleter[T]
}

// NewMockiDeleter creates a new mock instance.
func NewMockiDeleter[T any](ctrl *gomock.Controller) *MockiDeleter[T] {
	mock := &MockiDeleter[T]{ctrl: ctrl}
	mock.recorder = &MockiDeleterMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiDeleter[T]) EXPECT() *MockiDeleterMockRecorder[T] {
	return m.recorder
}

// DeleteMany mocks base method.
func (m *MockiDeleter[T]) DeleteMany(ctx context.Context) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMany", ctx)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockiDeleterMockRecorder[T]) DeleteMany(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockiDeleter[T])(nil).DeleteMany), ctx)
}

// DeleteManyWithOptions mocks base method.
func (m *MockiDeleter[T]) DeleteManyWithOptions(ctx context.Context, opts []*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteManyWithOptions", ctx, opts)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteManyWithOptions indicates an expected call of DeleteManyWithOptions.
func (mr *MockiDeleterMockRecorder[T]) DeleteManyWithOptions(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteManyWithOptions", reflect.TypeOf((*MockiDeleter[T])(nil).DeleteManyWithOptions), ctx, opts)
}

// DeleteOne mocks base method.
func (m *MockiDeleter[T]) DeleteOne(ctx context.Context) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", ctx)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockiDeleterMockRecorder[T]) DeleteOne(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockiDeleter[T])(nil).DeleteOne), ctx)
}

// DeleteOneWithOptions mocks base method.
func (m *MockiDeleter[T]) DeleteOneWithOptions(ctx context.Context, opts []*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneWithOptions", ctx, opts)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOneWithOptions indicates an expected call of DeleteOneWithOptions.
func (mr *MockiDeleterMockRecorder[T]) DeleteOneWithOptions(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneWithOptions", reflect.TypeOf((*MockiDeleter[T])(nil).DeleteOneWithOptions), ctx, opts)
}
