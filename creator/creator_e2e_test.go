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

	"github.com/stretchr/testify/require"

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
		opts   []*options.InsertOneOptions
		ctx    context.Context
		doc    *types.TestUser

		wantId    string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name:   "nil doc",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc:       nil,
			wantError: assert.Error,
		},
		{
			name:   "insert one successfully",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "chenmingyong"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &types.TestUser{
				Name: "chenmingyong",
				Age:  24,
			},
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
			insertOneResult, err := creator.InsertOne(tc.ctx, tc.doc, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if err == nil {
				require.NotNil(t, insertOneResult.InsertedID)
				require.NotZero(t, tc.doc.CreatedAt)
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

		ctx  context.Context
		docs []*types.TestUser
		opts []*options.InsertManyOptions

		wantIdsLength int
		wantError     assert.ErrorAssertionFunc
	}{
		{
			name:   "nil docs",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			ctx:       context.Background(),
			docs:      nil,
			wantError: assert.Error,
		},
		{
			name:   "insert many successfully",
			before: func(_ context.Context, _ *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			ctx: context.Background(),
			docs: []*types.TestUser{
				{
					Name: "chenmingyong",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  24,
				},
			},
			wantIdsLength: 2,
			wantError:     assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			insertManyResult, err := creator.InsertMany(tc.ctx, tc.docs, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if err == nil {
				require.NotNil(t, insertManyResult)
				require.Len(t, insertManyResult.InsertedIDs, tc.wantIdsLength)
				for _, doc := range tc.docs {
					require.NotZero(t, doc.CreatedAt)
				}
			}
		})
	}
}
