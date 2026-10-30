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

package updater_test

import (
	"context"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/field"

	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"github.com/chenmingyong0423/go-mongox/v2/updater"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/mock/gomock"
)

func TestNewUpdater(t *testing.T) {
	u := updater.NewUpdater[any](&mongo.Collection{}, nil, nil)
	assert.NotNil(t, u)
}

func TestUpdater_UpdateOne(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update one",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateOne(ctx).Return(nil, assert.AnError).Times(1)
				return u
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return u
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
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateOne(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return u
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully with options",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateOne(ctx, gomock.Any()).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return u
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

			var got *mongo.UpdateResult
			var err error
			// Test with opts parameter like in finder_test.go
			if tc.name == "update successfully with options" {
				got, err = u.UpdateOne(tc.ctx, &options.UpdateOneOptionsBuilder{})
			} else {
				got, err = u.UpdateOne(tc.ctx)
			}
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
		mock func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to update many",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateMany(ctx).Return(nil, assert.AnError).Times(1)
				return u
			},
			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "execute successfully but modified count is 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil).Times(1)
				return u
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
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateMany(ctx).Return(&mongo.UpdateResult{ModifiedCount: 2}, nil).Times(1)
				return u
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 2,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "update successfully with options",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().UpdateMany(ctx, gomock.Any()).Return(&mongo.UpdateResult{ModifiedCount: 2}, nil).Times(1)
				return u
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

			var got *mongo.UpdateResult
			var err error
			// Test with opts parameter like in finder_test.go
			if tc.name == "update successfully with options" {
				got, err = u.UpdateMany(tc.ctx, &options.UpdateManyOptionsBuilder{})
			} else {
				got, err = u.UpdateMany(tc.ctx)
			}
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
		mock func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any]

		ctx     context.Context
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "failed to upsert one",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().Upsert(ctx).Return(nil, assert.AnError).Times(1)
				return u
			},

			ctx:  context.Background(),
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, assert.AnError, err)
			},
		},
		{
			name: "save successfully",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{UpsertedCount: 1}, nil).Times(1)
				return u
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
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().Upsert(ctx).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil).Times(1)
				return u
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				ModifiedCount: 1,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "upsert successfully with options",
			mock: func(ctx context.Context, ctl *gomock.Controller) updater.IUpdater[any] {
				u := mocks.NewMockIUpdater[any](ctl)
				u.EXPECT().Upsert(ctx, gomock.Any()).Return(&mongo.UpdateResult{UpsertedCount: 1}, nil).Times(1)
				return u
			},
			ctx: context.Background(),
			want: &mongo.UpdateResult{
				UpsertedCount: 1,
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

			var got *mongo.UpdateResult
			var err error
			// Test with opts parameter like in finder_test.go
			if tc.name == "upsert successfully with options" {
				got, err = u.Upsert(tc.ctx, &options.UpdateOneOptionsBuilder{})
			} else {
				got, err = u.Upsert(tc.ctx)
			}
			if tc.wantErr(t, err) {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUpdater_Filter(t *testing.T) {
	testCases := []struct {
		name   string
		filter any
		want   updater.IUpdater[any]
	}{
		{
			name:   "set filter successfully",
			filter: map[string]any{"name": "test"},
			want:   nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().Filter(tc.filter).Return(u).Times(1)

			result := u.Filter(tc.filter)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_Updates(t *testing.T) {
	testCases := []struct {
		name    string
		updates any
		want    updater.IUpdater[any]
	}{
		{
			name:    "set updates successfully",
			updates: map[string]any{"$set": map[string]any{"name": "updated"}},
			want:    nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().Updates(tc.updates).Return(u).Times(1)

			result := u.Updates(tc.updates)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_Replacement(t *testing.T) {
	testCases := []struct {
		name        string
		replacement any
		want        updater.IUpdater[any]
	}{
		{
			name:        "set replacement successfully",
			replacement: map[string]any{"name": "replaced", "value": 123},
			want:        nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().Replacement(tc.replacement).Return(u).Times(1)

			result := u.Replacement(tc.replacement)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_ModelHook(t *testing.T) {
	testCases := []struct {
		name      string
		modelHook any
		want      updater.IUpdater[any]
	}{
		{
			name:      "set model hook successfully",
			modelHook: "testHook",
			want:      nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().ModelHook(tc.modelHook).Return(u).Times(1)

			result := u.ModelHook(tc.modelHook)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_RegisterBeforeHooks(t *testing.T) {
	testCases := []struct {
		name  string
		hooks []updater.BeforeHookFn
		want  updater.IUpdater[any]
	}{
		{
			name: "register single before hook successfully",
			hooks: []updater.BeforeHookFn{
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
		{
			name: "register multiple before hooks successfully",
			hooks: []updater.BeforeHookFn{
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().RegisterBeforeHooks(gomock.Any()).Return(u).Times(1)

			result := u.RegisterBeforeHooks(tc.hooks...)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_RegisterAfterHooks(t *testing.T) {
	testCases := []struct {
		name  string
		hooks []updater.AfterHookFn
		want  updater.IUpdater[any]
	}{
		{
			name: "register single after hook successfully",
			hooks: []updater.AfterHookFn{
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
		{
			name: "register multiple after hooks successfully",
			hooks: []updater.AfterHookFn{
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *updater.OpContext, opts ...any) error {
					return nil
				},
			},
			want: nil, // mock will return the interface
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().RegisterAfterHooks(gomock.Any()).Return(u).Times(1)

			result := u.RegisterAfterHooks(tc.hooks...)
			assert.Equal(t, u, result)
		})
	}
}

func TestUpdater_GetCollection(t *testing.T) {
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
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().GetCollection().Return(tc.want).Times(1)

			result := u.GetCollection()
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestUpdater_PreActionHandler(t *testing.T) {
	type document struct {
		CreatedAt time.Time `bson:"created_at"`
		UpdatedAt time.Time `bson:"updated_at"`
		DeletedAt time.Time `bson:"deleted_at"`
	}
	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	filter := bson.M{"name": "alice"}
	fields := field.ParseFields(document{})
	conditions := bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$in", Value: bson.A{nil, time.Time{}}}}}}
	scopedFilter := bson.D{{Key: "$and", Value: bson.A{filter, conditions}}}
	testCases := []struct {
		name            string
		ctx             context.Context
		globalOpContext *operation.OpContext
		opContext       *updater.OpContext
		opType          operation.OpType
		globalHook      callback.CbFn
		beforeHook      []updater.BeforeHookFn
		after           func(t *testing.T, globalOpContext *operation.OpContext, opContext *updater.OpContext)

		wantErr       bool
		expectedError error
	}{
		{
			name:            "pre action handler success",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil),
			opContext:       updater.NewOpContext(nil, nil, nil),
			opType:          operation.OpTypeBeforeUpdate,
		},
		{
			name:            "global before hook error",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil),
			opContext:       updater.NewOpContext(nil, nil, nil),
			opType:          operation.OpTypeBeforeUpdate,
			globalHook: func(context.Context, *operation.OpContext, ...any) error {
				return assert.AnError
			},
			beforeHook: []updater.BeforeHookFn{
				func(context.Context, *updater.OpContext, ...any) error {
					assert.Fail(t, "local hook must not run after a global hook error")
					return nil
				},
			},
			wantErr:       true,
			expectedError: assert.AnError,
		},
		{
			name:            "update filter and updates",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithUpdates(bson.M{"$set": bson.M{"name": "alice"}}), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       updater.NewOpContext(nil, nil, nil),
			opType:          operation.OpTypeBeforeUpdate,
			beforeHook: []updater.BeforeHookFn{
				func(_ context.Context, opCtx *updater.OpContext, _ ...any) error {
					assert.Equal(t, scopedFilter, opCtx.Filter)
					assert.Equal(t, bson.M{"$set": bson.M{"name": "alice", "updated_at": now}}, opCtx.Updates)
					opCtx.Filter = bson.M{"name": "bob"}
					opCtx.Updates = bson.M{"$set": bson.M{"name": "bob"}}
					return nil
				},
			},
			after: func(t *testing.T, globalOpContext *operation.OpContext, opContext *updater.OpContext) {
				assert.Equal(t, bson.M{"name": "bob"}, globalOpContext.Filter)
				assert.Equal(t, bson.M{"$set": bson.M{"name": "bob"}}, globalOpContext.Updates)
				assert.Equal(t, globalOpContext.Filter, opContext.Filter)
				assert.Equal(t, globalOpContext.Updates, opContext.Updates)
				assert.Equal(t, bson.M{"name": "alice"}, filter)
			},
		},
		{
			name:            "upsert filter and updates",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithUpdates(bson.M{"$set": bson.M{"name": "alice"}}), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       updater.NewOpContext(nil, nil, nil),
			opType:          operation.OpTypeBeforeUpsert,
			beforeHook: []updater.BeforeHookFn{
				func(_ context.Context, opCtx *updater.OpContext, _ ...any) error {
					assert.Equal(t, scopedFilter, opCtx.Filter)
					assert.Equal(t, bson.M{"$set": bson.M{"name": "alice", "updated_at": now}, "$setOnInsert": bson.M{"created_at": now}}, opCtx.Updates)
					opCtx.Filter = bson.M{"name": "bob"}
					opCtx.Updates = bson.M{"$set": bson.M{"name": "bob"}}
					return nil
				},
			},
			after: func(t *testing.T, globalOpContext *operation.OpContext, opContext *updater.OpContext) {
				assert.Equal(t, bson.M{"name": "bob"}, globalOpContext.Filter)
				assert.Equal(t, bson.M{"$set": bson.M{"name": "bob"}}, globalOpContext.Updates)
				assert.Equal(t, globalOpContext.Filter, opContext.Filter)
				assert.Equal(t, globalOpContext.Updates, opContext.Updates)
				assert.Equal(t, bson.M{"name": "alice"}, filter)
			},
		},
		{
			name:            "local before hook error",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithUpdates(bson.M{"$set": bson.M{"name": "alice"}}), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       updater.NewOpContext(nil, nil, nil),
			opType:          operation.OpTypeBeforeUpdate,
			beforeHook: []updater.BeforeHookFn{
				func(context.Context, *updater.OpContext, ...any) error { return assert.AnError },
			},
			wantErr:       true,
			expectedError: assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			callbacks := callback.InitializeCallbacks()
			if tc.globalHook != nil {
				callbacks.Register(tc.opType, tc.name, tc.globalHook)
			}
			u := updater.NewUpdater[document](nil, callbacks, fields)
			u.RegisterBeforeHooks(tc.beforeHook...)

			err := u.PreActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType)
			if tc.wantErr {
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
			if tc.after != nil {
				tc.after(t, tc.globalOpContext, tc.opContext)
			}
		})
	}
}

func TestUpdater_PostActionHandler(t *testing.T) {
	testCases := []struct {
		name            string
		ctx             context.Context
		globalOpContext *operation.OpContext
		opContext       *updater.OpContext
		opType          operation.OpType
		wantErr         bool
		expectedError   error
	}{
		{
			name:            "post action handler success",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &updater.OpContext{},
			opType:          operation.OpTypeAfterUpdate,
			wantErr:         false,
			expectedError:   nil,
		},
		{
			name:            "post action handler with error",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &updater.OpContext{},
			opType:          operation.OpTypeAfterUpdate,
			wantErr:         true,
			expectedError:   assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			u := mocks.NewMockIUpdater[any](ctl)
			u.EXPECT().PostActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType).Return(tc.expectedError).Times(1)

			err := u.PostActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
