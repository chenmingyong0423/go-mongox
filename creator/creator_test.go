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

package creator_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/creator"
	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewCreator(t *testing.T) {
	mongoCollection := &mongo.Collection{}
	c := creator.NewCreator[any](mongoCollection, nil, nil)

	assert.NotNil(t, c)
	assert.Equal(t, mongoCollection, c.GetCollection())
}

func TestCreator_One(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller, doc *TestUser) creator.ICreator[TestUser]
		ctx  context.Context
		doc  *TestUser
		opts []options.Lister[options.InsertOneOptions]

		wantErr error
	}{
		{
			name: "nil doc",
			mock: func(ctx context.Context, ctl *gomock.Controller, doc *TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, doc).Return(nil, errors.New("nil filter")).Times(1)
				return mockCollection
			},
			ctx:     context.Background(),
			doc:     nil,
			wantErr: errors.New("nil filter"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller, doc *TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, doc).Return(&mongo.InsertOneResult{InsertedID: "?"}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: &TestUser{
				Name: "chenmingyong",
				Age:  24,
			},
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller, doc *TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, doc, gomock.Any()).Return(&mongo.InsertOneResult{InsertedID: "with_opts"}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: &TestUser{
				Name: "chenmingyong",
				Age:  24,
			},
			opts: []options.Lister[options.InsertOneOptions]{options.InsertOne().SetComment("test")},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(tc.ctx, ctl, tc.doc)

			insertOneResult, err := c.InsertOne(tc.ctx, tc.doc, tc.opts...)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				assert.NotNil(t, insertOneResult.InsertedID)
			}
		})
	}
}

func TestCreator_Many(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller, docs []*TestUser) creator.ICreator[TestUser]
		ctx  context.Context
		docs []*TestUser
		opts []options.Lister[options.InsertManyOptions]

		wantIdsLength int
		wantErr       error
	}{
		{
			name: "nil docs",
			mock: func(ctx context.Context, ctl *gomock.Controller, docs []*TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, docs).Return(nil, errors.New("nil docs")).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			docs: []*TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "burt", Age: 25},
			},
			wantErr: errors.New("nil docs"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller, docs []*TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, docs).Return(&mongo.InsertManyResult{InsertedIDs: make([]interface{}, 2)}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			docs: []*TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "burt", Age: 25},
			},
			wantIdsLength: 2,
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller, docs []*TestUser) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, docs, gomock.Any()).Return(&mongo.InsertManyResult{InsertedIDs: []interface{}{"1", "2"}}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			docs: []*TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "burt", Age: 25},
			},
			opts:          []options.Lister[options.InsertManyOptions]{options.InsertMany().SetOrdered(false)},
			wantIdsLength: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(tc.ctx, ctl, tc.docs)

			insertResult, err := c.InsertMany(tc.ctx, tc.docs, tc.opts...)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				assert.Equal(t, tc.wantIdsLength, len(insertResult.InsertedIDs))
			}
		})
	}
}

func TestCreator_GetCollection(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) creator.ICreator[TestUser]
	}{
		{
			name: "get collection",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCollection := &mongo.Collection{}
				mockCollection.EXPECT().GetCollection().Return(expectedCollection).Times(1)
				return mockCollection
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(ctl)

			result := c.GetCollection()
			assert.NotNil(t, result)
		})
	}
}

func TestCreator_ModelHook(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) creator.ICreator[TestUser]

		modelHook any
	}{
		{
			name: "set model hook",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCreator := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().ModelHook(&TestUser{}).Return(expectedCreator).Times(1)
				return mockCollection
			},
			modelHook: &TestUser{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(ctl)

			result := c.ModelHook(tc.modelHook)
			assert.NotNil(t, result)
		})
	}
}

func TestCreator_RegisterBeforeHooks(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) creator.ICreator[TestUser]

		hooks []creator.HookFn[TestUser]
	}{
		{
			name: "register before hooks",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCreator := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().RegisterBeforeHooks(gomock.Any()).Return(expectedCreator).Times(1)
				return mockCollection
			},
			hooks: []creator.HookFn[TestUser]{
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
		{
			name: "register multiple before hooks - should trigger hooks loop",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCreator := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().RegisterBeforeHooks(gomock.Any(), gomock.Any()).Return(expectedCreator).Times(1)
				return mockCollection
			},
			hooks: []creator.HookFn[TestUser]{
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(ctl)

			result := c.RegisterBeforeHooks(tc.hooks...)
			assert.NotNil(t, result)
		})
	}
}

func TestCreator_RegisterAfterHooks(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) creator.ICreator[TestUser]

		hooks []creator.HookFn[TestUser]
	}{
		{
			name: "register after hooks",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCreator := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().RegisterAfterHooks(gomock.Any()).Return(expectedCreator).Times(1)
				return mockCollection
			},
			hooks: []creator.HookFn[TestUser]{
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
		{
			name: "register multiple after hooks - should trigger hooks loop",
			mock: func(ctl *gomock.Controller) creator.ICreator[TestUser] {
				mockCollection := mocks.NewMockICreator[TestUser](ctl)
				expectedCreator := mocks.NewMockICreator[TestUser](ctl)
				mockCollection.EXPECT().RegisterAfterHooks(gomock.Any(), gomock.Any()).Return(expectedCreator).Times(1)
				return mockCollection
			},
			hooks: []creator.HookFn[TestUser]{
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
				func(ctx context.Context, opCtx *creator.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			c := tc.mock(ctl)

			result := c.RegisterAfterHooks(tc.hooks...)
			assert.NotNil(t, result)
		})
	}
}
