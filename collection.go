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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ID = "_id"
)

type CollectionName interface {
	CollectionName() string
}

type Collection struct {
	coll *mongo.Collection
	err  error
}

func newCollection(coll *mongo.Collection, err error) *Collection {
	return &Collection{
		coll: coll,
		err:  err,
	}
}

func (c *Collection) FindById(ctx context.Context, id any, expectPtr any, opts ...*options.FindOneOptions) error {
	if c.err != nil {
		return c.err
	}
	return c.coll.FindOne(ctx, bson.M{ID: id}, opts...).Decode(expectPtr)
}
