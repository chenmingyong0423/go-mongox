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

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
type iFinder[T any] interface {
	FindOne(ctx context.Context) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	Count(ctx context.Context) (int64, error)
}

func NewFinder[T any](collection *mongo.Collection) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}}
}

var _ iFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection   *mongo.Collection
	findOneOpts  []*options.FindOneOptions
	findOpts     []*options.FindOptions
	countOptions []*options.CountOptions
	filter       any
}

// Filter is used to set the filter of the query
func (f *Finder[T]) Filter(filter any) *Finder[T] {
	f.filter = filter
	return f
}

// FilterKeyValue is used to set the filter of the query
func (f *Finder[T]) FilterKeyValue(bsonElements ...types.KeyValue[any]) *Finder[T] {
	if bsonElements == nil {
		f.filter = nil
	} else {
		f.filter = query.BsonBuilder().Add(bsonElements...).Build()
	}
	return f
}

func (f *Finder[T]) FindOne(ctx context.Context) (*T, error) {
	t := new(T)
	err := f.collection.FindOne(ctx, f.filter, f.findOneOpts...).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (f *Finder[T]) OneOptions(opts ...*options.FindOneOptions) *Finder[T] {
	f.findOneOpts = opts
	return f
}

func (f *Finder[T]) FindAll(ctx context.Context) ([]*T, error) {
	t := make([]*T, 0)
	cursor, err := f.collection.Find(ctx, f.filter, f.findOpts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (f *Finder[T]) AllOptions(opts ...*options.FindOptions) *Finder[T] {
	f.findOpts = opts
	return f
}

func (f *Finder[T]) CountOptions(opts ...*options.CountOptions) *Finder[T] {
	f.countOptions = opts
	return f
}

func (f *Finder[T]) Count(ctx context.Context) (int64, error) {
	cnt, err := f.collection.CountDocuments(ctx, f.filter, f.countOptions...)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
