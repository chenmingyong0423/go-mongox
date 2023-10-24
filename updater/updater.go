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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=updater.go -destination=../mock/updater.mock.go -package=mocks
type iUpdater interface {
	UpdateOne(ctx context.Context) (*mongo.UpdateResult, error)
	UpdateOneWithOptions(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context) (*mongo.UpdateResult, error)
	UpdateManyWithOptions(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

func NewUpdater(collection *mongo.Collection) *Updater {
	return &Updater{collection: collection, filter: nil}
}

type Updater struct {
	collection *mongo.Collection
	filter     bson.D
	updates    bson.D
	opts       []*options.UpdateOptions
}

func (u *Updater) Filter(filter bson.D) *Updater {
	u.filter = filter
	return u
}

func (u *Updater) Updates(updates bson.D) *Updater {
	u.updates = updates
	return u
}

func (u *Updater) UpdateOne(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.collection.UpdateOne(ctx, u.filter, u.updates, u.opts...)
}

func (u *Updater) UpdateOneWithOptions(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	u.opts = opts
	return u.UpdateOne(ctx)
}

func (u *Updater) UpdateMany(ctx context.Context) (*mongo.UpdateResult, error) {
	return u.collection.UpdateMany(ctx, u.filter, u.updates, u.opts...)
}

func (u *Updater) UpdateManyWithOptions(ctx context.Context, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	u.opts = opts
	return u.UpdateMany(ctx)
}
