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

package finder

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/bsonx"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

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
	require.NoError(t, err)
	require.NoError(t, client.Ping(context.Background(), readpref.Primary()))
	return client.Database("db-test").Collection("test_user")
}

func TestFinder_e2e_New(t *testing.T) {
	collection := getCollection(t)

	result := NewFinder[types.TestUser](collection)
	require.NotNil(t, result, "Expected non-nil Finder")
	require.Equal(t, collection, result.collection, "Expected finder field to be initialized correctly")
}

func TestFinder_e2e_FindOne(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

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
		opts       []*options.FindOneOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn[types.TestUser]

		ctx     context.Context
		want    *types.TestUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "chenmingyong",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "burt"),
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "find by name",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "chenmingyong",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "chenmingyong"),
			want: &types.TestUser{
				Name: "chenmingyong",
				Age:  24,
			},
		},
		{
			name: "ignore age field",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "chenmingyong",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "chenmingyong"),
			opts: []*options.FindOneOptions{
				{
					Projection: bsonx.M("age", 0),
				},
			},
			want: &types.TestUser{
				Name: "chenmingyong",
			},
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global before hook error")
					},
				},
			},
			wantErr: errors.New("global before hook error"),
		},
		{
			name: "global after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "Mingyong Chen",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("global after hook error")
					},
				},
			},
			wantErr: errors.New("global after hook error"),
		},
		{
			name: "global before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
							return errors.New("filter error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						user := opCtx.Doc.(*types.TestUser)
						if user.Name != "Mingyong Chen" || user.Age != 18 {
							return errors.New("result error")
						}
						return nil
					},
				},
			},
			want: &types.TestUser{
				Name: "Mingyong Chen",
				Age:  18,
			},
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: errors.New("before hook error"),
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			afterHook: []afterHookFn[types.TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[types.TestUser], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantErr: errors.New("after hook error"),
		},
		{
			name: "before and after hook",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "Mingyong Chen",
					Age:  18,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter.(bson.D)[0].Key != "name" || opCtx.Filter.(bson.D)[0].Value.(bson.D)[0].Value != "Mingyong Chen" {
						return errors.New("filter error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn[types.TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[types.TestUser], opts ...any) error {
					user := opCtx.Doc
					if user.Name != "Mingyong Chen" || user.Age != 18 {
						return errors.New("after error")
					}
					return nil
				},
			},
			want: &types.TestUser{
				Name: "Mingyong Chen",
				Age:  18,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			user, err := finder.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).
				FindOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				tc.want.ID = user.ID
				require.Equal(t, tc.want, user)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			finder.beforeHooks = nil
			finder.afterHooks = nil
		})
	}
}

func TestFinder_e2e_Find(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

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
		opts       []*options.FindOptions
		globalHook []globalHook
		beforeHook []beforeHookFn
		afterHook  []afterHookFn[types.TestUser]

		ctx     context.Context
		want    []*types.TestUser
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "nil filter error",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			filter:  nil,
			wantErr: require.Error,
		},
		{
			name: "decode error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.IllegalUser{
						Name: "chenmingyong",
						Age:  "24",
					},
					&types.IllegalUser{
						Name: "burt",
						Age:  "25",
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter:  bson.D{},
			wantErr: require.Error,
		},
		{
			name: "returns empty documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong",
						Age:  24,
					},
					&types.TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter:  query.Eq("name", "cmy"),
			want:    []*types.TestUser{},
			wantErr: require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong",
						Age:  24,
					},
					&types.TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: bson.D{},
			want: []*types.TestUser{
				{
					Name: "chenmingyong",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  25,
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "find by multiple name",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong",
						Age:  24,
					},
					&types.TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "chenmingyong", "burt"),
			want: []*types.TestUser{
				{
					Name: "chenmingyong",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  25,
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "ignore age field",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong",
						Age:  24,
					},
					&types.TestUser{
						Name: "burt",
						Age:  25,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "chenmingyong", "burt"),
			opts: []*options.FindOptions{
				{
					Projection: bsonx.M("age", 0),
				},
			},
			want: []*types.TestUser{
				{
					Name: "chenmingyong",
				},
				{
					Name: "burt",
				},
			},
			wantErr: require.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
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
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&types.TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "Mingyong Chen", "burt"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterFind,
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
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&types.TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "Mingyong Chen", "burt"),
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeFind,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx.Filter == nil {
							return errors.New("filter error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterFind,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						fmt.Println(opCtx.Doc)
						users := opCtx.Doc.([]*types.TestUser)
						if len(users) != 2 {
							return errors.New("result error")
						}
						return nil
					},
				},
			},
			wantErr: require.NoError,
			want: []*types.TestUser{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			filter: query.Eq("name", "Mingyong Chen"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				require.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name: "after hook error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&types.TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "Mingyong Chen", "burt"),
			afterHook: []afterHookFn[types.TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[types.TestUser], opts ...any) error {
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
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "Mingyong Chen",
						Age:  18,
					},
					&types.TestUser{
						Name: "burt",
						Age:  19,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: query.In("name", "Mingyong Chen", "burt"),
			beforeHook: []beforeHookFn{
				func(ctx context.Context, opCtx *OpContext, opts ...any) error {
					if opCtx.Filter == nil {
						return errors.New("filter error")
					}
					return nil
				},
			},
			afterHook: []afterHookFn[types.TestUser]{
				func(ctx context.Context, opCtx *AfterOpContext[types.TestUser], opts ...any) error {
					users := opCtx.Docs
					if len(users) != 2 {
						return errors.New("result error")
					}
					return nil
				},
			},
			wantErr: require.NoError,
			want: []*types.TestUser{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			users, err := finder.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).Filter(tc.filter).Find(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.Equal(t, len(tc.want), len(users))
				for _, user := range users {
					var zero primitive.ObjectID
					user.ID = zero
				}
				require.ElementsMatch(t, tc.want, users)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			finder.beforeHooks = nil
			finder.afterHooks = nil
		})
	}
}

