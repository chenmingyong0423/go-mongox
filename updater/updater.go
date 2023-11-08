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

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/builder/update"

	"github.com/chenmingyong0423/go-mongox/converter"
	"github.com/chenmingyong0423/go-mongox/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
type iUpdater[T any] interface {
	UpdateOne(ctx context.Context) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context) (*mongo.UpdateResult, error)
}

func NewUpdater[T any](collection *mongo.Collection) *Updater[T] {
	return &Updater[T]{collection: collection, filter: nil}
}

type Updater[T any] struct {
	collection *mongo.Collection
	filter     any
	updates    any
	opts       []*options.UpdateOptions
}

// Filter is used to set the filter of the query
func (u *Updater[T]) Filter(filter any) *Updater[T] {
	u.filter = filter
	return u
}

// FilterKeyValue is used to set the filter of the query
func (u *Updater[T]) FilterKeyValue(bsonElements ...types.KeyValue[any]) *Updater[T] {
	u.filter = query.BsonBuilder().Add(bsonElements...).Build()
	return u
}

// Updates is used to set the updates of the update
// if the type of updates is in [map[string]any, map[string]any pointer, struct, struct pointer], it will be converted to bson.D
// or it will be set to updates directly
func (u *Updater[T]) Updates(updates any) *Updater[T] {
	d := converter.ToSetBson(updates)
	if d != nil {
		u.updates = d
	} else {
		u.updates = updates
	}
	return u
}

// UpdatesKeyValue is used to set the updates of the update
func (u *Updater[T]) UpdatesKeyValue(bsonElements ...types.KeyValue[any]) *Updater[T] {
	if len(bsonElements) != 0 {
		u.updates = bson.D{bson.E{Key: types.Set, Value: update.BsonBuilder().Add(bsonElements...).Build()}}
	} else {
		u.updates = nil
	}
	return u
}

func (u *Updater[T]) UpdateOne(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.collection.UpdateOne(ctx, u.filter, u.updates, u.opts...)
}

func (u *Updater[T]) UpdateOptions(opts ...*options.UpdateOptions) *Updater[T] {
	u.opts = opts
	return u
}

func (u *Updater[T]) UpdateMany(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.collection.UpdateMany(ctx, u.filter, u.updates, u.opts...)
}
