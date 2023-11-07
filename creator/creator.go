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

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=creator.go -destination=../mock/creator.mock.go -package=mocks
type iCreator[T any] interface {
	InsertOne(ctx context.Context, docs T) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, docs []T) (*mongo.InsertManyResult, error)
}

type Creator[T any] struct {
	collection        *mongo.Collection
	insertManyOptions []*options.InsertManyOptions
	insertOneOptions  []*options.InsertOneOptions
}

func NewCreator[T any](collection *mongo.Collection) *Creator[T] {
	return &Creator[T]{
		collection: collection,
	}
}

func (c *Creator[T]) InsertOne(ctx context.Context, doc T) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(ctx, doc, c.insertOneOptions...)
}

func (c *Creator[T]) InsertOneOptions(opts ...*options.InsertOneOptions) *Creator[T] {
	c.insertOneOptions = opts
	return c
}

func (c *Creator[T]) InsertMany(ctx context.Context, docs []T) (*mongo.InsertManyResult, error) {
	return c.collection.InsertMany(ctx, utils.ToAnySlice(docs...), c.insertManyOptions...)
}

func (c *Creator[T]) InsertManyOptions(opts ...*options.InsertManyOptions) *Creator[T] {
	c.insertManyOptions = opts
	return c
}
