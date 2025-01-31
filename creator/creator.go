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

package creator

import (
	"context"
	"reflect"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"

	"github.com/chenmingyong0423/go-mongox/v2/internal/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=creator.go -destination=../mock/creator.mock.go -package=mocks
type ICreator[T any] interface {
	InsertOne(ctx context.Context, docs *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error)
}

var _ ICreator[any] = (*Creator[any])(nil)

type Creator[T any] struct {
	collection *mongo.Collection

	modelHook any

	dbCallbacks *callback.Callback
	beforeHooks []hookFn[T]
	afterHooks  []hookFn[T]

	fields []*field.Filed
}

func NewCreator[T any](collection *mongo.Collection, dbCallbacks *callback.Callback, fields []*field.Filed) *Creator[T] {
	return &Creator[T]{
		collection:  collection,
		dbCallbacks: dbCallbacks,
		fields:      fields,
	}
}

func (c *Creator[T]) ModelHook(modelHook any) *Creator[T] {
	c.modelHook = modelHook
	return c
}

// RegisterBeforeHooks is used to set the after hooks of the insert operation
// If you register the hook for InsertOne, the opContext.Docs will be nil
// If you register the hook for InsertMany, the opContext.Doc will be nil
func (c *Creator[T]) RegisterBeforeHooks(hooks ...hookFn[T]) *Creator[T] {
	c.beforeHooks = append(c.beforeHooks, hooks...)
	return c
}

func (c *Creator[T]) RegisterAfterHooks(hooks ...hookFn[T]) *Creator[T] {
	c.afterHooks = append(c.afterHooks, hooks...)
	return c
}

func (c *Creator[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opType operation.OpType) error {
	err := c.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range c.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Creator[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext[T], opType operation.OpType) error {
	err := c.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range c.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Creator[T]) InsertOne(ctx context.Context, doc *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	currentTime := time.Now()
	docValue := reflect.ValueOf(doc)

	globalOpContext := operation.NewOpContext(c.collection, operation.WithDoc(doc), operation.WithReflectValue(docValue), operation.WithMongoOptions(opts), operation.WithModelHook(c.modelHook), operation.WithStartTime(currentTime), operation.WithFields(c.fields))
	opContext := NewOpContext(c.collection, WithDoc(doc), WithReflectValue[T](docValue), WithStartTime[T](currentTime), WithMongoOptions[T](opts), WithModelHook[T](c.modelHook), WithFields[T](c.fields))

	err := c.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeInsert)
	if err != nil {
		return nil, err
	}

	result, err := c.collection.InsertOne(ctx, doc, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = c.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterInsert)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Creator[T]) InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	currentTime := time.Now()
	docsValue := reflect.ValueOf(docs)

	globalOpContext := operation.NewOpContext(c.collection, operation.WithDoc(docs), operation.WithReflectValue(docsValue), operation.WithStartTime(currentTime), operation.WithMongoOptions(opts), operation.WithModelHook(c.modelHook), operation.WithFields(c.fields))
	opContext := NewOpContext(c.collection, WithDocs(docs), WithReflectValue[T](docsValue), WithStartTime[T](currentTime), WithMongoOptions[T](opts), WithModelHook[T](c.modelHook), WithFields[T](c.fields))

	err := c.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeInsert)
	if err != nil {
		return nil, err
	}

	result, err := c.collection.InsertMany(ctx, utils.ToAnySlice(docs...), opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = c.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterInsert)
	if err != nil {
		return nil, err
	}
	return result, nil
}
