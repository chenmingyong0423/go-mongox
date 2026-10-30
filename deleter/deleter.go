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

package deleter

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=deleter.go -destination=../mock/deleter.mock.go -package=mocks
type IDeleter[T any] interface {
	DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error)
	UpdateOneOptions(opts ...options.Lister[options.UpdateOneOptions]) IDeleter[T]
	UpdateManyOptions(opts ...options.Lister[options.UpdateManyOptions]) IDeleter[T]
	Filter(filter any) IDeleter[T]
	ModelHook(modelHook any) IDeleter[T]
	RegisterAfterHooks(hooks ...AfterHookFn) IDeleter[T]
	RegisterBeforeHooks(hooks ...BeforeHookFn) IDeleter[T]
	PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error
	PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error
	GetCollection() *mongo.Collection
}

var _ IDeleter[any] = (*Deleter[any])(nil)

func NewDeleter[T any](collection *mongo.Collection, dbCallbacks *callback.Callback, fields []*field.Filed) *Deleter[T] {
	return &Deleter[T]{collection: collection, DBCallbacks: dbCallbacks, fields: fields}
}

type Deleter[T any] struct {
	collection *mongo.Collection
	fields     []*field.Filed

	filter    any
	modelHook any

	updateOneOptions  []options.Lister[options.UpdateOneOptions]
	updateManyOptions []options.Lister[options.UpdateManyOptions]

	DBCallbacks *callback.Callback
	BeforeHooks []BeforeHookFn
	AfterHooks  []AfterHookFn
}

func (d *Deleter[T]) RegisterBeforeHooks(hooks ...BeforeHookFn) IDeleter[T] {
	d.BeforeHooks = append(d.BeforeHooks, hooks...)
	return d
}

func (d *Deleter[T]) RegisterAfterHooks(hooks ...AfterHookFn) IDeleter[T] {
	d.AfterHooks = append(d.AfterHooks, hooks...)
	return d
}

func (d *Deleter[T]) PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := d.DBCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	opContext.Filter = globalOpContext.Filter
	opContext.Updates = globalOpContext.Updates
	for _, beforeHook := range d.BeforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	globalOpContext.Filter = opContext.Filter
	globalOpContext.Updates = opContext.Updates
	return nil
}

func (d *Deleter[T]) PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := d.DBCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range d.AfterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

// Filter is used to set the filter of the query
func (d *Deleter[T]) Filter(filter any) IDeleter[T] {
	d.filter = filter
	return d
}

func (d *Deleter[T]) ModelHook(modelHook any) IDeleter[T] {
	d.modelHook = modelHook
	return d
}

// UpdateOneOptions sets the driver options used when DeleteOne performs a soft deletion.
// Each call replaces the previous options. No options means using driver defaults.
func (d *Deleter[T]) UpdateOneOptions(opts ...options.Lister[options.UpdateOneOptions]) IDeleter[T] {
	d.updateOneOptions = opts
	return d
}

// UpdateManyOptions sets the driver options used when DeleteMany performs a soft deletion.
// Each call replaces the previous options. No options means using driver defaults.
func (d *Deleter[T]) UpdateManyOptions(opts ...options.Lister[options.UpdateManyOptions]) IDeleter[T] {
	d.updateManyOptions = opts
	return d
}

// DeleteOne marks a matching live document as deleted when T has an automatic
// deletion time field. Otherwise it physically deletes the document.
func (d *Deleter[T]) DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	return d.executeDelete(ctx, opts,
		func(opContext *OpContext) (*mongo.DeleteResult, error) {
			return d.collection.DeleteOne(ctx, opContext.Filter, opts...)
		},
		func(opContext *OpContext) (*mongo.UpdateResult, error) {
			return d.collection.UpdateOne(ctx, opContext.Filter, opContext.Updates, d.updateOneOptions...)
		},
	)
}

// DeleteMany marks matching live documents as deleted when T has an automatic
// deletion time field. Otherwise, it physically deletes matching documents.
func (d *Deleter[T]) DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	return d.executeDelete(ctx, opts,
		func(opContext *OpContext) (*mongo.DeleteResult, error) {
			return d.collection.DeleteMany(ctx, opContext.Filter, opts...)
		},
		func(opContext *OpContext) (*mongo.UpdateResult, error) {
			return d.collection.UpdateMany(ctx, opContext.Filter, opContext.Updates, d.updateManyOptions...)
		},
	)
}

// executeDelete runs the shared lifecycle; callers provide the typed driver operations.
func (d *Deleter[T]) executeDelete(ctx context.Context, opts any,
	deleteFn func(*OpContext) (*mongo.DeleteResult, error),
	updateFn func(*OpContext) (*mongo.UpdateResult, error),
) (*mongo.DeleteResult, error) {
	currentTime := time.Now()
	globalOpContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter), operation.WithMongoOptions(opts), operation.WithModelHook(d.modelHook), operation.WithFields(d.fields), operation.WithStartTime(currentTime))
	opContext := NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook), WithFields(d.fields), WithStartTime(currentTime))
	err := d.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	var result *mongo.DeleteResult
	if globalOpContext.SoftDelete {
		var updateResult *mongo.UpdateResult
		updateResult, err = updateFn(opContext)
		if err == nil {
			result = &mongo.DeleteResult{DeletedCount: updateResult.ModifiedCount, Acknowledged: updateResult.Acknowledged}
		}
	} else {
		result, err = deleteFn(opContext)
	}
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = d.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Deleter[T]) GetCollection() *mongo.Collection {
	return d.collection
}
