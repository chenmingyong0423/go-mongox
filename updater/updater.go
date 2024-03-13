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

	"github.com/chenmingyong0423/go-mongox/callback"

	"github.com/chenmingyong0423/go-mongox/operation"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
type iUpdater[T any] interface {
	UpdateOne(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Upsert(ctx context.Context, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
}

func NewUpdater[T any](collection *mongo.Collection) *Updater[T] {
	return &Updater[T]{collection: collection, filter: nil}
}

type Updater[T any] struct {
	collection  *mongo.Collection
	filter      any
	updates     any
	replacement any
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

func (u *Updater[T]) UpdatesWithOperator(operator string, value any) *Updater[T] {
	u.updates = bson.D{bson.E{Key: operator, Value: value}}
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

func (u *Updater[T]) UpdateOne(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	globalOpContext := operation.NewOpContext(u.collection, operation.WithFilter(u.filter), operation.WithUpdate(u.updates))
	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates))), operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateOne(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates))), operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) UpdateMany(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	globalOpContext := operation.NewOpContext(u.collection, operation.WithFilter(u.filter), operation.WithUpdate(u.updates))
	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates))), operation.OpTypeBeforeUpdate)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.UpdateMany(ctx, u.filter, u.updates, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithUpdates(u.updates))), operation.OpTypeAfterUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Updater[T]) Upsert(ctx context.Context, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	if len(opts) == 0 {
		opts = append(opts, options.Replace().SetUpsert(true))
	} else {
		opts[0].SetUpsert(true)
	}

	globalOpContext := operation.NewOpContext(u.collection, operation.WithFilter(u.filter), operation.WithReplacement(u.replacement))

	err := u.preActionHandler(ctx, globalOpContext, NewBeforeOpContext(u.collection, NewCondContext(u.filter, WithReplacement(u.replacement))), operation.OpTypeBeforeUpsert)
	if err != nil {
		return nil, err
	}

	result, err := u.collection.ReplaceOne(ctx, u.filter, u.replacement, opts...)
	if err != nil {
		return nil, err
	}

	err = u.postActionHandler(ctx, globalOpContext, NewAfterOpContext(u.collection, NewCondContext(u.filter, WithReplacement(u.replacement))), operation.OpTypeAfterUpsert)
	if err != nil {
		return nil, err
	}
	return result, nil
}
