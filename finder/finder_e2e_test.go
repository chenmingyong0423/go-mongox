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

////go:build e2e

package finder

import (
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/builder"
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

type illegalUser struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Age  string
}

//type updatedUser struct {
//	Name string `bson:"name"`
//	Age  int
//}
//
//type userName struct {
//	Name string `bson:"name"`
//}

func TestFinder_e2e_New(t *testing.T) {
	collection := getCollection(t)

	result := NewFinder[testUser](collection)
	assert.NotNil(t, result, "Expected non-nil Finder")
	assert.Equal(t, collection, result.collection, "Expected finder field to be initialized correctly")
}

func TestFinder_e2e_One(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[testUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[testUser]) *Finder[testUser]

		ctx     context.Context
		want    *testUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123"))
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Id("456").Build())
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123"))
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder
			},
			want: &testUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by id",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123"))
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Id("123").Build())
			},
			want: &testUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by name",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123"))
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Add("name", "cmy").Build())
			},
			want: &testUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "find by age",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123"))
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Add("age", 18).Build())
			},
			want: &testUser{
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
			user, err := finder.One(tc.ctx)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_e2e_OneWithOptions(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[testUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[testUser]) *Finder[testUser]
		opts   []*options.FindOneOptions

		ctx     context.Context
		want    *testUser
		wantErr error
	}{
		{
			name: "no document",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123").Build())
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Id("456").Build())
			},

			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "returns some of fields",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123").Build())
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Id("123").Build())
			},
			opts: []*options.FindOneOptions{
				{
					Projection: builder.NewBsonBuilder().Add("_id", 1).Add("name", 1).Build(),
				},
			},
			want: &testUser{
				Id:   "123",
				Name: "cmy",
			},
		},
		{
			name: "ignore _id field",
			before: func(ctx context.Context, t *testing.T) {
				insertOneResult, err := collection.InsertOne(ctx, &testUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.Equal(t, insertOneResult.InsertedID.(string), "123")
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteOneResult, err := collection.DeleteOne(ctx, builder.NewBsonBuilder().Id("123").Build())
				assert.Equal(t, int64(1), deleteOneResult.DeletedCount)
				assert.NoError(t, err)
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Id("123").Build())
			},
			opts: []*options.FindOneOptions{
				{
					Projection: builder.NewBsonBuilder().Add("_id", 0).Build(),
				},
			},
			want: &testUser{
				Name: "cmy",
				Age:  18,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			finder = tc.filter(finder)
			user, err := finder.OneWithOptions(tc.ctx, tc.opts)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_e2e_All(t *testing.T) {
	collection := getCollection(t)
	finder := NewFinder[testUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[testUser]) *Finder[testUser]

		ctx     context.Context
		want    []*testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "nil filter error",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
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
					&illegalUser{
						Id:   "123",
						Name: "cmy",
						Age:  "18",
					},
					&illegalUser{
						Id:   "456",
						Name: "cmy",
						Age:  "18",
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Eq("_id", "789").Build())
			},
			want: []*testUser{},
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder
			},
			want: []*testUser{
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().InString("_id", []string{"123", "456"}...).Build())
			},
			want: []*testUser{
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Add("name", "cmy").Build())
			},
			want: []*testUser{
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
			users, err := finder.All(tc.ctx)
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
	finder := NewFinder[testUser](collection)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		filter func(finder *Finder[testUser]) *Finder[testUser]
		opts   []*options.FindOptions

		ctx     context.Context
		want    []*testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "decode error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&illegalUser{
						Id:   "123",
						Name: "cmy",
						Age:  "18",
					},
					&illegalUser{
						Id:   "456",
						Name: "cmy",
						Age:  "18",
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.NoError(t, err)
				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder.Filter(builder.NewBsonBuilder().Eq("_id", "789").Build())
			},
			opts: nil,
			want: []*testUser{},
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder
			},
			opts: nil,
			want: []*testUser{
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder
			},
			opts: []*options.FindOptions{
				{
					Projection: builder.NewBsonBuilder().Add("_id", 1).Add("name", 1).Build(),
				},
			},
			want: []*testUser{
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
					&testUser{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
					&testUser{
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				})
				assert.ElementsMatch(t, []string{"123", "456"}, insertManyResult.InsertedIDs)
				assert.NoError(t, err)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteManyResult, err := collection.DeleteMany(ctx, builder.NewBsonBuilder().InString("_id", "123", "456").Build())
				assert.Equal(t, int64(2), deleteManyResult.DeletedCount)
				assert.NoError(t, err)

				finder.filter = bson.D{}
			},
			filter: func(finder *Finder[testUser]) *Finder[testUser] {
				return finder
			},
			opts: []*options.FindOptions{
				{
					Projection: builder.NewBsonBuilder().Add("_id", 0).Build(),
				},
			},
			want: []*testUser{
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
			users, err := finder.AllWithOptions(tc.ctx, tc.opts)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			assert.ElementsMatch(t, tc.want, users)
		})
	}
}
