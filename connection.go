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
	"reflect"

	mongoxErr "github.com/chenmingyong0423/go-mongox/error"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
	err    error
}

type Config struct {
	ClientOpts []*options.ClientOptions
	DbOpts     []*options.DatabaseOptions
}

func Open(ctx context.Context, db string, cfg *Config) (mc MongoClient, err error) {
	mc = MongoClient{}
	mc.client, err = mongo.Connect(ctx, cfg.ClientOpts...)
	if err != nil {
		return
	}
	mc.db = mc.client.Database(db, cfg.DbOpts...)
	return
}

func (mc *MongoClient) Disconnect(ctx context.Context) error {
	if err := mc.client.Disconnect(ctx); err != nil {
		return err
	}
	mc.client = nil
	return nil
}

func (mc *MongoClient) Coll(model any, opts ...*options.CollectionOptions) *Collection {
	collName := ""
	if model == nil {
		mc.err = mongoxErr.ErrModelIsNil
	}
	valType := reflect.TypeOf(model)
	switch valType.Kind() {
	case reflect.Struct:
		if col, ok := model.(CollectionName); ok {
			collName = col.CollectionName()
		} else {
			collName = valType.Name()
		}
	case reflect.Ptr:
		return mc.Coll(valType.Elem())
	default:
		mc.err = mongoxErr.ErrNotStructType
	}
	return newCollection(mc.db.Collection(collName, opts...), mc.err)
}

func (mc *MongoClient) CollectionByCollName(collName string, opts ...*options.CollectionOptions) *Collection {
	return newCollection(mc.db.Collection(collName, opts...), nil)
}
