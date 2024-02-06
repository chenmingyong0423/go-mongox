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

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=deleter.go -destination=../mock/deleter.mock.go -package=mocks
type iDeleter[T any] interface {
	DeleteOne(ctx context.Context, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

func NewDeleter[T any](collection *mongo.Collection) *Deleter[T] {
	return &Deleter[T]{collection: collection, filter: nil}
}

type Deleter[T any] struct {
	collection *mongo.Collection
	filter     any
}

// Filter is used to set the filter of the query
func (d *Deleter[T]) Filter(filter any) *Deleter[T] {
	d.filter = filter
	return d
}

func (d *Deleter[T]) DeleteOne(ctx context.Context, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	opContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter))
	err := callback.GetCallback().Execute(ctx, opContext, operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteOne(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	err = callback.GetCallback().Execute(ctx, opContext, operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Deleter[T]) DeleteMany(ctx context.Context, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	opContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter))
	err := callback.GetCallback().Execute(ctx, opContext, operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteMany(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	err = callback.GetCallback().Execute(ctx, opContext, operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}
