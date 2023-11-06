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
	"testing"

	"github.com/chenmingyong0423/go-mongox/converter"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

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

func TestFinder_e2e_New(t *testing.T) {
	collection := getCollection(t)

	result := NewFinder[types.TestUser](collection)
	assert.NotNil(t, result, "Expected non-nil Finder")
	assert.Equal(t, collection, result.collection, "Expected finder field to be initialized correctly")
}

func TestFinder_e2e_One(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[types.TestUser]) *Finder[types.TestUser]

		ctx     context.Context
		want    *types.TestUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, err)
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Id("456").Build())
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, err)
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by id",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, err)
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Id("123").Build())
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by name",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Add(converter.KeyValue("name", "cmy")).Build())
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by age",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, err)
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123"))
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Add(converter.KeyValue("age", 18)).Build())
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			finder = tc.filter(finder)
			user, err := finder.FindOne(tc.ctx)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_e2e_OneWithOptions(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[types.TestUser]) *Finder[types.TestUser]
		opts   []*options.FindOneOptions

		ctx     context.Context
		want    *types.TestUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Id("456").Build())
			},

			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "returns some of fields",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.FilterKeyValue(converter.KeyValue("_id", "123"))
			},
			opts: []*options.FindOneOptions{
				{
					Projection: query.BsonBuilder().Add(converter.KeyValue("_id", 1), converter.KeyValue("name", 1)).Build(),
				},
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
			},
		},
		{
			name: "ignore _id field",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, err)
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, query.BsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Id("123").Build())
			},
			opts: []*options.FindOneOptions{
				{
					Projection: query.BsonBuilder().Add(converter.KeyValue("_id", 0)).Build(),
				},
			},
			want: &types.TestUser{
				Name: "cmy",
				Age:  18,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			finder = tc.filter(finder)
			user, err := finder.FindOneWithOptions(tc.ctx, tc.opts)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_e2e_All(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[types.TestUser]) *Finder[types.TestUser]

		ctx     context.Context
		want    []*types.TestUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(nil)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected error, got nil")
					return false
				}
				return true
			},
		},
		{
			name: "decode error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.IllegalUser{
						Id:   "123",
						Name: "cmy",
						Age:  "18",
					},
					&types.IllegalUser{
						Id:   "456",
						Name: "cmy",
						Age:  "18",
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected error, got nil")
					return false
				}
				return true
			},
		},
		{
			name: "returns empty documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Eq("_id", "789").Build())
			},
			want: []*types.TestUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			want: []*types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "find by multiple id",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().InString("_id", []string{"123", "456"}...).Build())
			},
			want: []*types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "find by name",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Add(converter.KeyValue("name", "cmy")).Build())
			},
			want: []*types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			finder = tc.filter(finder)
			users, err := finder.FindAll(tc.ctx)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.ElementsMatch(t, tc.want, users)
		})
	}
}

func TestFinder_e2e_AllWithOptions(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[types.TestUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[types.TestUser]) *Finder[types.TestUser]
		opts   []*options.FindOptions

		ctx     context.Context
		want    []*types.TestUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "decode error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.IllegalUser{
						Id:   "123",
						Name: "cmy",
						Age:  "18",
					},
					&types.IllegalUser{
						Id:   "456",
						Name: "cmy",
						Age:  "18",
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected error, got nil")
					return false
				}
				return true
			},
		},
		{
			name: "returns empty documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)

			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder.Filter(query.BsonBuilder().Eq("_id", "789").Build())
			},
			opts: nil,
			want: []*types.TestUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "returns all documents",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)

			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			opts: nil,
			want: []*types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				},
				{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "returns some of fields",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			opts: []*options.FindOptions{
				{
					Projection: query.BsonBuilder().Add(converter.KeyValue("_id", 1)).Add(converter.KeyValue("name", 1)).Build(),
				},
			},
			want: []*types.TestUser{
				{
					Id:   "123",
					Name: "cmy",
				},
				{
					Id:   "456",
					Name: "cmy",
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "ignore _id field",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&types.TestUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.NoError(t, err)
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[types.TestUser]) *Finder[types.TestUser] {
				return finder
			},
			opts: []*options.FindOptions{
				{
					Projection: query.BsonBuilder().Add(converter.KeyValue("_id", 0)).Build(),
				},
			},
			want: []*types.TestUser{
				{
					Name: "cmy",
					Age:  18,
				},
				{
					Name: "cmy",
					Age:  18,
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected nil, got error: %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			finder = tc.filter(finder)
			users, err := finder.FindAllWithOptions(tc.ctx, tc.opts)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.ElementsMatch(t, tc.want, users)
		})
	}
}
