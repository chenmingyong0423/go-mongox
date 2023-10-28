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

package finder

import (
	"context"
	"errors"
	"testing"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/chenmingyong0423/go-mongox/types"

	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/mock/gomock"
)

func TestFinder_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}

	result := NewFinder[any](mongoCollection)
	assert.NotNil(t, result, "Expected non-nil Finder")
	assert.Equal(t, mongoCollection, result.collection, "Expected finder field to be initialized correctly")
}

func TestFinder_One(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser]
		ctx  context.Context

		want    *types.TestUser
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any()).Return(nil, mongo.ErrNoDocuments).Times(1)
				return mockCollection
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any()).Return(&types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				}, nil).Times(1)
				return mockCollection
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
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			user, err := finder.FindOne(tc.ctx)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_OneWithOptions(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, opts []*options.FindOneOptions, ctl *gomock.Controller) iFinder[types.TestUser]
		ctx  context.Context
		opts []*options.FindOneOptions

		want    *types.TestUser
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, opts []*options.FindOneOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOneWithOptions(ctx, opts).Return(nil, mongo.ErrNoDocuments).Times(1)
				return mockCollection
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			mock: func(ctx context.Context, opts []*options.FindOneOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOneWithOptions(ctx, opts).Return(&types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				}, nil).Times(1)
				return mockCollection
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
		},
		{
			name: "returns some of fields",
			mock: func(ctx context.Context, opts []*options.FindOneOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOneWithOptions(ctx, opts).Return(&types.TestUser{
					Id:   "123",
					Name: "cmy",
				}, nil).Times(1)
				return mockCollection
			},
			opts: []*options.FindOneOptions{
				{
					Projection: query.BsonBuilder().Add("_id", 1).Add("name", 1).Build(),
				},
			},
			want: &types.TestUser{
				Id:   "123",
				Name: "cmy",
			},
		},
		{
			name: "ignore _id field",
			mock: func(ctx context.Context, opts []*options.FindOneOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindOneWithOptions(ctx, opts).Return(&types.TestUser{
					Name: "cmy",
					Age:  18,
				}, nil).Times(1)
				return mockCollection
			},
			opts: []*options.FindOneOptions{
				{
					Projection: query.BsonBuilder().Add("_id", 0).Build(),
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
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, tc.opts, ctl)

			user, err := finder.FindOneWithOptions(tc.ctx, tc.opts)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_All(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser]
		ctx  context.Context

		want    []*types.TestUser
		wantErr error
	}{
		{
			name: "empty documents",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAll(ctx).Return([]*types.TestUser{}, nil).Times(1)
				return mockCollection
			},
			want: []*types.TestUser{},
		},
		{
			name: "matched",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAll(ctx).Return([]*types.TestUser{
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
				}, nil).Times(1)
				return mockCollection
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
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			users, err := finder.FindAll(tc.ctx)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}

func TestFinder_AllWithOptions(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser]
		ctx  context.Context
		opts []*options.FindOptions

		want    []*types.TestUser
		wantErr error
	}{
		{
			name: "cursor.all error",
			mock: func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAllWithOptions(ctx, opts).Return(nil, errors.New("decode failed")).Times(1)
				return mockCollection
			},
			wantErr: errors.New("decode failed"),
		},
		{
			name: "empty documents",
			mock: func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAllWithOptions(ctx, opts).Return([]*types.TestUser{}, nil).Times(1)
				return mockCollection
			},
			want: []*types.TestUser{},
		},
		{
			name: "matched",
			mock: func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAllWithOptions(ctx, opts).Return([]*types.TestUser{
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
				}, nil).Times(1)
				return mockCollection
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
		},
		{
			name: "returns some of fields",
			mock: func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAllWithOptions(ctx, opts).Return([]*types.TestUser{
					{
						Id:   "123",
						Name: "cmy",
					},
					{
						Id:   "456",
						Name: "cmy",
					},
				}, nil).Times(1)
				return mockCollection
			},
			opts: []*options.FindOptions{
				{
					Projection: query.BsonBuilder().Add("_id", 1).Add("name", 1).Build(),
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
		},
		{
			name: "ignore _id field",
			mock: func(ctx context.Context, opts []*options.FindOptions, ctl *gomock.Controller) iFinder[types.TestUser] {
				mockCollection := mocks.NewMockiFinder[types.TestUser](ctl)
				mockCollection.EXPECT().FindAllWithOptions(ctx, opts).Return([]*types.TestUser{
					{
						Name: "cmy",
						Age:  18,
					},
					{
						Name: "cmy",
						Age:  18,
					},
				}, nil).Times(1)
				return mockCollection
			},
			opts: []*options.FindOptions{
				{
					Projection: query.BsonBuilder().Add("_id", 0).Build(),
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
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, tc.opts, ctl)

			users, err := finder.FindAllWithOptions(tc.ctx, tc.opts)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}
