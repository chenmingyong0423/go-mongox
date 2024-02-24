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

package deleter

import (
	"context"
	"fmt"
	"testing"

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/bsonx"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type testTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

func newCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	require.NoError(t, err)
	require.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := client.Database("db-test").Collection("test_user")
	return collection
}

func TestDeleter_e2e_New(t *testing.T) {
	result := NewDeleter[any](newCollection(t))
	require.NotNil(t, result, "Expected non-nil Deleter")
}

func TestDeleter_e2e_DeleteOne(t *testing.T) {
	collection := newCollection(t)
	deleter := NewDeleter[testTempUser](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     bson.D
		opts       []*options.DeleteOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name:      "error: nil filter",
			before:    func(_ context.Context, _ *testing.T) {},
			after:     func(_ context.Context, _ *testing.T) {},
			filter:    nil,
			ctx:       context.Background(),
			opts:      []*options.DeleteOptions{options.Delete().SetComment("test")},
			want:      nil,
			wantError: require.Error,
		},
		{
			name: "deleted count: 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("2").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantError: require.NoError,
		},
		{
			name: "delete success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("before hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name:   "global after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("after hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					return fmt.Errorf("before hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name:   "after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					return fmt.Errorf("after hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertOne(ctx, testTempUser{Id: "1", Name: "Mingyong Chen"})
				require.NoError(t, err)
				require.Equal(t, "1", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).DeleteOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			deleter.beforeHooks = nil
			deleter.afterHooks = nil
		})
	}
}

func TestDeleter_e2e_DeleteMany(t *testing.T) {
	collection := newCollection(t)
	deleter := NewDeleter[testTempUser](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     any
		opts       []*options.DeleteOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name:      "error: nil filter",
			before:    func(_ context.Context, _ *testing.T) {},
			after:     func(_ context.Context, _ *testing.T) {},
			filter:    nil,
			ctx:       context.Background(),
			opts:      []*options.DeleteOptions{options.Delete().SetComment("test")},
			want:      nil,
			wantError: require.Error,
		},
		{
			name: "deleted count: 0",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantError: require.NoError,
		},
		{
			name: "delete success",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: bsonx.M("name", "Mingyong Chen"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("before hook error")
					},
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return fmt.Errorf("after hook error")
					},
				},
			},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want:   nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterDelete,
					name:   "before delete hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return fmt.Errorf("filter is nil")
						}
						return nil
					},
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					return fmt.Errorf("before hook error")
				},
			},
			want: nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "before hook error", err.Error())
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					return fmt.Errorf("after hook error")
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want:   nil,
			wantError: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, "after hook error", err.Error())
			},
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]testTempUser{
					{Id: "1", Name: "Mingyong Chen"},
					{Id: "2", Name: "Mingyong Chen"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"1", "2"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *BeforeOpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []afterHookFn{
				func(ctx context.Context, opCtx *AfterOpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).DeleteMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			deleter.beforeHooks = nil
			deleter.afterHooks = nil
		})
	}
}
