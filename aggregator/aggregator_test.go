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

package aggregator

import (
	"context"
	"errors"
	"testing"

	"github.com/chenmingyong0423/go-mongox/builder/aggregation"
	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/mock/gomock"
)

func TestNewAggregator(t *testing.T) {
	mongoCollection := &mongo.Collection{}
	aggregator := NewAggregator[any](mongoCollection)

	assert.NotNil(t, aggregator)
	assert.Equal(t, mongoCollection, aggregator.collection)
}

func TestAggregator_Aggregation(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser]
		ctx      context.Context
		pipeline any

		want    []*types.TestUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().Aggregation(ctx, nil).Return(nil, errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:      context.Background(),
			pipeline: nil,
			wantErr:  assert.Error,
		},
		{
			name: "empty result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().Aggregation(ctx, mongo.Pipeline{}).Return([]*types.TestUser{}, nil).Times(1)
				return aggregator
			},
			ctx:      context.Background(),
			pipeline: mongo.Pipeline{},
			want:     []*types.TestUser{},
			wantErr:  assert.NoError,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().Aggregation(ctx, mongo.Pipeline{
					bson.D{
						bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "age", Value: -1}}},
					},
				}).Return([]*types.TestUser{
					{Id: "1", Name: "cmy", Age: 24},
					{Id: "2", Name: "gopher", Age: 20},
				}, nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			pipeline: aggregation.StageBsonBuilder().
				Sort("age", -1).Build(),
			want: []*types.TestUser{
				{Id: "2", Name: "gopher", Age: 20},
				{Id: "1", Name: "cmy", Age: 24},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			aggregator := tc.mock(tc.ctx, ctl)

			result, err := aggregator.Aggregation(tc.ctx, tc.pipeline)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, result)
			}
		})
	}
}

func TestAggregator_AggregationWithOptions(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser]
		ctx      context.Context
		pipeline any
		opts     *options.AggregateOptions
		want     []*types.TestUser
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithOptions(ctx, nil, options.Aggregate().SetBatchSize(1)).Return(nil, errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			opts:     options.Aggregate().SetBatchSize(1),
			ctx:      context.Background(),
			pipeline: nil,
			wantErr:  assert.Error,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithOptions(ctx, mongo.Pipeline{
					bson.D{
						bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "age", Value: -1}}},
					},
				}, options.Aggregate().SetBatchSize(1)).Return([]*types.TestUser{
					{Id: "1", Name: "cmy", Age: 24},
				}, nil).Times(1)
				return aggregator
			},
			opts: options.Aggregate().SetBatchSize(1),
			ctx:  context.Background(),
			pipeline: aggregation.StageBsonBuilder().
				Sort("age", -1).Build(),
			want: []*types.TestUser{
				{Id: "1", Name: "cmy", Age: 24},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			aggregator := tc.mock(tc.ctx, ctl)

			result, err := aggregator.AggregationWithOptions(tc.ctx, tc.pipeline, tc.opts)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, result)
			}
		})
	}
}

func TestAggregator_AggregationWithCallback(t *testing.T) {
	type User struct {
		Id           string `bson:"_id"`
		Name         string `bson:"name"`
		Age          int64
		IsProgrammer bool `bson:"is_programmer"`
	}
	testCases := []struct {
		name          string
		mock          func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser]
		ctx           context.Context
		pipeline      any
		callbackParam []*User
		want          []*User
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallback(ctx, nil, gomock.Any()).Return(errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:           context.Background(),
			callbackParam: []*User{},
			pipeline:      nil,
			wantErr:       assert.Error,
		},
		{
			name: "empty result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallback(ctx, mongo.Pipeline{}, gomock.Any()).Return(nil).Times(1)
				return aggregator
			},
			ctx:           context.Background(),
			pipeline:      mongo.Pipeline{},
			callbackParam: []*User{},
			want:          []*User{},
			wantErr:       assert.NoError,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallback(ctx, mongo.Pipeline{
					bson.D{
						bson.E{Key: "$set", Value: bson.D{bson.E{Key: "is_programmer", Value: true}}},
					},
				}, gomock.Any()).Return(nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			pipeline: aggregation.StageBsonBuilder().
				Set("is_programmer", true).Build(),
			callbackParam: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
			},
			want: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			var callback types.ResultHandler = func(cursor *mongo.Cursor) error {
				return cursor.All(tc.ctx, &tc.callbackParam)
			}
			aggregator := tc.mock(tc.ctx, ctl)
			err := aggregator.AggregationWithCallback(tc.ctx, tc.pipeline, callback)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, tc.callbackParam)
			}
		})
	}
}

func TestAggregator_AggregationWithCallbackAndOptions(t *testing.T) {
	type User struct {
		Id           string `bson:"_id"`
		Name         string `bson:"name"`
		Age          int64
		IsProgrammer bool `bson:"is_programmer"`
	}
	testCases := []struct {
		name          string
		mock          func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser]
		ctx           context.Context
		pipeline      any
		opts          *options.AggregateOptions
		callbackParam []*User
		want          []*User
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallbackAndOptions(ctx, nil, options.Aggregate().SetBatchSize(1), gomock.Any()).Return(errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:           context.Background(),
			callbackParam: []*User{},
			opts:          options.Aggregate().SetBatchSize(1),
			pipeline:      nil,
			wantErr:       assert.Error,
		},
		{
			name: "empty result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallbackAndOptions(ctx, mongo.Pipeline{}, options.Aggregate().SetBatchSize(1), gomock.Any()).Return(nil).Times(1)
				return aggregator
			},
			ctx:           context.Background(),
			pipeline:      mongo.Pipeline{},
			callbackParam: []*User{},
			opts:          options.Aggregate().SetBatchSize(1),
			want:          []*User{},
			wantErr:       assert.NoError,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregationWithCallbackAndOptions(ctx, mongo.Pipeline{
					bson.D{
						bson.E{Key: "$set", Value: bson.D{bson.E{Key: "is_programmer", Value: true}}},
					},
				}, options.Aggregate().SetBatchSize(1), gomock.Any()).Return(nil).Times(1)
				return aggregator
			},
			opts: options.Aggregate().SetBatchSize(1),
			ctx:  context.Background(),
			pipeline: aggregation.StageBsonBuilder().
				Set("is_programmer", true).Build(),
			callbackParam: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
			},
			want: []*User{
				{Id: "1", Name: "cmy", Age: 24, IsProgrammer: true},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			var callback types.ResultHandler = func(cursor *mongo.Cursor) error {
				return cursor.All(tc.ctx, &tc.callbackParam)
			}
			aggregator := tc.mock(tc.ctx, ctl)
			err := aggregator.AggregationWithCallbackAndOptions(tc.ctx, tc.pipeline, tc.opts, callback)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, tc.callbackParam)
			}
		})
	}
}
