// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package finder

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"
	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
type IFinder[T any] interface {
	FindOne(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*T, error)
	Find(ctx context.Context, opts ...options.Lister[options.FindOptions]) ([]*T, error)
	Count(ctx context.Context, opts ...options.Lister[options.CountOptions]) (int64, error)
	Distinct(ctx context.Context, fieldName string, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult
	DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...options.Lister[options.DistinctOptions]) error
	Filter(filter any) IFinder[T]
	FindOneAndUpdate(ctx context.Context, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error)
	Limit(limit int64) IFinder[T]
	ModelHook(modelHook any) IFinder[T]
	RegisterAfterHooks(hooks ...AfterHookFn[T]) IFinder[T]
	RegisterBeforeHooks(hooks ...BeforeHookFn[T]) IFinder[T]
	Skip(skip int64) IFinder[T]
	Sort(sort any) IFinder[T]
	Updates(update any) IFinder[T]
	PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opTypes ...operation.OpType) (err error)
	PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opTypes ...operation.OpType) (err error)
	GetCollection() *mongo.Collection
}

func NewFinder[T any](collection *mongo.Collection, callbacks *callback.Callback, fields []*field.Filed) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}, dbCallbacks: callbacks, fields: fields}
}

var _ IFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection *mongo.Collection
	filter     any
	updates    any
	modelHook  any

	fields      []*field.Filed
	dbCallbacks *callback.Callback
	beforeHooks []BeforeHookFn[T]
	afterHooks  []AfterHookFn[T]

	skip, limit int64
	sort        any
}

func (f *Finder[T]) RegisterBeforeHooks(hooks ...BeforeHookFn[T]) IFinder[T] {
	f.beforeHooks = append(f.beforeHooks, hooks...)
	return f
}

// RegisterAfterHooks is used to set the after hooks of the query
// If you register the hook for FindOne, the opContext.Docs will be nil
// If you register the hook for Find, the opContext.Doc will be nil
func (f *Finder[T]) RegisterAfterHooks(hooks ...AfterHookFn[T]) IFinder[T] {
	f.afterHooks = append(f.afterHooks, hooks...)
	return f
}

// Filter is used to set the filter of the query
func (f *Finder[T]) Filter(filter any) IFinder[T] {
	f.filter = filter
	return f
}

func (f *Finder[T]) Limit(limit int64) IFinder[T] {
	f.limit = limit
	return f
}

func (f *Finder[T]) Skip(skip int64) IFinder[T] {
	f.skip = skip
	return f
}

func (f *Finder[T]) Sort(sort any) IFinder[T] {
	f.sort = sort
	return f
}

func (f *Finder[T]) Updates(update any) IFinder[T] {
	f.updates = update
	return f
}

func (f *Finder[T]) ModelHook(modelHook any) IFinder[T] {
	f.modelHook = modelHook
	return f
}

func (f *Finder[T]) PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opTypes ...operation.OpType) (err error) {
	for _, opType := range opTypes {
		err = f.dbCallbacks.Execute(ctx, globalOpContext, opType)
		if err != nil {
			return
		}
	}
	for _, beforeHook := range f.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return
}

func (f *Finder[T]) PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opTypes ...operation.OpType) (err error) {
	for _, opType := range opTypes {
		err = f.dbCallbacks.Execute(ctx, globalOpContext, opType)
		if err != nil {
			return
		}
	}
	for _, afterHook := range f.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return
		}
	}
	return
}

func (f *Finder[T]) FindOne(ctx context.Context, opts ...options.Lister[options.FindOneOptions]) (*T, error) {
	currentTime := time.Now()
	if f.sort != nil {
		opts = append(opts, options.FindOne().SetSort(f.sort))
	}

	t := new(T)

	globalOpContext := operation.NewOpContext(f.collection, operation.WithFilter(f.filter), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook), operation.WithStartTime(currentTime), operation.WithFields(f.fields))
	opContext := NewOpContext(f.collection, f.filter, WithMongoOptions[T](opts), WithModelHook[T](f.modelHook), WithStartTime[T](currentTime), WithFields[T](f.fields))
	err := f.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeFind)
	if err != nil {
		return nil, err
	}

	result := f.collection.FindOne(ctx, f.filter, opts...)
	err = result.Decode(t)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	globalOpContext.Doc = t
	opContext.Result = result
	opContext.Doc = t
	err = f.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Find(ctx context.Context, opts ...options.Lister[options.FindOptions]) ([]*T, error) {
	currentTime := time.Now()

	if f.sort != nil {
		opts = append(opts, options.Find().SetSort(f.sort))
	}
	if f.skip != 0 {
		opts = append(opts, options.Find().SetSkip(f.skip))
	}
	if f.limit != 0 {
		opts = append(opts, options.Find().SetLimit(f.limit))
	}

	t := make([]*T, 0)

	globalOpContext := operation.NewOpContext(f.collection, operation.WithFilter(f.filter), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook), operation.WithStartTime(currentTime), operation.WithFields(f.fields))
	opContext := NewOpContext(f.collection, f.filter, WithMongoOptions[T](opts), WithModelHook[T](f.modelHook), WithStartTime[T](currentTime), WithFields[T](f.fields))
	err := f.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeFind)
	if err != nil {
		return nil, err
	}

	cursor, err := f.collection.Find(ctx, f.filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &t)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = cursor
	globalOpContext.Doc = t
	opContext.Result = cursor
	opContext.Docs = t
	err = f.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Count(ctx context.Context, opts ...options.Lister[options.CountOptions]) (int64, error) {
	return f.collection.CountDocuments(ctx, f.filter, opts...)
}

func (f *Finder[T]) Distinct(ctx context.Context, fieldName string, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult {
	return f.collection.Distinct(ctx, fieldName, f.filter, opts...)
}

// DistinctWithParse is used to parse the result of Distinct
// result must be a pointer
func (f *Finder[T]) DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...options.Lister[options.DistinctOptions]) error {
	distinctResult := f.collection.Distinct(ctx, fieldName, f.filter, opts...)
	if distinctResult.Err() != nil {
		return distinctResult.Err()
	}
	err := distinctResult.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (f *Finder[T]) FindOneAndUpdate(ctx context.Context, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error) {
	currentTime := time.Now()
	t := new(T)

	updates := bsonx.ToBsonM(f.updates)
	if len(updates) != 0 {
		f.updates = updates
	}

	globalOpContext := operation.NewOpContext(f.collection, operation.WithFilter(f.filter), operation.WithUpdates(f.updates), operation.WithMongoOptions(opts), operation.WithModelHook(f.modelHook), operation.WithStartTime(currentTime), operation.WithFields(f.fields))
	opContext := NewOpContext(f.collection, f.filter, WithUpdates[T](f.updates), WithMongoOptions[T](opts), WithModelHook[T](f.modelHook), WithStartTime[T](currentTime), WithFields[T](f.fields))

	err := f.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeFind, operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result := f.collection.FindOneAndUpdate(ctx, f.filter, f.updates, opts...)
	err = result.Decode(t)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	globalOpContext.Doc = t
	opContext.Result = result
	opContext.Doc = t
	err = f.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterFind, operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) GetCollection() *mongo.Collection {
	return f.collection
}
