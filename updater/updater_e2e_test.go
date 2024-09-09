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
	"errors"
	"testing"

	"github.com/chenmingyong0423/go-mongox/hook/field"

	"github.com/chenmingyong0423/go-mongox/internal/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/builder/update"

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
	updater := NewUpdater[TestUser2](collection)
	assert.NotNil(t, updater)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx        context.Context
		filter     any
		updates    any
		opts       []*options.UpdateOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

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
			updates: User{Id: "1", Name: "Mingyong Chen", Age: 18},
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
			updates: map[string]any{"Id": "1", "Name": "Mingyong Chen", "Age": 18},
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, TestUser2{ID: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Id("1"))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.Id("2"),
			updates: update.NewBuilder().Set("name", "Mingyong Chen").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is bson.D",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, TestUser2{Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success when the updates is map[string]any",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, TestUser2{Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.Eq("name", "Mingyong Chen"),
			updates: map[string]any{
				"$set": map[string]any{
					"name": "chenmingyong",
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
				insertResult, err := collection.InsertOne(ctx, TestUser2{ID: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().InString("_id", "1", "2").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.Eq("_id", "2"),
			updates: update.Set("name", "chenmingyong"),
			opts: []*options.UpdateOptions{
				options.Update().SetUpsert(true),
			},
			want: &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "2"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update one success in field hook enabled",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, TestUser2{Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.NotNil(t, insertResult.InsertedID)
				callbacks := callback.GetCallback()
				callbacks.Register("beforeUpdate", "mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, "beforeUpdate", opts...)
				})
			},
			after: func(ctx context.Context, t *testing.T) {
				var user TestUser2
				err := collection.FindOne(ctx, query.Eq("name", "chenmingyong")).Decode(&user)
				require.NoError(t, err)
				require.Zero(t, user.CreatedAt)
				require.NotZero(t, user.UpdatedAt)

				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
				callbacks := callback.GetCallback()
				callbacks.Remove("beforeUpdate", "mongox:model")
			},
			ctx:     context.Background(),
			filter:  query.Eq("name", "Mingyong Chen"),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name:    "global before hook error",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "Mingyong Chen").Build(),
			want:    nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpdate,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, User{Id: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterUpdate,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, User{Id: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpdate,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("before hook error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterUpdate,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("after hook error")
						}
						return nil
					},
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: assert.NoError,
		},
		{
			name:    "before hook error",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "Mingyong Chen").Build(),
			want:    nil,
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opContext *BeforeOpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, User{Id: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opContext *AfterOpContext, opts ...any) error {
					return errors.New("after hook error")
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, User{Id: "1", Name: "Mingyong Chen", Age: 18})
				assert.NoError(t, err)
				assert.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opContext *BeforeOpContext, opts ...any) error {
					if opContext.Filter == nil || opContext.Updates == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opContext *AfterOpContext, opts ...any) error {
					if opContext.Filter == nil || opContext.Updates == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.NewBuilder().Set("name", "chenmingyong").Build(),
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			got, err := updater.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Updates(tc.updates).UpdateOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			updater.beforeHooks = nil
			updater.afterHooks = nil
		})
	}
}

