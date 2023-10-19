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

package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCollection[T any](collection *mongo.Collection) *Collection[T] {
	return &Collection[T]{collection: collection}
}

type Collection[T any] struct {
	collection *mongo.Collection
}

func (c *Collection[T]) FindById(ctx context.Context, id any, opts ...*options.FindOneOptions) (t *T, err error) {
	err = c.collection.FindOne(ctx, NewBsonBuilder().Id(id).Build(), opts...).Decode(t)
	return
}

func (c *Collection[T]) Find(ctx context.Context, filter any, opts ...*options.FindOptions) (t []*T, err error) {
	bsonFilter := toBson(filter)
	cursor, err := c.collection.Find(ctx, bsonFilter, opts...)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &t)
	return
}

func (c *Collection[T]) FindOneAndDelete(ctx context.Context, filter any, opts ...*options.FindOneAndDeleteOptions) (t *T, err error) {
	// 是否判断 bsonFilter 为 bson.D{}，待探究
	bsonFilter := toBson(filter)
	err = c.collection.FindOneAndDelete(ctx, bsonFilter, opts...).Decode(t)
	return
}

func (c *Collection[T]) FindOneAndUpdate(ctx context.Context, filter any, updates any, opts ...*options.FindOneAndUpdateOptions) (t *T, err error) {
	// 是否判断 bsonFilter 为 bson.D{}，待探究
	bsonFilter := toBson(filter)
	newUpdates := toSetBson(updates)
	err = c.collection.FindOneAndUpdate(ctx, bsonFilter, newUpdates, opts...).Decode(t)
	return
}

func (c *Collection[T]) FindOneAndReplace(ctx context.Context, filter any, updates any, opts ...*options.FindOneAndReplaceOptions) (t *T, err error) {
	// 是否判断 bsonFilter 为 bson.D{}，待探究
	bsonFilter := toBson(filter)
	err = c.collection.FindOneAndReplace(ctx, bsonFilter, updates, opts...).Decode(t)
	return
}
