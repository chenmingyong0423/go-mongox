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

	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestAggregator_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}
	aggregator := NewAggregator[any](mongoCollection)

	assert.NotNil(t, aggregator)
	assert.Equal(t, mongoCollection, aggregator.collection)
}

func TestAggregator_Aggregation(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser]
		ctx     context.Context
		want    []*types.TestUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().Aggregate(ctx).Return(nil, errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:     context.Background(),
			wantErr: assert.Error,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().Aggregate(ctx).Return([]*types.TestUser{
					{Name: "chenmingyong", Age: 24},
					{Name: "gopher", Age: 25},
				}, nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			want: []*types.TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "gopher", Age: 25},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			aggregator := tc.mock(tc.ctx, ctl)

			result, err := aggregator.Aggregate(tc.ctx)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, result)
			}
		})
	}
}

func TestAggregator_AggregateWithParse(t *testing.T) {
	type User struct {
		Id           string `bson:"_id"`
		Name         string `bson:"name"`
		Age          int64
		IsProgrammer bool `bson:"is_programmer"`
	}
	testCases := []struct {
		name    string
		mock    func(ctx context.Context, ctl *gomock.Controller, result any) iAggregator[types.TestUser]
		ctx     context.Context
		result  []*User
		want    []*User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller, result any) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregateWithParse(ctx, result).Return(errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:     context.Background(),
			result:  []*User{},
			wantErr: assert.Error,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller, result any) iAggregator[types.TestUser] {
				aggregator := mocks.NewMockiAggregator[types.TestUser](ctl)
				aggregator.EXPECT().AggregateWithParse(ctx, result).Return(nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			result: []*User{
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

			aggregator := tc.mock(tc.ctx, ctl, tc.result)
			err := aggregator.AggregateWithParse(tc.ctx, tc.result)
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, tc.result)
			}
		})
	}
}