func TestUpdater_e2e_UpdateMany(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[TestUser2](collection)
	assert.NotNil(t, updater)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx        context.Context
		filter     any
		updates    any
		opts       []*options.UpdateOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil updates",
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
			name:    "not contains any operator",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: bsonx.M("name", "Mingyong Chen"),
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "modified count is 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]TestUser2{
					{ID: "1", Name: "Mingyong Chen", Age: 18},
					{ID: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("_id", "1", "2"))
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				assert.NoError(t, err)
			},
			ctx:     context.Background(),
			filter:  query.In("_id", "3", "4"),
			updates: update.Set("name", "Mingyong Chen"),
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update many success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]TestUser2{
					{ID: "1", Name: "Mingyong Chen", Age: 18},
					{ID: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("_id", "1", "2"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  query.In("_id", "1", "2"),
			updates: update.Set("name", "hhh"),
			want:    &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "upserted count is 1",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]TestUser2{
					{ID: "1", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("_id", "1", "2"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx:     context.Background(),
			filter:  bsonx.Id("2"),
			updates: update.Set("name", "cmy"),
			opts: []*options.UpdateOptions{
				options.Update().SetUpsert(true),
			},
			want: &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1, UpsertedID: "2"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update many success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]TestUser2{
					{ID: "1", Name: "Mingyong Chen", Age: 18},
					{ID: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
				callbacks := callback.GetCallback()
				callbacks.Register("beforeUpdate", "mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, "beforeUpdate", opts...)
				})
			},
			after: func(ctx context.Context, t *testing.T) {
				users := make([]TestUser2, 2)
				cur, err := collection.Find(ctx, query.In("_id", "1", "2"))
				require.NoError(t, err)
				require.NoError(t, cur.All(ctx, &users))
				for _, user := range users {
					require.Zero(t, user.CreatedAt)
					require.NotZero(t, user.UpdatedAt)
				}

				deleteResult, err := collection.DeleteMany(ctx, query.In("_id", "1", "2"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				callbacks := callback.GetCallback()
				callbacks.Remove("beforeUpdate", "mongox:model")
			},
			ctx:     context.Background(),
			filter:  query.In("_id", "1", "2"),
			updates: update.Set("name", "hhh"),
			want:    &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name:    "global before hook error",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: nil,
			want:    nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpdate,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]User{
					{Id: "1", Name: "Mingyong Chen", Age: 18},
					{Id: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().InString("_id", "1", "2").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterUpdate,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.Set("name", "chenmingyong"),
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]User{
					{Id: "1", Name: "Mingyong Chen", Age: 18},
					{Id: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().InString("_id", "1", "2").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpdate,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("before hook error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterUpdate,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("after hook error")
						}
						return nil
					},
				},
			},
			ctx:     context.Background(),
			filter:  query.In("_id", "1", "2"),
			updates: update.Set("name", "chenmingyong"),
			want:    &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: assert.NoError,
		},
		{
			name:    "before hook error",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: nil,
			want:    nil,
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]User{
					{Id: "1", Name: "Mingyong Chen", Age: 18},
					{Id: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().InString("_id", "1", "2").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Updates == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					return errors.New("after hook error")
				},
			},
			ctx:     context.Background(),
			filter:  query.NewBuilder().Id("1").Build(),
			updates: update.Set("name", "chenmingyong"),
			want:    nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]User{
					{Id: "1", Name: "Mingyong Chen", Age: 18},
					{Id: "2", Name: "Mingyong Chen", Age: 18},
				}...))
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().InString("_id", "1", "2").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Updates == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Updates == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			ctx:     context.Background(),
			filter:  query.In("_id", "1", "2"),
			updates: update.Set("name", "chenmingyong"),
			want:    &mongo.UpdateResult{MatchedCount: 2, ModifiedCount: 2, UpsertedCount: 0, UpsertedID: nil},
			wantErr: assert.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			got, err := updater.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Updates(tc.updates).UpdateMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			updater.beforeHooks = nil
			updater.afterHooks = nil
		})
	}
}

