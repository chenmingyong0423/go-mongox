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

//go:build e2e

package updater

import (
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/bsonx"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/builder/update"

	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	return client.Database("db-test").Collection("test_user")
}

func TestUpdater_e2e_New(t *testing.T) {
	updater := NewUpdater[any](getCollection(t))
	assert.NotNil(t, updater)
}

func TestUpdater_e2e_UpdateOne(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[any](collection)
	assert.NotNil(t, updater)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  any
		updates any
		opts    []*options.UpdateOptions

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: bson.D{},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "invalid filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  6,
			updates: bson.D{},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "nil updates,failed to update one",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: nil,
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "invalid updates",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: 6,
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "got error when updates(struct) not contains any operator",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: types.TestUser{Id: "123", Name: "cmy", Age: 24},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "got error when updates(map) not contains any operator",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: map[string]any{"Id": "123", "Name": "cmy", "Age": 24},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.BsonBuilder().Id("456").Build(),
			updates: update.BsonBuilder().Set(bsonx.D(bsonx.E("name", "cmy"))).Build(),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is bson.D",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.BsonBuilder().Id("123").Build(),
			updates: update.BsonBuilder().Set(bsonx.D(bsonx.E("name", "hhh"))).Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is map[string]any",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.BsonBuilder().Id("123").Build(),
			updates: map[string]any{
				"$set": map[string]any{
					"name": "hhh",
				},
			},
			want: &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "upserted count is 1",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.BsonBuilder().Id("456").Build(),
			updates: update.BsonBuilder().Set(bsonx.D(bsonx.E("name", "cmy"))).Build(),
			opts: []*options.UpdateOptions{
				options.Update().SetUpsert(true),
			},
			want: &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "456"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).Updates(tc.updates).Options(tc.opts...).UpdateOne(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_e2e_UpdateMany(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[any](collection)
	assert.NotNil(t, updater)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx      context.Context
		filter   any
		operator string
		updates  any
		opts     []*options.UpdateOptions

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:     "nil updates",
			before:   func(ctx context.Context, t *testing.T) {},
			after:    func(ctx context.Context, t *testing.T) {},
			ctx:      context.Background(),
			filter:   nil,
			operator: types.Set,
			updates:  nil,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:     "not contains any operator",
			before:   func(ctx context.Context, t *testing.T) {},
			after:    func(ctx context.Context, t *testing.T) {},
			ctx:      context.Background(),
			filter:   nil,
			operator: "",
			updates:  bsonx.M("name", "cmy"),
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
					{Id: "456", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				assert.NoError(t, err)
			},
			ctx:      context.Background(),
			operator: types.Set,
			filter:   query.BsonBuilder().InString("_id", "789", "000").Build(),
			updates:  bsonx.M("name", "cmy"),
			want:     &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update many success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
					{Id: "456", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			operator: types.Set,
			ctx:      context.Background(),
			filter:   query.BsonBuilder().InString("_id", "123", "456").Build(),
			updates:  bsonx.M("name", "hhh"),
			want:     &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "upserted count is 1",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			operator: types.Set,
			ctx:      context.Background(),
			filter:   bsonx.Id("456"),
			updates:  bsonx.M("name", "cmy"),
			opts: []*options.UpdateOptions{
				options.Update().SetUpsert(true),
			},
			want: &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "456"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).UpdatesWithOperator(tc.operator, tc.updates).Options(tc.opts...).UpdateMany(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_e2e_UpdatesWithOperator(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[any](collection)
	assert.NotNil(t, updater)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx      context.Context
		operator string
		filter   any
		updates  any
		opts     []*options.UpdateOptions

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: bson.D{},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "invalid filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  6,
			updates: bson.D{},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:     "nil updates,failed to update one",
			before:   func(ctx context.Context, t *testing.T) {},
			after:    func(ctx context.Context, t *testing.T) {},
			ctx:      context.Background(),
			operator: types.Set,
			filter:   bson.D{},
			updates:  nil,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:     "invalid updates",
			before:   func(ctx context.Context, t *testing.T) {},
			after:    func(ctx context.Context, t *testing.T) {},
			ctx:      context.Background(),
			filter:   bson.D{},
			operator: types.Set,
			updates:  6,
			want:     nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name:    "got error when operator not contains any operator",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: map[string]any{"Id": "123", "Name": "cmy", "Age": 24},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			operator: types.Set,
			ctx:      context.Background(),
			filter:   query.BsonBuilder().Id("456").Build(),
			updates:  update.BsonBuilder().AddToSet(bsonx.D(bsonx.E("name", "cmy"))).Build(),
			want:     &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is bson.D",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			operator: types.Set,
			ctx:      context.Background(),
			filter:   query.BsonBuilder().Id("123").Build(),
			updates:  update.BsonBuilder().Add(types.KV("name", "hhh")).Build(),
			want:     &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is map[string]any",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:      context.Background(),
			filter:   query.BsonBuilder().Id("123").Build(),
			operator: types.Set,
			updates: map[string]any{
				"name": "hhh",
			},
			want: &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "upserted count is 1",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:      context.Background(),
			operator: types.Set,
			filter:   query.BsonBuilder().Id("456").Build(),
			updates:  update.BsonBuilder().AddToSet(bsonx.D(bsonx.E("name", "cmy"))).Build(),
			opts: []*options.UpdateOptions{
				options.Update().SetUpsert(true),
			},
			want: &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "456"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).UpdatesWithOperator(tc.operator, tc.updates).Options(tc.opts...).UpdateOne(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
