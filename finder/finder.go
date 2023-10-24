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
	FindOne(ctx context.Context) (*T, error)
	FindOneWithOptions(ctx context.Context, opts []*options.FindOneOptions) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	FindAllWithOptions(ctx context.Context, opts []*options.FindOptions) ([]*T, error)
}

func NewFinder[T any](collection *mongo.Collection) *Finder[T] {
	return &Finder[T]{collection: collection, filter: bson.D{}}
}

var _ iFinder[any] = (*Finder[any])(nil)

type Finder[T any] struct {
	collection  *mongo.Collection
	findOneOpts []*options.FindOneOptions
	findOpts    []*options.FindOptions
	filter      bson.D
}

func (f *Finder[T]) Filter(d bson.D) *Finder[T] {
	f.filter = d
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

func (f *Finder[T]) FindOneWithOptions(ctx context.Context, opts []*options.FindOneOptions) (*T, error) {
	f.findOneOpts = opts
	return f.FindOne(ctx)
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

func (f *Finder[T]) FindAllWithOptions(ctx context.Context, opts []*options.FindOptions) ([]*T, error) {
	f.findOpts = opts
	return f.FindAll(ctx)
}
