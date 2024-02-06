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

package aggregator

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/bsonx"

	"github.com/chenmingyong0423/go-mongox/builder/aggregation"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/types"
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

func TestAggregator_e2e_New(t *testing.T) {
	collection := getCollection(t)

	result := NewAggregator[types.TestUser](collection)
	require.NotNil(t, result, "Expected non-nil Aggregator")
	require.Equal(t, collection, result.collection, "Expected collection field to be initialized correctly")
}

func TestAggregator_e2e_Aggregation(t *testing.T) {
	collection := getCollection(t)
	aggregator := NewAggregator[types.TestUser](collection)

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		pipeline           any
		aggregationOptions []*options.AggregateOptions

		ctx     context.Context
		want    []*types.TestUser
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:   "got error when pipeline is nil",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			pipeline:           nil,
			aggregationOptions: nil,

			ctx:     context.Background(),
			want:    nil,
			wantErr: require.Error,
		},
		{
			name: "decode error",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.IllegalUser{
						Name: "chenmingyong", Age: "24",
					},
					&types.IllegalUser{
						Name: "gopher", Age: "20",
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "gopher"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline:           mongo.Pipeline{},
			aggregationOptions: nil,
			want:               []*types.TestUser{},
			ctx:                context.Background(),
			wantErr:            require.Error,
		},
		{
			name: "got result when pipeline is empty",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong", Age: 24,
					},
					&types.TestUser{
						Name: "gopher", Age: 20,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "gopher"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline:           mongo.Pipeline{},
			aggregationOptions: nil,
			want: []*types.TestUser{
				{
					Name: "chenmingyong", Age: 24,
				},
				{
					Name: "gopher", Age: 20,
				},
			},
			ctx:     context.Background(),
			wantErr: require.NoError,
		},
		{
			name: "got result by pipeline with match stage",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong", Age: 24,
					},
					&types.TestUser{
						Name: "gopher", Age: 20,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "gopher"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline: aggregation.StageBsonBuilder().Sort(bsonx.M("age", -1)).Build(),
			want: []*types.TestUser{
				{
					Name: "chenmingyong", Age: 24,
				},
				{
					Name: "gopher", Age: 20,
				},
			},
			ctx:     context.Background(),
			wantErr: require.NoError,
		},
		{
			name: "got result with aggregation options",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					&types.TestUser{
						Name: "chenmingyong", Age: 24,
					},
					&types.TestUser{
						Name: "gopher", Age: 20,
					},
				})
				require.NoError(t, err)
				require.Len(t, insertManyResult.InsertedIDs, 2)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "chenmingyong", "gopher"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline: aggregation.StageBsonBuilder().Sort(bsonx.M("age", 1)).Build(),
			aggregationOptions: []*options.AggregateOptions{
				options.Aggregate().SetCollation(&options.Collation{Locale: "en", Strength: 2}),
			},
			want: []*types.TestUser{
				{
					Name: "chenmingyong", Age: 24,
				},
				{
					Name: "gopher", Age: 20,
				},
			},
			ctx:     context.Background(),
			wantErr: require.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			testUsers, err := aggregator.Pipeline(tc.pipeline).Aggregate(tc.ctx, tc.aggregationOptions...)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.Equal(t, len(tc.want), len(testUsers))
				for _, tu := range testUsers {
					var zero primitive.ObjectID
					tu.ID = zero
				}
				require.ElementsMatch(t, tc.want, testUsers)
			}
		})
	}
}

func TestAggregator_e2e_AggregateWithParse(t *testing.T) {
	collection := getCollection(t)
	aggregator := NewAggregator[types.TestUser](collection)

	type User struct {
		Id           string `bson:"_id"`
		Name         string `bson:"name"`
		Age          int64
		IsProgrammer bool `bson:"is_programmer"`
	}

	testCases := []struct {
		name               string
		before             func(ctx context.Context, t *testing.T)
		after              func(ctx context.Context, t *testing.T)
		pipeline           any
		aggregationOptions *options.AggregateOptions
		ctx                context.Context
		result             any
		want               []*User
		wantErr            require.ErrorAssertionFunc
	}{
		{
			name:   "got error when pipeline is nil",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			pipeline:           nil,
			aggregationOptions: nil,
			ctx:                context.Background(),
			want:               nil,
			wantErr:            require.Error,
		},
		{
			name: "got result by pipeline with match stage",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					types.TestTempUser{Id: "2", Name: "gopher", Age: 20},
					types.TestTempUser{Id: "1", Name: "cmy", Age: 24},
				})
				require.NoError(t, err)
				require.ElementsMatch(t, []any{"1", "2"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline: aggregation.StageBsonBuilder().Set(bsonx.M("is_programmer", true)).Build(),
			result:   make([]*User, 0, 4),
			want: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
				{Id: "2", Name: "gopher", Age: 20, IsProgrammer: true},
			},
			ctx:     context.Background(),
			wantErr: require.NoError,
		},
		{
			name: "got result with aggregation options",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					types.TestTempUser{Id: "2", Name: "gopher", Age: 20},
					types.TestTempUser{Id: "1", Name: "cmy", Age: 24},
				})
				require.NoError(t, err)
				require.ElementsMatch(t, []any{"1", "2"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline: aggregation.StageBsonBuilder().Set(bsonx.M("is_programmer", true)).Sort(bsonx.M("name", 1)).Build(),
			result:   make([]*User, 0, 4),
			want: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
				{Id: "2", Name: "gopher", Age: 20, IsProgrammer: true},
			},
			aggregationOptions: options.Aggregate().SetCollation(&options.Collation{Locale: "en", Strength: 2}),
			ctx:                context.Background(),
			wantErr:            require.NoError,
		},
		{
			name: "got error from cursor",
			before: func(ctx context.Context, t *testing.T) {
				insertManyResult, err := collection.InsertMany(ctx, []any{
					types.TestTempUser{Id: "2", Name: "gopher", Age: 20},
					types.TestTempUser{Id: "1", Name: "cmy", Age: 24},
				})
				require.NoError(t, err)
				require.ElementsMatch(t, []any{"1", "2"}, insertManyResult.InsertedIDs)
			},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.BsonBuilder().InString("_id", []string{"1", "2"}...).Build())
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			pipeline:           aggregation.StageBsonBuilder().Set(bsonx.M("is_programmer", true)).Sort(bsonx.M("name", 1)).Build(),
			result:             []string{},
			want:               []*User{},
			aggregationOptions: options.Aggregate().SetCollation(&options.Collation{Locale: "en", Strength: 2}),
			ctx:                context.Background(),
			wantErr:            require.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			err := aggregator.Pipeline(tc.pipeline).AggregateWithParse(tc.ctx, &tc.result, tc.aggregationOptions)
			tc.after(tc.ctx, t)
			tc.wantErr(t, err)
			if err == nil {
				require.ElementsMatch(t, tc.want, tc.result)
			}
		})
	}
}