func TestFinder_e2e_Count(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter any
		opts   []*options.CountOptions

		ctx     context.Context
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "nil filter error",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			filter:  nil,
			wantErr: require.Error,
		},
		{
			name:    "returns 0",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			filter:  bson.D{},
			wantErr: require.NoError,
		},
		{
			name: "returns 1",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Name: "chenmingyong",
					Age:  24,
				})
				require.NoError(t, err)
				require.NotNil(t, insertOneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			opts: []*options.CountOptions{
				options.Count().SetComment("test"),
			},
			filter:  bson.D{},
			want:    1,
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			count, err := finder.Filter(tc.filter).Count(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.Equal(t, tc.want, count)
			}
		})
	}
}

func TestFinder_e2e_Distinct(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		fieldName string
		filter    any
		opts      []*options.DistinctOptions

		ctx     context.Context
		want    []any
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			filter:  "name",
			ctx:     context.Background(),
			want:    nil,
			wantErr: require.Error,
		},
		{
			name:      "returns empty documents",
			before:    func(ctx context.Context, t *testing.T) {},
			after:     func(ctx context.Context, t *testing.T) {},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want:      []any{},
			wantErr:   require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*types.TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
					{
						Name: "burt",
						Age:  45,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want: []any{
				"chenmingyong",
				"burt",
			},
			wantErr: require.NoError,
		},
		{
			name: "name distinct",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*types.TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
					{
						Name: "chenmingyong",
						Age:  25,
					},
					{
						Name: "burt",
						Age:  26,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 3)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(3), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			want: []any{
				"chenmingyong",
				"burt",
			},
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			result, err := finder.Filter(tc.filter).Distinct(tc.ctx, tc.fieldName, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.ElementsMatch(t, tc.want, result)
			}
		})
	}
}

func TestFinder_e2e_DistinctWithParse(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		fieldName string
		filter    any
		result    []string
		opts      []*options.DistinctOptions

		ctx     context.Context
		want    []string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			filter:  "name",
			ctx:     context.Background(),
			wantErr: require.Error,
		},
		{
			name:      "returns empty documents",
			before:    func(ctx context.Context, t *testing.T) {},
			after:     func(ctx context.Context, t *testing.T) {},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want:      []string{},
			wantErr:   require.NoError,
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*types.TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
					{
						Name: "burt",
						Age:  45,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want: []string{
				"chenmingyong",
				"burt",
			},
			wantErr: require.NoError,
		},
		{
			name: "name distinct",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]*types.TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
					{
						Name: "chenmingyong",
						Age:  25,
					},
					{
						Name: "burt",
						Age:  26,
					},
				}...))
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 3)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(3), deleteResult.DeletedCount)
			},
			filter:    bson.D{},
			fieldName: "name",
			ctx:       context.Background(),
			result:    []string{},
			want: []string{
				"chenmingyong",
				"burt",
			},
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			err := finder.Filter(tc.filter).DistinctWithParse(tc.ctx, tc.fieldName, &tc.result, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.ElementsMatch(t, tc.want, tc.result)
			}
		})
	}
	t.Run("parse error", func(t *testing.T) {
		var result []int
		err := finder.Filter(bson.D{}).DistinctWithParse(context.Background(), "name", result)
		require.Error(t, err)
	})
}