func TestUpdater_e2e_Upsert(t *testing.T) {
	collection := getCollection(t)
	updater := NewUpdater[TestUser](collection)
	assert.NotNil(t, updater)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx        context.Context
		filter     any
		updates    any
		opts       []*options.UpdateOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		want    *mongo.UpdateResult
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "nil filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: bson.D{},
			want:    nil,
			wantErr: require.Error,
		},
		{
			name:    "invalid filter",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  6,
			updates: bson.D{},
			want:    nil,
			wantErr: require.Error,
		},
		{
			name:    "nil updates",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: nil,
			want:    nil,
			wantErr: require.Error,
		},
		{
			name:    "invalid updates",
			before:  func(ctx context.Context, t *testing.T) {},
			after:   func(ctx context.Context, t *testing.T) {},
			ctx:     context.Background(),
			filter:  bson.D{},
			updates: 6,
			want:    nil,
			wantErr: require.Error,
		},
		{
			name: "update successfully",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				var user TestUser
				err := collection.FindOne(ctx, query.Eq("name", "chenmingyong")).Decode(&user)
				require.NoError(t, err)
				require.Zero(t, user.CreatedAt)
				require.Zero(t, user.UpdatedAt)

				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.Eq("name", "Mingyong Chen"),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
				},
			},
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: require.NoError,
		},
		{
			name: "update successfully in field hook enabled",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
				callbacks := callback.GetCallback()
				callbacks.Register("beforeUpsert", "mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, "beforeUpsert", opts...)
				})
			},
			after: func(ctx context.Context, t *testing.T) {
				var user TestUser
				err := collection.FindOne(ctx, query.Eq("name", "chenmingyong")).Decode(&user)
				require.NoError(t, err)
				require.Zero(t, user.CreatedAt)
				require.NotZero(t, user.UpdatedAt)

				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)

				callbacks := callback.GetCallback()
				callbacks.Remove("beforeUpsert", "mongox:model")
			},
			ctx:    context.Background(),
			filter: query.Eq("name", "Mingyong Chen"),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
				},
			},
			want:    &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			wantErr: require.NoError,
		},
		{
			name:   "save successfully",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				var user TestUser
				err := collection.FindOne(ctx, query.Eq("name", "Mingyong Chen")).Decode(&user)
				require.NoError(t, err)
				require.Zero(t, user.CreatedAt)
				require.Zero(t, user.UpdatedAt)
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			opts:   nil,
			updates: bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "name", Value: "Mingyong Chen"},
				}},
			},
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1},
			wantErr: require.NoError,
		},
		{
			name: "save successfully in field hook enabled",
			before: func(ctx context.Context, t *testing.T) {
				callbacks := callback.GetCallback()
				callbacks.Register("beforeUpsert", "mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, "beforeUpsert", opts...)
				})
			},
			after: func(ctx context.Context, t *testing.T) {
				var user TestUser
				err := collection.FindOne(ctx, query.Eq("name", "Mingyong Chen")).Decode(&user)
				require.NoError(t, err)
				require.NotZero(t, user.CreatedAt)
				require.NotZero(t, user.UpdatedAt)
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
				callbacks := callback.GetCallback()
				callbacks.Remove("beforeUpsert", "mongox:model")
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			opts:   nil,
			updates: bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "name", Value: "Mingyong Chen"},
				}},
			},
			want:    &mongo.UpdateResult{MatchedCount: 0, ModifiedCount: 0, UpsertedCount: 1},
			wantErr: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			filter: query.NewBuilder().Id("1").Build(),
			want:   nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpsert,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
					"age":  18,
				},
			},
			want: nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterUpsert,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
					"age":  18,
				},
			},
			want: &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeUpsert,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("before hook error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterUpsert,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil || opCtx.Updates == nil {
							return errors.New("after hook error")
						}
						return nil
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			filter: query.NewBuilder().Id("1").Build(),
			want:   nil,
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Replacement == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
					"age":  18,
				},
			},
			want: nil,
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, &TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx:    context.Background(),
			filter: query.NewBuilder().Eq("name", "Mingyong Chen").Build(),
			updates: bson.M{
				"$set": bson.M{
					"name": "chenmingyong",
					"age":  18,
				},
			},
			want: &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1, UpsertedCount: 0, UpsertedID: nil},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Updates == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					if opCtx.Filter == nil || opCtx.Updates == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			got, err := updater.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Updates(tc.updates).Upsert(tc.ctx, tc.opts...)
			tc.wantErr(t, err)
			tc.after(tc.ctx, t)

			if err == nil {
				require.Equal(t, tc.want.MatchedCount, got.MatchedCount)
				require.Equal(t, tc.want.ModifiedCount, got.ModifiedCount)
				require.Equal(t, tc.want.UpsertedCount, got.UpsertedCount)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			updater.beforeHooks = nil
			updater.afterHooks = nil
		})
	}
}
