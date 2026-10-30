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

package deleter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/internal/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"
	xdeleter "github.com/chenmingyong0423/go-mongox/v2/deleter"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type testTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type testSoftDeleteUser struct {
	testTempUser `bson:",inline"`
	DeletedAt    time.Time `bson:"deleted_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func newCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
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
	result := xdeleter.NewDeleter[testTempUser](newCollection(t), callback.InitializeCallbacks(), field.ParseFields(testTempUser{}))
	require.NotNil(t, result, "Expected non-nil Deleter")
}

func TestDeleter_e2e_DeleteOne(t *testing.T) {
	t.Run("physical delete", testDeleterDeleteOne)
	t.Run("soft delete", testSoftDeleterDeleteOne)
}

func testDeleterDeleteOne(t *testing.T) {
	t.Helper()
	collection := newCollection(t)
	deleter := xdeleter.NewDeleter[testTempUser](collection, callback.InitializeCallbacks(), field.ParseFields(testTempUser{}))

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
		opts       []options.Lister[options.DeleteOneOptions]
		updateOpts []options.Lister[options.UpdateOneOptions]
		globalHook []globalHook
		beforeHook []xdeleter.BeforeHookFn
		afterHook  []xdeleter.AfterHookFn

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
			opts:      []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
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
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("2").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
				Acknowledged: true,
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
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter:     query.NewBuilder().Id("1").Build(),
			ctx:        context.Background(),
			opts:       []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test").SetHint("_id_")},
			updateOpts: []options.Lister[options.UpdateOneOptions]{options.UpdateOne().SetHint("unused_index")},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
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
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
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
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
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
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			beforeHook: []xdeleter.BeforeHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
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
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			afterHook: []xdeleter.AfterHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
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
				deleteResult, err := collection.DeleteOne(ctx, query.NewBuilder().Id("1").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.NewBuilder().Id("1").Build(),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("test")},
			beforeHook: []xdeleter.BeforeHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []xdeleter.AfterHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				deleter.DBCallbacks.Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).UpdateOneOptions(tc.updateOpts...).DeleteOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				deleter.DBCallbacks.Remove(hook.opType, hook.name)
			}
			deleter.BeforeHooks = nil
			deleter.AfterHooks = nil
		})
	}
}

func testSoftDeleterDeleteOne(t *testing.T) {
	t.Helper()
	collection := newCollection(t)
	past := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	deleter := xdeleter.NewDeleter[testSoftDeleteUser](collection, callback.InitializeCallbacks(), field.ParseFields(testSoftDeleteUser{}))

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     bson.D
		opts       []options.Lister[options.DeleteOneOptions]
		updateOpts []options.Lister[options.UpdateOneOptions]
		beforeHook []xdeleter.BeforeHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "success",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertMany(ctx, []any{
					bson.M{"_id": "soft", "updated_at": past},
					bson.M{"_id": "deleted", "deleted_at": past, "updated_at": past},
				})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				var stored testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("soft")).Decode(&stored))
				require.False(t, stored.DeletedAt.IsZero())
				require.Equal(t, past, stored.UpdatedAt)
				result, err := deleter.DeleteOne(ctx)
				require.NoError(t, err)
				require.Zero(t, result.DeletedCount)
				var repeated testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("soft")).Decode(&repeated))
				require.Equal(t, stored.DeletedAt, repeated.DeletedAt)
				var deleted testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("deleted")).Decode(&deleted))
				require.Equal(t, past, deleted.DeletedAt)
				result, err = collection.DeleteMany(ctx, query.In("_id", "soft", "deleted"))
				require.NoError(t, err)
				require.Equal(t, int64(2), result.DeletedCount)
			},
			filter:     query.NewBuilder().Id("soft").Build(),
			opts:       []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetHint("unused_index")},
			updateOpts: []options.Lister[options.UpdateOneOptions]{options.UpdateOne().SetHint("_id_")},
			ctx:        context.Background(),
			want:       &mongo.DeleteResult{DeletedCount: 1, Acknowledged: true},
			wantError:  require.NoError,
		},
		{
			name: "already deleted document",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertOne(ctx, bson.M{"_id": "deleted", "deleted_at": past})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				var stored testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("deleted")).Decode(&stored))
				require.Equal(t, past, stored.DeletedAt)
				result, err := collection.DeleteOne(ctx, bsonx.Id("deleted"))
				require.NoError(t, err)
				require.Equal(t, int64(1), result.DeletedCount)
			},
			filter:    query.NewBuilder().Id("deleted").Build(),
			ctx:       context.Background(),
			want:      &mongo.DeleteResult{Acknowledged: true},
			wantError: require.NoError,
		},
		{
			name: "nil updates",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertOne(ctx, testTempUser{Id: "soft"})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				var stored testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("soft")).Decode(&stored))
				require.True(t, stored.DeletedAt.IsZero())
				result, err := collection.DeleteOne(ctx, bsonx.Id("soft"))
				require.NoError(t, err)
				require.Equal(t, int64(1), result.DeletedCount)
			},
			filter: query.NewBuilder().Id("soft").Build(),
			beforeHook: []xdeleter.BeforeHookFn{
				func(_ context.Context, opCtx *xdeleter.OpContext, _ ...any) error {
					opCtx.Updates = nil
					return nil
				},
			},
			ctx:  context.Background(),
			want: nil,
			wantError: func(t require.TestingT, err error, _ ...interface{}) {
				require.ErrorIs(t, err, mongo.ErrNilDocument)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).Filter(tc.filter).UpdateOneOptions(tc.updateOpts...).DeleteOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			deleter.BeforeHooks = nil
		})
	}
}

func TestDeleter_e2e_DeleteMany(t *testing.T) {
	t.Run("physical delete", testDeleterDeleteMany)
	t.Run("soft delete", testSoftDeleterDeleteMany)
}

func testDeleterDeleteMany(t *testing.T) {
	t.Helper()
	collection := newCollection(t)
	deleter := xdeleter.NewDeleter[testTempUser](collection, callback.InitializeCallbacks(), field.ParseFields(testTempUser{}))

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
		opts       []options.Lister[options.DeleteManyOptions]
		updateOpts []options.Lister[options.UpdateManyOptions]
		globalHook []globalHook
		beforeHook []xdeleter.BeforeHookFn
		afterHook  []xdeleter.AfterHookFn

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
			opts:      []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
				Acknowledged: true,
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter:     bsonx.M("name", "Mingyong Chen"),
			ctx:        context.Background(),
			opts:       []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test").SetHint("_id_")},
			updateOpts: []options.Lister[options.UpdateManyOptions]{options.UpdateMany().SetHint("unused_index")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
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
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
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
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: bsonx.Id("789"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			beforeHook: []xdeleter.BeforeHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			afterHook: []xdeleter.AfterHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
					return fmt.Errorf("after hook error")
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
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
				deleteResult, err := collection.DeleteMany(ctx, query.NewBuilder().Eq("name", "Mingyong Chen").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			beforeHook: []xdeleter.BeforeHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			afterHook: []xdeleter.AfterHookFn{
				func(ctx context.Context, opCtx *xdeleter.OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return fmt.Errorf("filter is nil")
					}
					return nil
				},
			},
			filter: query.In("_id", "1", "2"),
			ctx:    context.Background(),
			opts:   []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 2,
				Acknowledged: true,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				deleter.DBCallbacks.Register(hook.opType, hook.name, hook.fn)
			}
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).UpdateManyOptions(tc.updateOpts...).DeleteMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			for _, hook := range tc.globalHook {
				deleter.DBCallbacks.Remove(hook.opType, hook.name)
			}
			deleter.BeforeHooks = nil
			deleter.AfterHooks = nil
		})
	}
}

func testSoftDeleterDeleteMany(t *testing.T) {
	t.Helper()
	collection := newCollection(t)
	past := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	deleter := xdeleter.NewDeleter[testSoftDeleteUser](collection, callback.InitializeCallbacks(), field.ParseFields(testSoftDeleteUser{}))

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter     any
		opts       []options.Lister[options.DeleteManyOptions]
		updateOpts []options.Lister[options.UpdateManyOptions]
		beforeHook []xdeleter.BeforeHookFn

		ctx       context.Context
		want      *mongo.DeleteResult
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "success",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertMany(ctx, []any{
					bson.M{"_id": "missing", "updated_at": past},
					bson.M{"_id": "null", "deleted_at": nil, "updated_at": past},
					bson.M{"_id": "zero", "deleted_at": time.Time{}, "updated_at": past},
					bson.M{"_id": "deleted", "deleted_at": past, "updated_at": past},
				})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				cursor, err := collection.Find(ctx, bson.M{})
				require.NoError(t, err)
				var stored []testSoftDeleteUser
				require.NoError(t, cursor.All(ctx, &stored))
				require.Len(t, stored, 4)
				for _, doc := range stored {
					require.False(t, doc.DeletedAt.IsZero())
					require.Equal(t, past, doc.UpdatedAt)
					if doc.Id == "deleted" {
						require.Equal(t, past, doc.DeletedAt)
					}
				}
				result, err := deleter.DeleteMany(ctx)
				require.NoError(t, err)
				require.Zero(t, result.DeletedCount)
				result, err = collection.DeleteMany(ctx, query.In("_id", "missing", "null", "zero", "deleted"))
				require.NoError(t, err)
				require.Equal(t, int64(4), result.DeletedCount)
			},
			filter:     bson.M{},
			opts:       []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetHint("unused_index")},
			updateOpts: []options.Lister[options.UpdateManyOptions]{options.UpdateMany().SetHint("_id_")},
			ctx:        context.Background(),
			want:       &mongo.DeleteResult{DeletedCount: 3, Acknowledged: true},
			wantError:  require.NoError,
		},
		{
			name: "already deleted documents",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertOne(ctx, bson.M{"_id": "deleted", "deleted_at": past})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				var stored testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("deleted")).Decode(&stored))
				require.Equal(t, past, stored.DeletedAt)
				result, err := collection.DeleteMany(ctx, bsonx.Id("deleted"))
				require.NoError(t, err)
				require.Equal(t, int64(1), result.DeletedCount)
			},
			filter:    bsonx.Id("deleted"),
			ctx:       context.Background(),
			want:      &mongo.DeleteResult{Acknowledged: true},
			wantError: require.NoError,
		},
		{
			name: "nil updates",
			before: func(ctx context.Context, t *testing.T) {
				_, err := collection.InsertOne(ctx, testTempUser{Id: "soft"})
				require.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				var stored testSoftDeleteUser
				require.NoError(t, collection.FindOne(ctx, bsonx.Id("soft")).Decode(&stored))
				require.True(t, stored.DeletedAt.IsZero())
				result, err := collection.DeleteMany(ctx, bsonx.Id("soft"))
				require.NoError(t, err)
				require.Equal(t, int64(1), result.DeletedCount)
			},
			filter: bsonx.Id("soft"),
			beforeHook: []xdeleter.BeforeHookFn{
				func(_ context.Context, opCtx *xdeleter.OpContext, _ ...any) error {
					opCtx.Updates = nil
					return nil
				},
			},
			ctx:  context.Background(),
			want: nil,
			wantError: func(t require.TestingT, err error, _ ...interface{}) {
				require.ErrorIs(t, err, mongo.ErrNilDocument)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			result, err := deleter.RegisterBeforeHooks(tc.beforeHook...).Filter(tc.filter).UpdateManyOptions(tc.updateOpts...).DeleteMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
			deleter.BeforeHooks = nil
		})
	}
}
