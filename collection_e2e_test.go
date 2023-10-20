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

package mongox

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type testUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int
	UnknownField string `bson:"-"`
}

type updatedUser struct {
	Name         string `bson:"name"`
	Age          int
	UnknownField string `bson:"-"`
}

func TestCollection_e2e_FindOneAndReplace(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx         context.Context
		filter      any
		replacement any
		opts        []*options.FindOneAndReplaceOptions

		wantT   *testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "filter is not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil filter",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx: context.Background(),

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil replacement",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:         context.Background(),
			filter:      bson.D{},
			replacement: nil,
			wantT:       nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNilDocument) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty replacement",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:         context.Background(),
			filter:      NewBsonBuilder().Id("123").Build(),
			replacement: map[string]any{},
			wantT:       &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "replace by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, "ccc", user.Name)
				assert.NoError(t, err)

				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:         context.Background(),
			filter:      NewBsonBuilder().Id("123").Build(),
			replacement: testUser{Id: "123", Name: "ccc"},
			wantT:       &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindOneAndReplace(tc.ctx, tc.filter, tc.replacement, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_FindOneAndUpdate(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  any
		updates any
		opts    []*options.FindOneAndUpdateOptions

		wantT   *testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "filter is not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:     context.Background(),
			filter:  1,
			updates: bson.D{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "update is not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:     context.Background(),
			filter:  bson.D{},
			updates: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil filter",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx: context.Background(),

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil updates",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D updates",
			before: func(ctx context.Context, t *testing.T) {
			},
			after: func(ctx context.Context, t *testing.T) {
			},

			ctx:     context.Background(),
			filter:  bson.D{},
			updates: bson.D{},
			wantT:   nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "update by bson.D",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "ccc", user.Name)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:     context.Background(),
			filter:  NewBsonBuilder().Id("123").Build(),
			updates: bson.D{bson.E{Key: set, Value: bson.D{bson.E{Key: "name", Value: "ccc"}}}},
			wantT:   &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map updates",
			before: func(ctx context.Context, t *testing.T) {
			},
			after: func(ctx context.Context, t *testing.T) {},

			ctx:     context.Background(),
			filter:  map[string]any{},
			updates: map[string]any{},
			wantT:   nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "update by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "ccc", user.Name)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: NewBsonBuilder().Id("123").Build(),
			updates: map[string]any{
				"name": "ccc",
			},
			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "", user.Name)
				assert.Equal(t, 0, user.Age)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:     context.Background(),
			filter:  NewBsonBuilder().Id("123").Build(),
			updates: updatedUser{},
			wantT:   &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "update by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "ccc", user.Name)
				assert.Equal(t, 24, user.Age)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:     context.Background(),
			filter:  NewBsonBuilder().Id("123").Build(),
			updates: updatedUser{Name: "ccc", Age: 24},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "ccc", user.Name)
				assert.Equal(t, 24, user.Age)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:     context.Background(),
			filter:  NewBsonBuilder().Id("123").Build(),
			updates: &updatedUser{Name: "ccc", Age: 24},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "update by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, err := collection.FindOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, "ccc", user.Name)
				assert.Equal(t, 24, user.Age)
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:     context.Background(),
			filter:  NewBsonBuilder().Id("123").Build(),
			updates: &updatedUser{Name: "ccc", Age: 24},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindOneAndUpdate(tc.ctx, tc.filter, tc.updates, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_FindOneAndDelete(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx    context.Context
		filter any
		opts   []*options.FindOneAndDeleteOptions

		wantT   *testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindById(ctx, NewBsonBuilder().Id("123"))
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)
				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindById(ctx, bson.D{bson.E{Key: id, Value: "123"}})
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)
				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{bson.E{Key: id, Value: "123"}},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindById(ctx, NewBsonBuilder().Id("123"))
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)
				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: map[string]any{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindById(ctx, map[string]any{
					"_id": "123",
				})
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)
				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx: context.Background(),
			filter: map[string]any{
				"_id": "123",
			},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindOne(ctx, testUser{Id: "123", Name: "cmy", Age: 18})
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)

				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				result1, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.Equal(t, result1.DeletedCount, int64(1))
				assert.NoError(t, fErr)
				result2, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result2.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				user, fErr := collection.FindOne(ctx, &testUser{Id: "123", Name: "cmy", Age: 18})
				assert.Equal(t, mongo.ErrNoDocuments, fErr)
				assert.Nil(t, user)

				result, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.Equal(t, result.DeletedCount, int64(1))
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindOneAndDelete(tc.ctx, tc.filter, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_FindOne(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx    context.Context
		filter any
		opts   []*options.FindOneOptions

		wantT   *testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil filter",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: nil,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{bson.E{Key: id, Value: "123"}},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: map[string]any{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx: context.Background(),
			filter: map[string]any{
				"_id": "123",
			},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindOne(tc.ctx, tc.filter, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_Find(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx    context.Context
		filter any
		opts   []*options.FindOptions

		wantT   []*testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "456", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "decode failed",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, NewBsonBuilder().Add(id, "123").Add("name", "cmy").Add("age", "18").Build())
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: NewBsonBuilder().Id("123").Build(),

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "get one by bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{bson.E{Key: id, Value: "123"}},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: map[string]any{},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "456", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx: context.Background(),
			filter: map[string]any{
				"_id": "123",
			},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{},

			wantT: []*testUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{},

			wantT: []*testUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.Find(tc.ctx, tc.filter, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_FindById(t *testing.T) {
	collection := getCollection(t)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx  context.Context
		id   string
		opts []*options.FindOneOptions

		wantT   *testUser
		wantErr error
	}{
		{
			name:   "no document",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:  context.Background(),
			id:   "123",
			opts: nil,

			wantT:   nil,
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "found",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},

			ctx:  context.Background(),
			id:   "123",
			opts: nil,

			wantT: &testUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindById(tc.ctx, tc.id, tc.opts...)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantT, gotT)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func getCollection(t *testing.T) *Collection[testUser] {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := NewCollection[testUser](client.Database("db-test").Collection("test_user"))
	return collection
}
