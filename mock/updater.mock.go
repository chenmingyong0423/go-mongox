// Code generated by MockGen. DO NOT EDIT.
// Source: updater.go
//
// Generated by this command:
//
//	mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	operation "github.com/chenmingyong0423/go-mongox/v2/operation"
	updater "github.com/chenmingyong0423/go-mongox/v2/updater"
	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockIUpdater is a mock of IUpdater interface.
type MockIUpdater[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdaterMockRecorder[T]
	isgomock struct{}
}

// MockIUpdaterMockRecorder is the mock recorder for MockIUpdater.
type MockIUpdaterMockRecorder[T any] struct {
	mock *MockIUpdater[T]
}

// NewMockIUpdater creates a new mock instance.
func NewMockIUpdater[T any](ctrl *gomock.Controller) *MockIUpdater[T] {
	mock := &MockIUpdater[T]{ctrl: ctrl}
	mock.recorder = &MockIUpdaterMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUpdater[T]) EXPECT() *MockIUpdaterMockRecorder[T] {
	return m.recorder
}

// Filter mocks base method.
func (m *MockIUpdater[T]) Filter(filter any) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", filter)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockIUpdaterMockRecorder[T]) Filter(filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockIUpdater[T])(nil).Filter), filter)
}

// GetCollection mocks base method.
func (m *MockIUpdater[T]) GetCollection() *mongo.Collection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*mongo.Collection)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockIUpdaterMockRecorder[T]) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockIUpdater[T])(nil).GetCollection))
}

// ModelHook mocks base method.
func (m *MockIUpdater[T]) ModelHook(modelHook any) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelHook", modelHook)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// ModelHook indicates an expected call of ModelHook.
func (mr *MockIUpdaterMockRecorder[T]) ModelHook(modelHook any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelHook", reflect.TypeOf((*MockIUpdater[T])(nil).ModelHook), modelHook)
}

// PostActionHandler mocks base method.
func (m *MockIUpdater[T]) PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *updater.OpContext, opType operation.OpType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostActionHandler", ctx, globalOpContext, opContext, opType)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostActionHandler indicates an expected call of PostActionHandler.
func (mr *MockIUpdaterMockRecorder[T]) PostActionHandler(ctx, globalOpContext, opContext, opType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostActionHandler", reflect.TypeOf((*MockIUpdater[T])(nil).PostActionHandler), ctx, globalOpContext, opContext, opType)
}

// PreActionHandler mocks base method.
func (m *MockIUpdater[T]) PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *updater.OpContext, opType operation.OpType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreActionHandler", ctx, globalOpContext, opContext, opType)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreActionHandler indicates an expected call of PreActionHandler.
func (mr *MockIUpdaterMockRecorder[T]) PreActionHandler(ctx, globalOpContext, opContext, opType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreActionHandler", reflect.TypeOf((*MockIUpdater[T])(nil).PreActionHandler), ctx, globalOpContext, opContext, opType)
}

// RegisterAfterHooks mocks base method.
func (m *MockIUpdater[T]) RegisterAfterHooks(hooks ...updater.AfterHookFn) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterAfterHooks", varargs...)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// RegisterAfterHooks indicates an expected call of RegisterAfterHooks.
func (mr *MockIUpdaterMockRecorder[T]) RegisterAfterHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAfterHooks", reflect.TypeOf((*MockIUpdater[T])(nil).RegisterAfterHooks), hooks...)
}

// RegisterBeforeHooks mocks base method.
func (m *MockIUpdater[T]) RegisterBeforeHooks(hooks ...updater.BeforeHookFn) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterBeforeHooks", varargs...)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// RegisterBeforeHooks indicates an expected call of RegisterBeforeHooks.
func (mr *MockIUpdaterMockRecorder[T]) RegisterBeforeHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBeforeHooks", reflect.TypeOf((*MockIUpdater[T])(nil).RegisterBeforeHooks), hooks...)
}

// Replacement mocks base method.
func (m *MockIUpdater[T]) Replacement(replacement any) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replacement", replacement)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// Replacement indicates an expected call of Replacement.
func (mr *MockIUpdaterMockRecorder[T]) Replacement(replacement any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replacement", reflect.TypeOf((*MockIUpdater[T])(nil).Replacement), replacement)
}

// UpdateMany mocks base method.
func (m *MockIUpdater[T]) UpdateMany(ctx context.Context, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMany", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMany indicates an expected call of UpdateMany.
func (mr *MockIUpdaterMockRecorder[T]) UpdateMany(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockIUpdater[T])(nil).UpdateMany), varargs...)
}

// UpdateOne mocks base method.
func (m *MockIUpdater[T]) UpdateOne(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockIUpdaterMockRecorder[T]) UpdateOne(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockIUpdater[T])(nil).UpdateOne), varargs...)
}

// Updates mocks base method.
func (m *MockIUpdater[T]) Updates(updates any) updater.IUpdater[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", updates)
	ret0, _ := ret[0].(updater.IUpdater[T])
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockIUpdaterMockRecorder[T]) Updates(updates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockIUpdater[T])(nil).Updates), updates)
}

// Upsert mocks base method.
func (m *MockIUpdater[T]) Upsert(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upsert", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockIUpdaterMockRecorder[T]) Upsert(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockIUpdater[T])(nil).Upsert), varargs...)
}
