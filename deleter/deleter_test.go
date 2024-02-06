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

package deleter

import (
	"context"
	"errors"
	"testing"

	"github.com/chenmingyong0423/go-mongox/types"

	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestDeleter_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}

	result := NewDeleter[any](mongoCollection)
	assert.NotNil(t, result, "Expected non-nil Deleter")
	assert.Equal(t, mongoCollection, result.collection, "Expected deleter field to be initialized correctly")
}

func TestDeleter_DeleteOne(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser]
		ctx  context.Context

		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteOne(ctx).Return(nil, errors.New("nil filter")).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("nil filter"), err)
			},
		},
		{
			name: "deleted count: 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteOne(ctx).Return(&mongo.DeleteResult{DeletedCount: 0}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "delete success",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteOne(ctx).Return(&mongo.DeleteResult{DeletedCount: 1}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			deleter := tc.mock(tc.ctx, ctl)

			got, err := deleter.DeleteOne(tc.ctx)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}

}

func TestDeleter_DeleteMany(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser]
		ctx  context.Context

		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteMany(ctx).Return(nil, errors.New("nil filter")).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("nil filter"), err)
			},
		},
		{
			name: "deleted count: 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteMany(ctx).Return(&mongo.DeleteResult{DeletedCount: 0}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "delete success",
			mock: func(ctx context.Context, ctl *gomock.Controller) iDeleter[types.TestUser] {
				mockCollection := mocks.NewMockiDeleter[types.TestUser](ctl)
				mockCollection.EXPECT().DeleteMany(ctx).Return(&mongo.DeleteResult{DeletedCount: 2}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 2,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			deleter := tc.mock(tc.ctx, ctl)

			got, err := deleter.DeleteMany(tc.ctx)
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
