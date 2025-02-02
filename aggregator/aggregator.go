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

package aggregator

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/field"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=aggregator.go -destination=../mock/aggregator.mock.go -package=mocks
type IAggregator[T any] interface {
	Aggregate(ctx context.Context, opts ...options.Lister[options.AggregateOptions]) ([]*T, error)
	AggregateWithParse(ctx context.Context, result any, opts ...options.Lister[options.AggregateOptions]) error
}

var _ IAggregator[any] = (*Aggregator[any])(nil)

type Aggregator[T any] struct {
	collection *mongo.Collection
	pipeline   any

	dbCallbacks *callback.Callback
	fields      []*field.Filed

	modelHook   any
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn
}

func NewAggregator[T any](collection *mongo.Collection, dbCallbacks *callback.Callback, fields []*field.Filed) *Aggregator[T] {
	return &Aggregator[T]{
		collection:  collection,
		dbCallbacks: dbCallbacks,
		fields:      fields,
	}
}

func (a *Aggregator[T]) ModelHook(modelHook any) *Aggregator[T] {
	a.modelHook = modelHook
	return a
}
func (a *Aggregator[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Aggregator[T] {
	a.beforeHooks = append(a.beforeHooks, hooks...)
	return a
}

func (a *Aggregator[T]) RegisterAfterHooks(hooks ...afterHookFn) *Aggregator[T] {
	a.afterHooks = append(a.afterHooks, hooks...)
	return a
}

func (a *Aggregator[T]) Pipeline(pipeline any) *Aggregator[T] {
	a.pipeline = pipeline
	return a
}

func (a *Aggregator[T]) Aggregate(ctx context.Context, opts ...options.Lister[options.AggregateOptions]) ([]*T, error) {
	currentTime := time.Now()
	globalOpContext := operation.NewOpContext(a.collection, operation.WithPipeline(a.pipeline), operation.WithMongoOptions(opts), operation.WithModelHook(a.modelHook), operation.WithStartTime(currentTime), operation.WithFields(a.fields))
	opContext := NewOpContext(a.collection, a.pipeline, WithMongoOptions(opts), WithModelHook(a.modelHook), WithStartTime(currentTime), WithFields(a.fields))

	err := a.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeInsert)
	if err != nil {
		return nil, err
	}

	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*T, 0)
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = cursor
	opContext.Result = cursor
	err = a.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AggregateWithParse is used to parse the result of the aggregation
// result must be a pointer to a slice
func (a *Aggregator[T]) AggregateWithParse(ctx context.Context, result any, opts ...options.Lister[options.AggregateOptions]) error {

	currentTime := time.Now()
	globalOpContext := operation.NewOpContext(a.collection, operation.WithPipeline(a.pipeline), operation.WithMongoOptions(opts), operation.WithModelHook(a.modelHook), operation.WithStartTime(currentTime), operation.WithFields(a.fields))
	opContext := NewOpContext(a.collection, a.pipeline, WithMongoOptions(opts), WithModelHook(a.modelHook), WithStartTime(currentTime), WithFields(a.fields))

	err := a.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeInsert)
	if err != nil {
		return err
	}

	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}

	globalOpContext.Result = cursor
	opContext.Result = cursor
	err = a.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterInsert)
	if err != nil {
		return err
	}

	return nil
}

func (a *Aggregator[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := a.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range a.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Aggregator[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := a.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range a.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}
