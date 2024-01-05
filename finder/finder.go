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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
type iFinder[T any] interface {
	FindOne(ctx context.Context, opts ...*options.FindOneOptions) (*T, error)
	Find(ctx context.Context, opts ...*options.FindOptions) ([]*T, error)
	Count(ctx context.Context, opts ...*options.CountOptions) (int64, error)
}

func NewFinder[T any](collection *mongo.Collection) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}}
}

var _ iFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection *mongo.Collection
	filter     any
}

// Filter is used to set the filter of the query
func (f *Finder[T]) Filter(filter any) *Finder[T] {
	f.filter = filter
	return f
}

func (f *Finder[T]) FindOne(ctx context.Context, opts ...*options.FindOneOptions) (*T, error) {
	t := new(T)
	err := f.collection.FindOne(ctx, f.filter, opts...).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (f *Finder[T]) Find(ctx context.Context, opts ...*options.FindOptions) ([]*T, error) {
	t := make([]*T, 0)
	cursor, err := f.collection.Find(ctx, f.filter, opts...)
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

func (f *Finder[T]) Count(ctx context.Context, opts ...*options.CountOptions) (int64, error) {
	cnt, err := f.collection.CountDocuments(ctx, f.filter, opts...)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
