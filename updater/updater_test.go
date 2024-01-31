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

package updater

import (
	"context"
	"testing"

	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestNewUpdater(t *testing.T) {
	updater := NewUpdater[any](&mongo.Collection{})
	assert.NotNil(t, updater)
}

func TestUpdater_UpdateOne(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update one",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.UpdateOne(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_UpdateMany(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update many",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},
			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 0,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 2}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 2,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.UpdateMany(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_Upsert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to upsert one",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(nil, assert.AnError).Times(1)
				return updater
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "save successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{UpsertedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				UpsertedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) iUpdater[any] {
				updater := mocks.NewMockiUpdater[any](ctl)
				updater.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return updater
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := tc.mock(context.Background(), ctl)
			got, err := u.Upsert(tc.ctx)
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
