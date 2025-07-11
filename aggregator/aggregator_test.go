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
	"time"

	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
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

type TestTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type IllegalUser struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
	Age  string
}

func TestAggregator_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}
	aggregator := NewAggregator[any](mongoCollection, nil, nil)

	assert.NotNil(t, aggregator)
	assert.Equal(t, mongoCollection, aggregator.collection)
}

func TestAggregator_Aggregation(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctx context.Context, ctl *gomock.Controller) IAggregator[TestUser]
		ctx     context.Context
		opts    []options.Lister[options.AggregateOptions]
		want    []*TestUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().Aggregate(ctx).Return(nil, errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:     context.Background(),
			wantErr: assert.Error,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().Aggregate(ctx).Return([]*TestUser{
					{Name: "chenmingyong", Age: 24},
					{Name: "gopher", Age: 25},
				}, nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			want: []*TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "gopher", Age: 25},
			},
			wantErr: assert.NoError,
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().Aggregate(ctx, gomock.Any()).Return([]*TestUser{
					{Name: "chenmingyong", Age: 24},
				}, nil).Times(1)
				return aggregator
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.AggregateOptions]{options.Aggregate().SetComment("test aggregation")},
			want: []*TestUser{
				{Name: "chenmingyong", Age: 24},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			aggregator := tc.mock(tc.ctx, ctl)

			var result []*TestUser
			var err error
			if tc.opts != nil {
				result, err = aggregator.Aggregate(tc.ctx, tc.opts...)
			} else {
				result, err = aggregator.Aggregate(tc.ctx)
			}
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
		mock    func(ctx context.Context, ctl *gomock.Controller, result any) IAggregator[TestUser]
		ctx     context.Context
		opts    []options.Lister[options.AggregateOptions]
		result  []*User
		want    []*User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "got error",
			mock: func(ctx context.Context, ctl *gomock.Controller, result any) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().AggregateWithParse(ctx, result).Return(errors.New("can only marshal slices and arrays into aggregation pipelines, but got invalid")).Times(1)
				return aggregator
			},
			ctx:     context.Background(),
			result:  []*User{},
			wantErr: assert.Error,
		},
		{
			name: "got result",
			mock: func(ctx context.Context, ctl *gomock.Controller, result any) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().AggregateWithParse(ctx, result).Return(nil).Times(1)
				return aggregator
			},
			ctx: context.Background(),
			result: []*User{
				{Id: "1", Name: "chenmingyong", Age: 24, IsProgrammer: true},
			},
			want: []*User{
				{Id: "1", Name: "chenmingyong", Age: 24, IsProgrammer: true},
			},
			wantErr: assert.NoError,
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller, result any) IAggregator[TestUser] {
				aggregator := mocks.NewMockIAggregator[TestUser](ctl)
				aggregator.EXPECT().AggregateWithParse(ctx, result, gomock.Any()).Return(nil).Times(1)
				return aggregator
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.AggregateOptions]{options.Aggregate().SetComment("test aggregation")},
			result: []*User{
				{Id: "1", Name: "chenmingyong", Age: 25, IsProgrammer: false},
			},
			want: []*User{
				{Id: "1", Name: "chenmingyong", Age: 25, IsProgrammer: false},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			aggregator := tc.mock(tc.ctx, ctl, tc.result)
			var err error
			if tc.opts != nil {
				err = aggregator.AggregateWithParse(tc.ctx, tc.result, tc.opts...)
			} else {
				err = aggregator.AggregateWithParse(tc.ctx, tc.result)
			}
			if tc.wantErr(t, err) {
				assert.ElementsMatch(t, tc.want, tc.result)
			}
		})
	}
}

func TestAggregator_CorrectOpTypes(t *testing.T) {
	t.Run("aggregation should use correct OpTypes", func(t *testing.T) {
		// Setup
		col := &mongo.Collection{}
		dbCallbacks := callback.InitializeCallbacks()

		// Track which hooks are called
		var calledHooks []string

		// Register hooks for different operation types
		dbCallbacks.Register(operation.OpTypeBeforeInsert, "insert-before", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			calledHooks = append(calledHooks, "insert-before")
			return nil
		})

		dbCallbacks.Register(operation.OpTypeAfterInsert, "insert-after", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			calledHooks = append(calledHooks, "insert-after")
			return nil
		})

		dbCallbacks.Register(operation.OpTypeBeforeAggregate, "aggregate-before", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			calledHooks = append(calledHooks, "aggregate-before")
			return nil
		})

		dbCallbacks.Register(operation.OpTypeAfterAggregate, "aggregate-after", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			calledHooks = append(calledHooks, "aggregate-after")
			return nil
		})

		// Create aggregator
		aggregator := NewAggregator[map[string]interface{}](col, dbCallbacks, nil)
		pipeline := []interface{}{map[string]interface{}{"$match": map[string]interface{}{}}}

		// Test preActionHandler and postActionHandler directly
		ctx := context.Background()
		globalOpContext := operation.NewOpContext(col)
		opContext := NewOpContext(col, pipeline)

		// Test before handler
		err := aggregator.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeAggregate)
		assert.NoError(t, err)

		// Test after handler
		err = aggregator.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterAggregate)
		assert.NoError(t, err)

		// Verify only aggregation hooks were called
		assert.Contains(t, calledHooks, "aggregate-before")
		assert.Contains(t, calledHooks, "aggregate-after")
		assert.NotContains(t, calledHooks, "insert-before")
		assert.NotContains(t, calledHooks, "insert-after")
	})

	t.Run("insert hooks should not be triggered by aggregation", func(t *testing.T) {
		// Setup
		col := &mongo.Collection{}
		dbCallbacks := callback.InitializeCallbacks()

		// Register an insert hook that should NOT be called
		insertHookCalled := false
		dbCallbacks.Register(operation.OpTypeBeforeInsert, "should-not-be-called", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			insertHookCalled = true
			return errors.New("insert hook incorrectly called during aggregation")
		})

		// Create aggregator
		aggregator := NewAggregator[map[string]interface{}](col, dbCallbacks, nil)
		pipeline := []interface{}{map[string]interface{}{"$match": map[string]interface{}{}}}

		// Test that insert hooks are NOT called
		ctx := context.Background()
		globalOpContext := operation.NewOpContext(col)
		opContext := NewOpContext(col, pipeline)

		// This should NOT trigger insert hooks
		err := aggregator.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeAggregate)
		assert.NoError(t, err)
		assert.False(t, insertHookCalled, "Insert hook should not be called during aggregation")
	})
}
