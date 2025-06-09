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

package deleter_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/deleter"
	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"
	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/mock/gomock"
)

type TestUser struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (tu *TestUser) DefaultCreatedAt() {
	if tu.CreatedAt.IsZero() {
		tu.CreatedAt = time.Now().Local()
	}
}

func (tu *TestUser) DefaultUpdatedAt() {
	tu.UpdatedAt = time.Now().Local()
}

func TestDeleter_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}

	result := deleter.NewDeleter[any](mongoCollection, nil, nil)
	assert.NotNil(t, result, "Expected non-nil Deleter")
	assert.Equal(t, mongoCollection, result.GetCollection(), "Expected deleter field to be initialized correctly")
}

func TestDeleter_DeleteOne(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser]
		ctx  context.Context

		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
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
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
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
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
				mockCollection.EXPECT().DeleteOne(ctx).Return(&mongo.DeleteResult{DeletedCount: 1}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "delete success with options",
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
				mockCollection.EXPECT().DeleteOne(ctx, gomock.Any()).Return(&mongo.DeleteResult{DeletedCount: 1}, nil).Times(1)
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
			d := tc.mock(tc.ctx, ctl)

			var got *mongo.DeleteResult
			var err error
			if tc.name == "delete success with options" {
				got, err = d.DeleteOne(tc.ctx, &options.DeleteOneOptionsBuilder{})
			} else {
				got, err = d.DeleteOne(tc.ctx)
			}
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

		mock func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser]
		ctx  context.Context

		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
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
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
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
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
				mockCollection.EXPECT().DeleteMany(ctx).Return(&mongo.DeleteResult{DeletedCount: 2}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			want: &mongo.DeleteResult{
				DeletedCount: 2,
			},
			wantErr: assert.NoError,
		},
		{
			name: "delete success with options",
			mock: func(ctx context.Context, ctl *gomock.Controller) deleter.IDeleter[TestUser] {
				mockCollection := mocks.NewMockIDeleter[TestUser](ctl)
				mockCollection.EXPECT().DeleteMany(ctx, gomock.Any()).Return(&mongo.DeleteResult{DeletedCount: 2}, nil).Times(1)
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
			d := tc.mock(tc.ctx, ctl)

			var got *mongo.DeleteResult
			var err error
			if tc.name == "delete success with options" {
				got, err = d.DeleteMany(tc.ctx, &options.DeleteManyOptionsBuilder{})
			} else {
				got, err = d.DeleteMany(tc.ctx)
			}
			if !tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDeleter_Filter(t *testing.T) {
	testCases := []struct {
		name   string
		filter any
		want   deleter.IDeleter[TestUser]
	}{
		{
			name:   "set filter successfully",
			filter: bson.M{"name": "test"},
			want:   nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().Filter(tc.filter).Return(d).Times(1)

			result := d.Filter(tc.filter)
			assert.Equal(t, d, result)
		})
	}
}

func TestDeleter_ModelHook(t *testing.T) {
	testCases := []struct {
		name      string
		modelHook any
		want      deleter.IDeleter[TestUser]
	}{
		{
			name:      "set model hook successfully",
			modelHook: &TestUser{Name: "hook"},
			want:      nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().ModelHook(tc.modelHook).Return(d).Times(1)

			result := d.ModelHook(tc.modelHook)
			assert.Equal(t, d, result)
		})
	}
}

func TestDeleter_RegisterBeforeHooks(t *testing.T) {
	testCases := []struct {
		name  string
		hooks []deleter.BeforeHookFn
		want  deleter.IDeleter[TestUser]
	}{
		{
			name: "register single before hook successfully",
			hooks: []deleter.BeforeHookFn{
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
		{
			name: "register multiple before hooks successfully",
			hooks: []deleter.BeforeHookFn{
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().RegisterBeforeHooks(gomock.Any()).Return(d).Times(1)

			result := d.RegisterBeforeHooks(tc.hooks...)
			assert.Equal(t, d, result)
		})
	}
}

func TestDeleter_RegisterAfterHooks(t *testing.T) {
	testCases := []struct {
		name  string
		hooks []deleter.AfterHookFn
		want  deleter.IDeleter[TestUser]
	}{
		{
			name: "register single after hook successfully",
			hooks: []deleter.AfterHookFn{
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
		{
			name: "register multiple after hooks successfully",
			hooks: []deleter.AfterHookFn{
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *deleter.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().RegisterAfterHooks(gomock.Any()).Return(d).Times(1)

			result := d.RegisterAfterHooks(tc.hooks...)
			assert.Equal(t, d, result)
		})
	}
}

func TestDeleter_GetCollection(t *testing.T) {
	testCases := []struct {
		name string
		want *mongo.Collection
	}{
		{
			name: "get collection successfully",
			want: &mongo.Collection{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().GetCollection().Return(tc.want).Times(1)

			result := d.GetCollection()
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestDeleter_PreActionHandler(t *testing.T) {
	testCases := []struct {
		name            string
		ctx             context.Context
		globalOpContext *operation.OpContext
		opContext       *deleter.OpContext
		opType          operation.OpType
		wantErr         bool
		expectedError   error
	}{
		{
			name:            "pre action handler success",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			wantErr:         false,
			expectedError:   nil,
		},
		{
			name:            "pre action handler with error",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			wantErr:         true,
			expectedError:   assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().PreActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType).Return(tc.expectedError).Times(1)

			err := d.PreActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleter_PostActionHandler(t *testing.T) {
	testCases := []struct {
		name            string
		ctx             context.Context
		globalOpContext *operation.OpContext
		opContext       *deleter.OpContext
		opType          operation.OpType
		wantErr         bool
		expectedError   error
	}{
		{
			name:            "post action handler success",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeAfterDelete,
			wantErr:         false,
			expectedError:   nil,
		},
		{
			name:            "post action handler with error",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeAfterDelete,
			wantErr:         true,
			expectedError:   assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			d := mocks.NewMockIDeleter[TestUser](ctl)
			d.EXPECT().PostActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType).Return(tc.expectedError).Times(1)

			err := d.PostActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
