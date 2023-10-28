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

package creator

import (
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := client.Database("db-test").Collection("test_user")
	return collection
}

func TestCreator_e2e_One(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[types.TestUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx context.Context
		t   types.TestUser

		wantId    string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "duplicate",
			before: func(ctx context.Context, t *testing.T) {
				oneResult, err := collection.InsertOne(ctx, types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				})
				assert.NoError(t, err)
				assert.Equal(t, "123", oneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			wantId: "",
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return mongo.IsDuplicateKeyError(err)
			},
		},
		{
			name:   "insert one successfully",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			wantId: "123",
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			insertOneResult, err := creator.InsertOne(tc.ctx, tc.t)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if insertOneResult != nil {
				assert.Equal(t, tc.wantId, insertOneResult.InsertedID)
			}
		})
	}
}

func TestCreator_e2e_OneWithOptions(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[types.TestUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx  context.Context
		opts []*options.InsertOneOptions
		t    types.TestUser

		wantId    string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "duplicate",
			before: func(ctx context.Context, t *testing.T) {
				oneResult, err := collection.InsertOne(ctx, types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				})
				assert.Equal(t, "123", oneResult.InsertedID)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			ctx: context.Background(),
			t: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			wantId: "",
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return mongo.IsDuplicateKeyError(err)
			},
		},
		{
			name:   "insert one successfully",
			before: func(_ context.Context, _ *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			wantId: "123",
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			insertOneResult, err := creator.InsertOneWithOptions(tc.ctx, tc.t, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if insertOneResult != nil {
				assert.Equal(t, tc.wantId, insertOneResult.InsertedID)
			}
		})
	}
}

func TestCreator_e2e_Many(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[types.TestUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx context.Context
		t   []types.TestUser

		wantIds   []string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "duplicate",
			before: func(ctx context.Context, t *testing.T) {
				oneResult, err := collection.InsertOne(ctx, types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				})
				assert.NoError(t, err)
				assert.Equal(t, "123", oneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: []types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				},
			},
			wantIds: nil,
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return mongo.IsDuplicateKeyError(err)
			},
		},
		{
			name:   "insert many successfully",
			before: func(_ context.Context, _ *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: []types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  24,
				},
			},
			wantIds:   []string{"123", "456"},
			wantError: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			insertManyResult, err := creator.InsertMany(tc.ctx, tc.t)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if insertManyResult != nil {
				assert.ElementsMatch(t, tc.wantIds, insertManyResult.InsertedIDs)
			}
		})
	}
}

func TestCreator_e2e_ManyWithOptions(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[types.TestUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		opts []*options.InsertManyOptions

		ctx context.Context
		t   []types.TestUser

		wantIds   []string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "duplicate",
			before: func(ctx context.Context, t *testing.T) {
				oneResult, err := collection.InsertOne(ctx, types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				})
				assert.NoError(t, err)
				assert.Equal(t, "123", oneResult.InsertedID)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			t: []types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				},
			},
			wantIds: nil,
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return mongo.IsDuplicateKeyError(err)
			},
		},
		{
			name:   "insert many successfully",
			before: func(_ context.Context, _ *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			t: []types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  24,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  24,
				},
			},
			wantIds:   []string{"123", "456"},
			wantError: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			insertManyResult, err := creator.InsertManyWithOptions(tc.ctx, tc.t, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if insertManyResult != nil {
				assert.ElementsMatch(t, tc.wantIds, insertManyResult.InsertedIDs)
			}
		})
	}
}
