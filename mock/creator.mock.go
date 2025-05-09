// Code generated by MockGen. DO NOT EDIT.
// Source: creator.go
//
// Generated by this command:
//
//	mockgen -source=creator.go -destination=../mock/creator.mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	creator "github.com/chenmingyong0423/go-mongox/v2/creator"
	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockICreator is a mock of ICreator interface.
type MockICreator[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockICreatorMockRecorder[T]
	isgomock struct{}
}

// MockICreatorMockRecorder is the mock recorder for MockICreator.
type MockICreatorMockRecorder[T any] struct {
	mock *MockICreator[T]
}

// NewMockICreator creates a new mock instance.
func NewMockICreator[T any](ctrl *gomock.Controller) *MockICreator[T] {
	mock := &MockICreator[T]{ctrl: ctrl}
	mock.recorder = &MockICreatorMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreator[T]) EXPECT() *MockICreatorMockRecorder[T] {
	return m.recorder
}

// GetCollection mocks base method.
func (m *MockICreator[T]) GetCollection() *mongo.Collection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*mongo.Collection)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockICreatorMockRecorder[T]) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockICreator[T])(nil).GetCollection))
}

// InsertMany mocks base method.
func (m *MockICreator[T]) InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, docs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertMany", varargs...)
	ret0, _ := ret[0].(*mongo.InsertManyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockICreatorMockRecorder[T]) InsertMany(ctx, docs any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, docs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockICreator[T])(nil).InsertMany), varargs...)
}

// InsertOne mocks base method.
func (m *MockICreator[T]) InsertOne(ctx context.Context, docs *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, docs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockICreatorMockRecorder[T]) InsertOne(ctx, docs any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, docs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockICreator[T])(nil).InsertOne), varargs...)
}

// ModelHook mocks base method.
func (m *MockICreator[T]) ModelHook(modelHook any) creator.ICreator[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelHook", modelHook)
	ret0, _ := ret[0].(creator.ICreator[T])
	return ret0
}

// ModelHook indicates an expected call of ModelHook.
func (mr *MockICreatorMockRecorder[T]) ModelHook(modelHook any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelHook", reflect.TypeOf((*MockICreator[T])(nil).ModelHook), modelHook)
}

// RegisterAfterHooks mocks base method.
func (m *MockICreator[T]) RegisterAfterHooks(hooks ...creator.HookFn[T]) creator.ICreator[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterAfterHooks", varargs...)
	ret0, _ := ret[0].(creator.ICreator[T])
	return ret0
}

// RegisterAfterHooks indicates an expected call of RegisterAfterHooks.
func (mr *MockICreatorMockRecorder[T]) RegisterAfterHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAfterHooks", reflect.TypeOf((*MockICreator[T])(nil).RegisterAfterHooks), hooks...)
}

// RegisterBeforeHooks mocks base method.
func (m *MockICreator[T]) RegisterBeforeHooks(hooks ...creator.HookFn[T]) creator.ICreator[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterBeforeHooks", varargs...)
	ret0, _ := ret[0].(creator.ICreator[T])
	return ret0
}

// RegisterBeforeHooks indicates an expected call of RegisterBeforeHooks.
func (mr *MockICreatorMockRecorder[T]) RegisterBeforeHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBeforeHooks", reflect.TypeOf((*MockICreator[T])(nil).RegisterBeforeHooks), hooks...)
}
