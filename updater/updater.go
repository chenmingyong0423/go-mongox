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

	"github.com/chenmingyong0423/go-mongox/v2/internal/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"

	"github.com/chenmingyong0423/go-mongox/v2/callback"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
type iUpdater[T any] interface {
	UpdateOne(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error)
	Upsert(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error)
}

func NewUpdater[T any](collection *mongo.Collection) *Updater[T] {
	return &Updater[T]{collection: collection, filter: nil}
}

var _ iUpdater[any] = (*Updater[any])(nil)

type Updater[T any] struct {
	collection  *mongo.Collection
	filter      any
	updates     any
	replacement any
	modelHook   any
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn
}

// Filter is used to set the filter of the query
func (u *Updater[T]) Filter(filter any) *Updater[T] {
	u.filter = filter
	return u
}

// Updates is used to set the updates of the update
func (u *Updater[T]) Updates(updates any) *Updater[T] {
	u.updates = updates
	return u
}

func (u *Updater[T]) Replacement(replacement any) *Updater[T] {
	u.replacement = replacement
	return u
}

func (u *Updater[T]) ModelHook(modelHook any) *Updater[T] {
	u.modelHook = modelHook
	return u
}

func (u *Updater[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Updater[T] {
	u.beforeHooks = append(u.beforeHooks, hooks...)
	return u
}

func (u *Updater[T]) RegisterAfterHooks(hooks ...afterHookFn) *Updater[T] {
	u.afterHooks = append(u.afterHooks, hooks...)
	return u
}

func (u *Updater[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *BeforeOpContext, opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range u.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *AfterOpContext, opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range u.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Updater[T]) UpdateOne(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error) {
	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook))
	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateOne(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) UpdateMany(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error) {
	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook))
	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateMany(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) Upsert(ctx context.Context, opts ...options.Lister[options.UpdateOptions]) (*mongo.UpdateResult, error) {
	if len(opts) == 0 {
		opts = append(opts, options.Update().SetUpsert(true))
	} else {
		if uob, ok := opts[0].(*options.UpdateOptionsBuilder); ok {
			uob.Opts = append(uob.Opts, func(o *options.UpdateOptions) error {
				o.Upsert = utils.ToPtr(true)
				return nil
			})
		}
	}

	updates := bsonx.ToBsonM(u.updates)
	if len(updates) != 0 {
		u.updates = updates
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithDoc(new(T)), operation.WithFilter(u.filter), operation.WithUpdates(u.updates), operation.WithMongoOptions(opts), operation.WithModelHook(u.modelHook))

	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeBeforeUpsert)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateOne(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates), WithMongoOptions(opts), WithModelHook(u.modelHook))), operation.OpTypeAfterUpsert)
	if err != nil {
		return nil, err
	}
	return result, nil
}
