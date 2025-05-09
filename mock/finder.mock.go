// Code generated by MockGen. DO NOT EDIT.
// Source: finder.go
//
// Generated by this command:
//
//	mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	finder "github.com/chenmingyong0423/go-mongox/v2/finder"
	operation "github.com/chenmingyong0423/go-mongox/v2/operation"
	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockIFinder is a mock of IFinder interface.
type MockIFinder[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIFinderMockRecorder[T]
	isgomock struct{}
}

// MockIFinderMockRecorder is the mock recorder for MockIFinder.
type MockIFinderMockRecorder[T any] struct {
	mock *MockIFinder[T]
}

// NewMockIFinder creates a new mock instance.
func NewMockIFinder[T any](ctrl *gomock.Controller) *MockIFinder[T] {
	mock := &MockIFinder[T]{ctrl: ctrl}
	mock.recorder = &MockIFinderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFinder[T]) EXPECT() *MockIFinderMockRecorder[T] {
	return m.recorder
}

// Count mocks base method.
func (m *MockIFinder[T]) Count(ctx context.Context, opts ...options.Lister[options.CountOptions]) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIFinderMockRecorder[T]) Count(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIFinder[T])(nil).Count), varargs...)
}

// Distinct mocks base method.
func (m *MockIFinder[T]) Distinct(ctx context.Context, fieldName string, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult {
	m.ctrl.T.Helper()
	varargs := []any{ctx, fieldName}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Distinct", varargs...)
	ret0, _ := ret[0].(*mongo.DistinctResult)
	return ret0
}

// Distinct indicates an expected call of Distinct.
func (mr *MockIFinderMockRecorder[T]) Distinct(ctx, fieldName any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, fieldName}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*MockIFinder[T])(nil).Distinct), varargs...)
}

// DistinctWithParse mocks base method.
func (m *MockIFinder[T]) DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...options.Lister[options.DistinctOptions]) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, fieldName, result}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistinctWithParse", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistinctWithParse indicates an expected call of DistinctWithParse.
func (mr *MockIFinderMockRecorder[T]) DistinctWithParse(ctx, fieldName, result any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, fieldName, result}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistinctWithParse", reflect.TypeOf((*MockIFinder[T])(nil).DistinctWithParse), varargs...)
}

// Filter mocks base method.
func (m *MockIFinder[T]) Filter(filter any) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", filter)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockIFinderMockRecorder[T]) Filter(filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockIFinder[T])(nil).Filter), filter)
}

// Find mocks base method.
func (m *MockIFinder[T]) Find(ctx context.Context, opts ...options.Lister[options.FindOptions]) ([]*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIFinderMockRecorder[T]) Find(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIFinder[T])(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockIFinder[T]) FindOne(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockIFinderMockRecorder[T]) FindOne(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIFinder[T])(nil).FindOne), varargs...)
}

// FindOneAndUpdate mocks base method.
func (m *MockIFinder[T]) FindOneAndUpdate(ctx context.Context, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndUpdate", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockIFinderMockRecorder[T]) FindOneAndUpdate(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockIFinder[T])(nil).FindOneAndUpdate), varargs...)
}

// GetCollection mocks base method.
func (m *MockIFinder[T]) GetCollection() *mongo.Collection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*mongo.Collection)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockIFinderMockRecorder[T]) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockIFinder[T])(nil).GetCollection))
}

// Limit mocks base method.
func (m *MockIFinder[T]) Limit(limit int64) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *MockIFinderMockRecorder[T]) Limit(limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockIFinder[T])(nil).Limit), limit)
}

// ModelHook mocks base method.
func (m *MockIFinder[T]) ModelHook(modelHook any) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelHook", modelHook)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// ModelHook indicates an expected call of ModelHook.
func (mr *MockIFinderMockRecorder[T]) ModelHook(modelHook any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelHook", reflect.TypeOf((*MockIFinder[T])(nil).ModelHook), modelHook)
}

// PostActionHandler mocks base method.
func (m *MockIFinder[T]) PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *finder.OpContext[T], opTypes ...operation.OpType) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, globalOpContext, opContext}
	for _, a := range opTypes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostActionHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostActionHandler indicates an expected call of PostActionHandler.
func (mr *MockIFinderMockRecorder[T]) PostActionHandler(ctx, globalOpContext, opContext any, opTypes ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, globalOpContext, opContext}, opTypes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostActionHandler", reflect.TypeOf((*MockIFinder[T])(nil).PostActionHandler), varargs...)
}

// PreActionHandler mocks base method.
func (m *MockIFinder[T]) PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *finder.OpContext[T], opTypes ...operation.OpType) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, globalOpContext, opContext}
	for _, a := range opTypes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreActionHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreActionHandler indicates an expected call of PreActionHandler.
func (mr *MockIFinderMockRecorder[T]) PreActionHandler(ctx, globalOpContext, opContext any, opTypes ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, globalOpContext, opContext}, opTypes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreActionHandler", reflect.TypeOf((*MockIFinder[T])(nil).PreActionHandler), varargs...)
}

// RegisterAfterHooks mocks base method.
func (m *MockIFinder[T]) RegisterAfterHooks(hooks ...finder.AfterHookFn[T]) finder.IFinder[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterAfterHooks", varargs...)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// RegisterAfterHooks indicates an expected call of RegisterAfterHooks.
func (mr *MockIFinderMockRecorder[T]) RegisterAfterHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAfterHooks", reflect.TypeOf((*MockIFinder[T])(nil).RegisterAfterHooks), hooks...)
}

// RegisterBeforeHooks mocks base method.
func (m *MockIFinder[T]) RegisterBeforeHooks(hooks ...finder.BeforeHookFn[T]) finder.IFinder[T] {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range hooks {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterBeforeHooks", varargs...)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// RegisterBeforeHooks indicates an expected call of RegisterBeforeHooks.
func (mr *MockIFinderMockRecorder[T]) RegisterBeforeHooks(hooks ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBeforeHooks", reflect.TypeOf((*MockIFinder[T])(nil).RegisterBeforeHooks), hooks...)
}

// Skip mocks base method.
func (m *MockIFinder[T]) Skip(skip int64) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skip", skip)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// Skip indicates an expected call of Skip.
func (mr *MockIFinderMockRecorder[T]) Skip(skip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockIFinder[T])(nil).Skip), skip)
}

// Sort mocks base method.
func (m *MockIFinder[T]) Sort(sort any) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sort", sort)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// Sort indicates an expected call of Sort.
func (mr *MockIFinderMockRecorder[T]) Sort(sort any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockIFinder[T])(nil).Sort), sort)
}

// Updates mocks base method.
func (m *MockIFinder[T]) Updates(update any) finder.IFinder[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", update)
	ret0, _ := ret[0].(finder.IFinder[T])
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockIFinderMockRecorder[T]) Updates(update any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockIFinder[T])(nil).Updates), update)
}
