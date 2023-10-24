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

	"github.com/chenmingyong0423/go-mongox/builder"
	"github.com/chenmingyong0423/go-mongox/internal/types"
	"github.com/chenmingyong0423/go-mongox/pkg"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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
		filter  bson.D
		updates bson.D

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "failed to update one",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,
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
				deleteResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().Id("456").Build(),
			updates: builder.NewBsonBuilder().Set("name", "cmy").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, types.TestUser{Id: "123", Name: "cmy", Age: 24})
				assert.NoError(t, err)
				assert.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().Id("123").Build(),
			updates: builder.NewBsonBuilder().Set("name", "hhh").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).Updates(tc.updates).UpdateOne(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_e2e_UpdateOneWithOptions(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[any](collection)
	assert.NotNil(t, updater)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  bson.D
		updates bson.D
		opts    *options.UpdateOptions

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "failed to update one",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,
			opts:    nil,
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
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
				deleteResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().Id("456").Build(),
			updates: builder.NewBsonBuilder().Set("name", "cmy").Build(),
			opts:    options.Update().SetUpsert(true),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "456"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).Updates(tc.updates).UpdateOneWithOptions(tc.ctx, tc.opts)
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

		ctx     context.Context
		filter  bson.D
		updates bson.D

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "failed to update many",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, pkg.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
					{Id: "456", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				assert.NoError(t, err)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().InString("_id", []string{"789", "000"}...).Build(),
			updates: builder.NewBsonBuilder().Set("name", "hhh").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update many success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, pkg.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
					{Id: "456", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				assert.NoError(t, err)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().InString("_id", []string{"123", "456"}...).Build(),
			updates: builder.NewBsonBuilder().Set("name", "hhh").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).Updates(tc.updates).UpdateMany(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}

}

func TestUpdater_e2e_UpdateManyWithOptions(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[any](collection)
	assert.NotNil(t, updater)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  bson.D
		updates bson.D
		opts    *options.UpdateOptions

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "failed to update many",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,
			opts:    nil,
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "upserted count is 2",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, pkg.ToAnySlice([]types.TestUser{
					{Id: "123", Name: "cmy", Age: 24},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				assert.NoError(t, err)
			},
			ctx:     context.Background(),
			filter:  builder.NewBsonBuilder().Id("456").Build(),
			updates: builder.NewBsonBuilder().Set("name", "cmy").Build(),
			opts:    options.Update().SetUpsert(true),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "456"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			got, err := updater.Filter(tc.filter).Updates(tc.updates).UpdateManyWithOptions(tc.ctx, tc.opts)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
