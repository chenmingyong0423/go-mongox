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

	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		mock func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser]
		ctx  context.Context

		want    *TestUser
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any()).Return(nil, mongo.ErrNoDocuments).Times(1)
				return mockCollection
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any()).Return(&TestUser{
					Name: "chenmingyong",
					Age:  24,
				}, nil).Times(1)
				return mockCollection
			},
			want: &TestUser{
				Name: "chenmingyong",
				Age:  24,
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

func TestFinder_All(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser]
		ctx  context.Context

		want    []*TestUser
		wantErr error
	}{
		{
			name: "empty documents",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().Find(ctx).Return([]*TestUser{}, nil).Times(1)
				return mockCollection
			},
			want: []*TestUser{},
		},
		{
			name: "matched",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().Find(ctx).Return([]*TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
					{
						Name: "burt",
						Age:  25,
					},
				}, nil).Times(1)
				return mockCollection
			},
			want: []*TestUser{
				{
					Name: "chenmingyong",
					Age:  24,
				},
				{
					Name: "burt",
					Age:  25,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			users, err := finder.Find(tc.ctx)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}

func TestFinder_Count(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser]
		ctx  context.Context

		want    int64
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(0), errors.New("nil filter error")).Times(1)
				return mockCollection
			},
			want:    0,
			wantErr: errors.New("nil filter error"),
		},
		{
			name: "matched 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(0), nil).Times(1)
				return mockCollection
			},
			want: 0,
		},
		{
			name: "matched 1",
			mock: func(ctx context.Context, ctl *gomock.Controller) iFinder[TestUser] {
				mockCollection := mocks.NewMockiFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(1), nil).Times(1)
				return mockCollection
			},
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			users, err := finder.Count(tc.ctx)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}
