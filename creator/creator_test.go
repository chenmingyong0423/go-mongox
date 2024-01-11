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
		mock func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser]
		ctx  context.Context
		doc  *types.TestUser

		wantId  string
		wantErr error
	}{
		{
			name: "duplicate",
			mock: func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				}).Return(nil, errors.New("duplicate key")).Times(1)
				return mockCollection
			},
			ctx:     context.Background(),
			doc:     &types.TestUser{Id: "123", Name: "cmy", Age: 18},
			wantId:  "",
			wantErr: errors.New("duplicate key"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertOne(ctx, &types.TestUser{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				}).Return(&mongo.InsertOneResult{InsertedID: "123"}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
			wantId: "123",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			creator := tc.mock(tc.ctx, ctl)

			insertOneResult, err := creator.InsertOne(tc.ctx, tc.doc)
			assert.Equal(t, tc.wantErr, err)
			if err == nil {
				assert.Equal(t, tc.wantId, insertOneResult.InsertedID)
			}
		})
	}
}

func TestCreator_Many(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser]
		ctx  context.Context
		doc  []*types.TestUser

		wantIds []string
		wantErr error
	}{
		{
			name: "duplicate",
			mock: func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, []*types.TestUser{
					{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					}, {
						Id:   "123",
						Name: "cmy",
						Age:  18,
					},
				}).Return(nil, errors.New("duplicate key")).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: []*types.TestUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantIds: nil,
			wantErr: errors.New("duplicate key"),
		},
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) iCreator[types.TestUser] {
				mockCollection := mocks.NewMockiCreator[types.TestUser](ctl)
				mockCollection.EXPECT().InsertMany(ctx, []*types.TestUser{
					{
						Id:   "123",
						Name: "cmy",
						Age:  18,
					}, {
						Id:   "456",
						Name: "cmy",
						Age:  18,
					},
				}).Return(&mongo.InsertManyResult{InsertedIDs: []any{"123", "456"}}, nil).Times(1)
				return mockCollection
			},
			ctx: context.Background(),
			doc: []*types.TestUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "456", Name: "cmy", Age: 18},
			},
			wantIds: []string{"123", "456"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			creator := tc.mock(tc.ctx, ctl)

			insertResult, err := creator.InsertMany(tc.ctx, tc.doc)
			assert.Equal(t, tc.wantErr, err)
			if err == nil {
				assert.ElementsMatch(t, tc.wantIds, insertResult.InsertedIDs)
			}
		})
	}
}
