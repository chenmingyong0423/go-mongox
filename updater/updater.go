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

package updater

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"

	"github.com/chenmingyong0423/go-mongox/v2/internal/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
type IUpdater[T any] interface {
	UpdateOne(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error)
	Upsert(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)
	Filter(filter any) IUpdater[T]
	ModelHook(modelHook any) IUpdater[T]
	RegisterAfterHooks(hooks ...AfterHookFn) IUpdater[T]
	RegisterBeforeHooks(hooks ...BeforeHookFn) IUpdater[T]
	Replacement(replacement any) IUpdater[T]
	Updates(updates any) IUpdater[T]
	PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error
	PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error
	GetCollection() *mongo.Collection
}

func NewUpdater[T any](collection *mongo.Collection, dbCallbacks *callback.Callback, fields []*field.Filed) *Updater[T] {
	return &Updater[T]{collection: collection, DBCallbacks: dbCallbacks, fields: fields}
}

var _ IUpdater[any] = (*Updater[any])(nil)

type Updater[T any] struct {
	collection *mongo.Collection
	fields     []*field.Filed

	filter      any
	updates     any
	replacement any
	modelHook   any

	DBCallbacks *callback.Callback
	BeforeHooks []BeforeHookFn
	AfterHooks  []AfterHookFn
}

// Filter is used to set the filter of the query
func (u *Updater[T]) Filter(filter any) IUpdater[T] {
	u.filter = filter
	return u
}

// Updates is used to set the updates of the update
func (u *Updater[T]) Updates(updates any) IUpdater[T] {
	u.updates = updates
	return u
}

func (u *Updater[T]) Replacement(replacement any) IUpdater[T] {
	u.replacement = replacement
	return u
}

func (u *Updater[T]) ModelHook(modelHook any) IUpdater[T] {
	u.modelHook = modelHook
	return u
}

func (u *Updater[T]) RegisterBeforeHooks(hooks ...BeforeHookFn) IUpdater[T] {
	u.BeforeHooks = append(u.BeforeHooks, hooks...)
	return u
}

func (u *Updater[T]) RegisterAfterHooks(hooks ...AfterHookFn) IUpdater[T] {
	u.AfterHooks = append(u.AfterHooks, hooks...)
	return u
}

func (u *Updater[T]) PreActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := u.DBCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range u.BeforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater[T]) PostActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := u.DBCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range u.AfterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater[T]) UpdateOne(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {

	currentTime := time.Now()

	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook), operation.WithFields(u.fields), operation.WithStartTime(currentTime))
	opContext := NewOpContext(u.collection, u.filter, u.updates, WithMongoOptions(opts), WithModelHook(u.modelHook), WithFields(u.fields), WithStartTime(currentTime))
	err := u.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateOne(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = u.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) UpdateMany(ctx context.Context, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	currentTime := time.Now()

	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook), operation.WithFields(u.fields), operation.WithStartTime(currentTime))
	opContext := NewOpContext(u.collection, u.filter, u.updates, WithMongoOptions(opts), WithModelHook(u.modelHook), WithFields(u.fields), WithStartTime(currentTime))

	err := u.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateMany(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = u.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) Upsert(ctx context.Context, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	currentTime := time.Now()

	if len(opts) == 0 {
		opts = append(opts, options.UpdateOne().SetUpsert(true))
	} else {
		if uob, ok := opts[0].(*options.UpdateOneOptionsBuilder); ok {
			uob.Opts = append(uob.Opts, func(o *options.UpdateOneOptions) error {
				o.Upsert = utils.ToPtr(true)
				return nil
			})
		}
	}

	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook), operation.WithStartTime(currentTime), operation.WithFields(u.fields))
	opContext := NewOpContext(u.collection, u.filter, u.updates, WithMongoOptions(opts), WithModelHook(u.modelHook), WithStartTime(currentTime), WithFields(u.fields))
	err := u.PreActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeUpsert)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateOne(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = u.PostActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterUpsert)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *Updater[T]) GetCollection() *mongo.Collection {
	return u.collection
}
