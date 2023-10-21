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

package mongox

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/chenmingyong0423/go-mongox/internal/types"

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
	Name string `bson:"name"`
	Age  int
}

type userName struct {
	Name string `bson:"name"`
}

func TestCollection_e2e_UpdateMany(t *testing.T) {
	collection := getCollection[testUser](t)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  any
		updates any
		opts    []*options.UpdateOptions

		result  func(t *testing.T, us *mongo.UpdateResult)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil filter",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,

			result: func(t *testing.T, ur *mongo.UpdateResult) {
				assert.Nil(t, ur)
			},
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

			ctx:     context.Background(),
			filter:  bson.D{},
			updates: nil,

			result: func(t *testing.T, ur *mongo.UpdateResult) {
				assert.Nil(t, ur)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "update records which does not exist.",
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
			filter:  map[string]any{types.Id: "789"},
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(0), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is bson",
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
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				two, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("456").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "456", Name: "ccc", Age: 18}, two)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)

				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},
			filter:  bson.D{bson.E{Key: "name", Value: "cmy"}},
			updates: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "name", Value: "ccc"}}}},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(2), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is map",
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
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				two, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("456").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "456", Name: "ccc", Age: 18}, two)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)

				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},
			filter:  map[string]any{"name": "cmy"},
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(2), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is struct",
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
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				two, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("456").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "456", Name: "ccc", Age: 18}, two)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)

				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},
			filter:  userName{Name: "cmy"},
			updates: updatedUser{Name: "ccc", Age: 18},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(2), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is struct pointer",
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
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				two, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("456").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "456", Name: "ccc", Age: 18}, two)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)

				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},
			filter:  &userName{Name: "cmy"},
			updates: &updatedUser{Name: "ccc", Age: 18},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(2), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			updateResult, err := collection.UpdateMany(tc.ctx, tc.filter, tc.updates, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			tc.result(t, updateResult)
		})
	}
}

func TestCollection_e2e_UpdateId(t *testing.T) {
	collection := getCollection[testUser](t)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		id      any
		updates any
		opts    []*options.UpdateOptions

		result  func(t *testing.T, us *mongo.UpdateResult)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "nil updates",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:     context.Background(),
			id:      "",
			updates: nil,

			result: func(t *testing.T, ur *mongo.UpdateResult) {
				assert.Nil(t, ur)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "id does not exist.",
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
			id:      "456",
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(0), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "bson updates",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			id:      "123",
			updates: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "name", Value: "ccc"}}}},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "map updates",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			id:      "123",
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "struct updates",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 24}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			id:      "123",
			updates: updatedUser{Name: "ccc", Age: 24},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "struct pointer updates",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 24}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			id:      "123",
			updates: &updatedUser{Name: "ccc", Age: 24},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			updateResult, err := collection.UpdateId(tc.ctx, tc.id, tc.updates, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			tc.result(t, updateResult)
		})
	}
}

func TestCollection_e2e_UpdateOne(t *testing.T) {
	collection := getCollection[testUser](t)
	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx     context.Context
		filter  any
		updates any
		opts    []*options.UpdateOptions

		result  func(t *testing.T, us *mongo.UpdateResult)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nil filter",
			before:  func(_ context.Context, _ *testing.T) {},
			after:   func(_ context.Context, _ *testing.T) {},
			ctx:     context.Background(),
			filter:  nil,
			updates: nil,

			result: func(t *testing.T, ur *mongo.UpdateResult) {
				assert.Nil(t, ur)
			},
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

			ctx:     context.Background(),
			filter:  bson.D{},
			updates: nil,

			result: func(t *testing.T, ur *mongo.UpdateResult) {
				assert.Nil(t, ur)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "update a record which does not exist.",
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
			filter:  map[string]any{types.Id: "456"},
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(0), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is bson",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			filter:  bson.D{bson.E{Key: types.Id, Value: "123"}},
			updates: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "name", Value: "ccc"}}}},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is map",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 18}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			filter:  map[string]any{types.Id: "123"},
			updates: map[string]any{"name": "ccc"},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 24}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			filter:  testUser{Id: "123", Name: "cmy", Age: 18},
			updates: updatedUser{Name: "ccc", Age: 24},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "the type of filter and updates is struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				one, err := collection.FindOne(context.Background(), NewBsonBuilder().Id("123").Build())
				assert.NoError(t, err)
				assert.Equal(t, &testUser{Id: "123", Name: "ccc", Age: 24}, one)

				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},
			filter:  &testUser{Id: "123", Name: "cmy", Age: 18},
			updates: &updatedUser{Name: "ccc", Age: 24},
			result: func(t *testing.T, us *mongo.UpdateResult) {
				assert.Equal(t, int64(1), us.ModifiedCount)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got one %v", err)
					return false
				}
				return true
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			updateResult, err := collection.UpdateOne(tc.ctx, tc.filter, tc.updates, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantErr(t, err) {
				return
			}
			tc.result(t, updateResult)
		})
	}
}

func TestCollection_e2e_FindOne(t *testing.T) {
	collection := getCollection[testUser](t)
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
			filter: bson.D{bson.E{Key: types.Id, Value: "123"}},

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
	collection := getCollection[testUser](t)
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
				_, fErr := collection.collection.InsertOne(ctx, NewBsonBuilder().Add(types.Id, "123").Add("name", "cmy").Add("age", "18").Build())
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
			filter: bson.D{bson.E{Key: types.Id, Value: "123"}},

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
	collection := getCollection[testUser](t)

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

func getCollection[T any](t *testing.T) *Collection[T] {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := NewCollection[T](client.Database("db-test").Collection("test_user"))
	return collection
}

func TestCollection_e2e_Insert(t *testing.T) {
	collection := getCollection[testUser](t)

	// InsertOne
	{
		user := testUser{
			Id:   "123",
			Name: "cmy",
			Age:  24,
		}
		insertOneResult, err := collection.InsertOne(context.Background(), user)
		assert.NoError(t, err)
		assert.Equal(t, "123", insertOneResult.InsertedID.(string))

		foundUser := new(testUser)
		err = collection.collection.FindOneAndDelete(context.Background(), NewBsonBuilder().Id("123").Build()).Decode(foundUser)
		assert.NoError(t, err)
		assert.Equal(t, &user, foundUser)
	}

	// InsertMany
	{
		users := []testUser{
			{Id: "123", Name: "cmy", Age: 24},
			{Id: "456", Name: "cmy", Age: 24},
		}
		ids := []any{"123", "456"}
		insertManyResult, err := collection.InsertMany(context.Background(), users)
		assert.NoError(t, err)
		assert.Equal(t, ids, insertManyResult.InsertedIDs)

		foundUsers, err := collection.Find(context.Background(), NewBsonBuilder().In("_id", ids...).Build())
		assert.NoError(t, err)
		fmt.Println(foundUsers)
		assert.True(t, reflect.DeepEqual([]*testUser{
			{Id: "123", Name: "cmy", Age: 24},
			{Id: "456", Name: "cmy", Age: 24},
		}, foundUsers))

		deleteResult, err := collection.collection.DeleteMany(context.Background(), NewBsonBuilder().In("_id", ids...))
		assert.NoError(t, err)
		assert.Equal(t, int64(2), deleteResult.DeletedCount)
	}
}
