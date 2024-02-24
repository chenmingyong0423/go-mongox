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

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
type iFinder[T any] interface {
	FindOne(ctx context.Context, opts ...*options.FindOneOptions) (*T, error)
	Find(ctx context.Context, opts ...*options.FindOptions) ([]*T, error)
	Count(ctx context.Context, opts ...*options.CountOptions) (int64, error)
	Distinct(ctx context.Context, fieldName string, opts ...*options.DistinctOptions) ([]any, error)
	DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...*options.DistinctOptions) error
}

func NewFinder[T any](collection *mongo.Collection) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}}
}

var _ iFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection  *mongo.Collection
	filter      any
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn[T]
}

func (f *Finder[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Finder[T] {
	f.beforeHooks = append(f.beforeHooks, hooks...)
	return f
}

// RegisterAfterHooks is used to set the after hooks of the query
// If you register the hook for FindOne, the opContext.Docs will be nil
// If you register the hook for Find, the opContext.Doc will be nil
func (f *Finder[T]) RegisterAfterHooks(hooks ...afterHookFn[T]) *Finder[T] {
	f.afterHooks = append(f.afterHooks, hooks...)
	return f
}

// Filter is used to set the filter of the query
func (f *Finder[T]) Filter(filter any) *Finder[T] {
	f.filter = filter
	return f
}

func (f *Finder[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range f.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Finder[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *AfterOpContext[T], opType operation.OpType) error {
	err := callback.GetCallback().Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range f.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Finder[T]) FindOne(ctx context.Context, opts ...*options.FindOneOptions) (*T, error) {
	t := new(T)

	globalOpContext := operation.NewOpContext(f.collection, operation.WithDoc(t), operation.WithFilter(f.filter))
	err := f.preActionHandler(ctx, globalOpContext, NewOpContext(f.collection, f.filter), operation.OpTypeBeforeFind)
	if err != nil {
		return nil, err
	}

	err = f.collection.FindOne(ctx, f.filter, opts...).Decode(t)
	if err != nil {
		return nil, err
	}

	err = f.postActionHandler(ctx, globalOpContext, NewAfterOpContext[T](NewOpContext(f.collection, f.filter), WithDoc(t)), operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Find(ctx context.Context, opts ...*options.FindOptions) ([]*T, error) {
	t := make([]*T, 0)

	opContext := operation.NewOpContext(f.collection, operation.WithFilter(f.filter))
	err := f.preActionHandler(ctx, opContext, NewOpContext(f.collection, f.filter), operation.OpTypeBeforeFind)
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

	opContext.Doc = t
	err = f.postActionHandler(ctx, opContext, NewAfterOpContext[T](NewOpContext(f.collection, f.filter), WithDocs(t)), operation.OpTypeAfterFind)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (f *Finder[T]) Count(ctx context.Context, opts ...*options.CountOptions) (int64, error) {
	return f.collection.CountDocuments(ctx, f.filter, opts...)
}

func (f *Finder[T]) Distinct(ctx context.Context, fieldName string, opts ...*options.DistinctOptions) ([]any, error) {
	return f.collection.Distinct(ctx, fieldName, f.filter, opts...)
}

// DistinctWithParse is used to parse the result of Distinct
// result must be a pointer
func (f *Finder[T]) DistinctWithParse(ctx context.Context, fieldName string, result any, opts ...*options.DistinctOptions) error {
	docs, err := f.collection.Distinct(ctx, fieldName, f.filter, opts...)
	if err != nil {
		return err
	}

	valueType, valueBytes, err := bson.MarshalValue(docs)
	if err != nil {
		return err
	}
	rawValue := bson.RawValue{Type: valueType, Value: valueBytes}
	err = rawValue.Unmarshal(result)
	if err != nil {
		return err
	}
	return nil
}
