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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"github.com/chenmingyong0423/go-mongox/bsonx"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
	deleter := NewDeleter[types.TestTempUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter bson.D
		opts   []*options.DeleteOptions

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
				insertResult, err := collection.InsertOne(ctx, types.TestTempUser{Id: "123", Name: "cmy"})
				require.NoError(t, err)
				require.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("456").Build(),
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
				insertResult, err := collection.InsertOne(ctx, types.TestTempUser{Id: "123", Name: "cmy"})
				require.NoError(t, err)
				require.Equal(t, "123", insertResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: query.BsonBuilder().Id("123").Build(),
			ctx:    context.Background(),
			opts:   []*options.DeleteOptions{options.Delete().SetComment("test")},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantError: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			result, err := deleter.Filter(tc.filter).DeleteOne(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
		})
	}
	t.Run("before hook error", func(t *testing.T) {
		ctx := context.Background()
		callback.GetCallback().Register(operation.OpTypeBeforeDelete, "before hook error", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			return errors.New("before hook error")
		})
		deleteResult, err := deleter.Filter(bsonx.D()).DeleteOne(ctx)
		require.Equal(t, err, errors.New("before hook error"))
		require.Nil(t, deleteResult)
		callback.GetCallback().Remove(operation.OpTypeBeforeDelete, "before hook error")
	})
	t.Run("before hook error", func(t *testing.T) {
		ctx := context.Background()
		callback.GetCallback().Register(operation.OpTypeAfterDelete, "after hook error", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			return errors.New("after hook error")
		})
		insertResult, err := collection.InsertOne(ctx, types.TestUser{Name: "chenmingyong"})
		require.NoError(t, err)
		require.NotNil(t, insertResult.InsertedID)
		deleteResult, err := deleter.Filter(query.Eq("name", "chenmingyong")).DeleteOne(ctx)
		require.Equal(t, err, errors.New("after hook error"))
		require.Nil(t, deleteResult)
		callback.GetCallback().Remove(operation.OpTypeAfterDelete, "after hook error")
	})
}

func TestDeleter_e2e_DeleteMany(t *testing.T) {
	collection := newCollection(t)
	deleter := NewDeleter[types.TestTempUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter any
		opts   []*options.DeleteOptions

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
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]types.TestTempUser{
					{Id: "123", Name: "cmy"},
					{Id: "456", Name: "cmy"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "cmy").Build())
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
				insertResult, err := collection.InsertMany(ctx, utils.ToAnySlice([]types.TestTempUser{
					{Id: "123", Name: "cmy"},
					{Id: "456", Name: "cmy"},
				}...))
				require.NoError(t, err)
				require.ElementsMatch(t, []string{"123", "456"}, insertResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().Eq("name", "cmy").Build())
				require.NoError(t, err)
				require.Equal(t, int64(0), deleteResult.DeletedCount)
			},
			filter: bsonx.M("name", "cmy"),
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
			result, err := deleter.Filter(tc.filter).DeleteMany(tc.ctx, tc.opts...)
			tc.after(tc.ctx, t)
			tc.wantError(t, err)
			require.Equal(t, tc.want, result)
		})
	}
	t.Run("before hook error", func(t *testing.T) {
		ctx := context.Background()
		callback.GetCallback().Register(operation.OpTypeBeforeDelete, "before hook error", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			return errors.New("before hook error")
		})
		deleteResult, err := deleter.Filter(bsonx.D()).DeleteMany(ctx)
		require.Equal(t, err, errors.New("before hook error"))
		require.Nil(t, deleteResult)
		callback.GetCallback().Remove(operation.OpTypeBeforeDelete, "before hook error")
	})
	t.Run("before hook error", func(t *testing.T) {
		ctx := context.Background()
		callback.GetCallback().Register(operation.OpTypeAfterDelete, "after hook error", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			return errors.New("after hook error")
		})
		insertResult, err := collection.InsertOne(ctx, types.TestUser{Name: "chenmingyong"})
		require.NoError(t, err)
		require.NotNil(t, insertResult.InsertedID)
		deleteResult, err := deleter.Filter(query.Eq("name", "chenmingyong")).DeleteMany(ctx)
		require.Equal(t, err, errors.New("after hook error"))
		require.Nil(t, deleteResult)
		callback.GetCallback().Remove(operation.OpTypeAfterDelete, "after hook error")
	})
}
