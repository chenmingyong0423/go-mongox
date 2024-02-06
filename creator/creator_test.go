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

package creator

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/types"

	mocks "github.com/chenmingyong0423/go-mongox/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestNewCreator(t *testing.T) {
	mongoCollection := &mongo.Collection{}
	creator := NewCreator[any](mongoCollection)

	assert.NotNil(t, creator)
	assert.Equal(t, mongoCollection, creator.collection)
}

func TestCreator_One(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller, doc *types.TestUser) iCreator[types.TestUser]
		ctx  context.Context
		doc  *types.TestUser

		wantErr error
	}{
		{
			name: "nil doc",
			mock: func(ctx context.Context, ctl *gomock.Controller, doc *types.TestUser) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, doc).Return(nil, errors.New("nil filter")).Times(1)
				return mockCollection
			},
			ctx:     context.Background(),
			doc:     nil,
			wantErr: errors.New("nil filter"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller, doc *types.TestUser) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, doc).Return(&mongo.InsertOneResult{InsertedID: "?"}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: &types.TestUser{
				Name: "chenmingyong",
				Age:  24,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			creator := tc.mock(tc.ctx, ctl, tc.doc)

			insertOneResult, err := creator.InsertOne(tc.ctx, tc.doc)
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
		mock func(ctx context.Context, ctl *gomock.Controller, docs []*types.TestUser) iCreator[types.TestUser]
		ctx  context.Context
		docs []*types.TestUser

		wantIdsLength int
		wantErr       error
	}{
		{
			name: "nil docs",
			mock: func(ctx context.Context, ctl *gomock.Controller, docs []*types.TestUser) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, docs).Return(nil, errors.New("nil docs")).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			docs: []*types.TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "burt", Age: 25},
			},
			wantErr: errors.New("nil docs"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller, docs []*types.TestUser) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, docs).Return(&mongo.InsertManyResult{InsertedIDs: make([]interface{}, 2)}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			docs: []*types.TestUser{
				{Name: "chenmingyong", Age: 24},
				{Name: "burt", Age: 25},
			},
			wantIdsLength: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			creator := tc.mock(tc.ctx, ctl, tc.docs)

			insertResult, err := creator.InsertMany(tc.ctx, tc.docs)
			require.Equal(t, tc.wantErr, err)
			if err == nil {
				assert.Equal(t, tc.wantIdsLength, len(insertResult.InsertedIDs))
			}
		})
	}
}
