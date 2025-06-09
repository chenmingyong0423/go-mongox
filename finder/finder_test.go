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

package finder_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/finder"
	mocks "github.com/chenmingyong0423/go-mongox/v2/mock"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestFinder_New(t *testing.T) {
	mongoCollection := &mongo.Collection{}

	result := finder.NewFinder[any](mongoCollection, nil, nil)
	assert.NotNil(t, result, "Expected non-nil Finder")
	assert.Equal(t, mongoCollection, result.GetCollection(), "Expected finder field to be initialized correctly")
}

func TestFinder_One(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.FindOneOptions]

		want    *TestUser
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any()).Return(nil, mongo.ErrNoDocuments).Times(1)
				return mockCollection
			},
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
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
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(&TestUser{
					Name: "chenmingyong",
				}, nil).Times(1)
				return mockCollection
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.FindOneOptions]{options.FindOne().SetProjection(bson.M{"age": 0})},
			want: &TestUser{
				Name: "chenmingyong",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			var user *TestUser
			var err error
			if tc.opts != nil {
				user, err = finder.FindOne(tc.ctx, tc.opts...)
			} else {
				user, err = finder.FindOne(tc.ctx)
			}
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_All(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.FindOptions]

		want    []*TestUser
		wantErr error
	}{
		{
			name: "empty documents",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Find(ctx).Return([]*TestUser{}, nil).Times(1)
				return mockCollection
			},
			want: []*TestUser{},
		},
		{
			name: "matched",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
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
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Find(ctx, gomock.Any()).Return([]*TestUser{
					{
						Name: "chenmingyong",
						Age:  24,
					},
				}, nil).Times(1)
				return mockCollection
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.FindOptions]{options.Find().SetLimit(1)},
			want: []*TestUser{
				{
					Name: "chenmingyong",
					Age:  24,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			var users []*TestUser
			var err error
			if tc.opts != nil {
				users, err = finder.Find(tc.ctx, tc.opts...)
			} else {
				users, err = finder.Find(tc.ctx)
			}
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}

func TestFinder_Count(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.CountOptions]

		want    int64
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(0), errors.New("nil filter error")).Times(1)
				return mockCollection
			},
			want:    0,
			wantErr: errors.New("nil filter error"),
		},
		{
			name: "matched 0",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(0), nil).Times(1)
				return mockCollection
			},
			want: 0,
		},
		{
			name: "matched 1",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx).Return(int64(1), nil).Times(1)
				return mockCollection
			},
			want: 1,
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Count(ctx, gomock.Any()).Return(int64(2), nil).Times(1)
				return mockCollection
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.CountOptions]{options.Count().SetComment("test")},
			want: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			var users int64
			var err error
			if tc.opts != nil {
				users, err = finder.Count(tc.ctx, tc.opts...)
			} else {
				users, err = finder.Count(tc.ctx)
			}
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, users)
		})
	}
}

func TestFindOneAndUpdate(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.FindOneAndUpdateOptions]

		want    *TestUser
		wantErr error
	}{
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOneAndUpdate(ctx).Return(nil, mongo.ErrNoDocuments).Times(1)
				return mockCollection
			},

			ctx:     context.Background(),
			want:    nil,
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "match the first one and update",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOneAndUpdate(ctx).Return(&TestUser{Name: "hejiangda", Age: 18}, nil).Times(1)
				return mockCollection
			},
			ctx:  context.Background(),
			want: &TestUser{Name: "hejiangda", Age: 18},
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().FindOneAndUpdate(ctx, gomock.Any()).Return(&TestUser{Name: "updated", Age: 30}, nil).Times(1)
				return mockCollection
			},
			ctx:  context.Background(),
			opts: []options.Lister[options.FindOneAndUpdateOptions]{options.FindOneAndUpdate().SetReturnDocument(options.After)},
			want: &TestUser{Name: "updated", Age: 30},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			user, err := finder.FindOneAndUpdate(tc.ctx, tc.opts...)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, user)
		})
	}
}

func TestFinder_Distinct(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.DistinctOptions]

		fieldName string
		want      *mongo.DistinctResult
		wantErr   error
	}{
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedResult := &mongo.DistinctResult{}
				mockCollection.EXPECT().Distinct(ctx, "name").Return(expectedResult).Times(1)
				return mockCollection
			},
			ctx:       context.Background(),
			fieldName: "name",
			want:      &mongo.DistinctResult{},
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedResult := &mongo.DistinctResult{}
				mockCollection.EXPECT().Distinct(ctx, "name", gomock.Any()).Return(expectedResult).Times(1)
				return mockCollection
			},
			ctx:       context.Background(),
			fieldName: "name",
			opts:      []options.Lister[options.DistinctOptions]{options.Distinct().SetComment("test")},
			want:      &mongo.DistinctResult{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			var result *mongo.DistinctResult
			if tc.opts != nil {
				result = finder.Distinct(tc.ctx, tc.fieldName, tc.opts...)
			} else {
				result = finder.Distinct(tc.ctx, tc.fieldName)
			}
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestFinder_DistinctWithParse(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context
		opts []options.Lister[options.DistinctOptions]

		fieldName string
		wantErr   error
	}{
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().DistinctWithParse(ctx, "name", gomock.Any()).Return(nil).Times(1)
				return mockCollection
			},
			ctx:       context.Background(),
			fieldName: "name",
			wantErr:   nil,
		},
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().DistinctWithParse(ctx, "name", gomock.Any()).Return(errors.New("distinct error")).Times(1)
				return mockCollection
			},
			ctx:       context.Background(),
			fieldName: "name",
			wantErr:   errors.New("distinct error"),
		},
		{
			name: "with options - should trigger opts loop",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().DistinctWithParse(ctx, "name", gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return mockCollection
			},
			ctx:       context.Background(),
			fieldName: "name",
			opts:      []options.Lister[options.DistinctOptions]{options.Distinct().SetComment("test")},
			wantErr:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			var names []string
			var err error
			if tc.opts != nil {
				err = finder.DistinctWithParse(tc.ctx, tc.fieldName, &names, tc.opts...)
			} else {
				err = finder.DistinctWithParse(tc.ctx, tc.fieldName, &names)
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestFinder_Filter(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		filter any
		want   finder.IFinder[TestUser]
	}{
		{
			name: "set filter",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Filter(bson.M{"name": "test"}).Return(expectedFinder).Times(1)
				return mockCollection
			},
			filter: bson.M{"name": "test"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.Filter(tc.filter)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_Limit(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		limit int64
	}{
		{
			name: "set limit",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Limit(int64(10)).Return(expectedFinder).Times(1)
				return mockCollection
			},
			limit: 10,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.Limit(tc.limit)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_Skip(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		skip int64
	}{
		{
			name: "set skip",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Skip(int64(5)).Return(expectedFinder).Times(1)
				return mockCollection
			},
			skip: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.Skip(tc.skip)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_Sort(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		sort any
	}{
		{
			name: "set sort",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Sort(bson.M{"age": 1}).Return(expectedFinder).Times(1)
				return mockCollection
			},
			sort: bson.M{"age": 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.Sort(tc.sort)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_Updates(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		update any
	}{
		{
			name: "set updates",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().Updates(bson.M{"$set": bson.M{"age": 25}}).Return(expectedFinder).Times(1)
				return mockCollection
			},
			update: bson.M{"$set": bson.M{"age": 25}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.Updates(tc.update)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_ModelHook(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		modelHook any
	}{
		{
			name: "set model hook",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				hookObj := &TestUser{Name: "hook"}
				mockCollection.EXPECT().ModelHook(hookObj).Return(expectedFinder).Times(1)
				return mockCollection
			},
			modelHook: &TestUser{Name: "hook"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.ModelHook(tc.modelHook)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_RegisterBeforeHooks(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		hooks []finder.BeforeHookFn[TestUser]
	}{
		{
			name: "register before hooks", mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().RegisterBeforeHooks(gomock.Any(), gomock.Any()).Return(expectedFinder).Times(1)
				return mockCollection
			},
			hooks: []finder.BeforeHookFn[TestUser]{
				func(ctx context.Context, opContext *finder.OpContext[TestUser], opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *finder.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.RegisterBeforeHooks(tc.hooks...)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_RegisterAfterHooks(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		hooks []finder.AfterHookFn[TestUser]
	}{
		{
			name: "register after hooks", mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedFinder := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().RegisterAfterHooks(gomock.Any(), gomock.Any()).Return(expectedFinder).Times(1)
				return mockCollection
			},
			hooks: []finder.AfterHookFn[TestUser]{
				func(ctx context.Context, opContext *finder.OpContext[TestUser], opts ...any) error {
					return nil
				},
				func(ctx context.Context, opContext *finder.OpContext[TestUser], opts ...any) error {
					return nil
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.RegisterAfterHooks(tc.hooks...)
			assert.NotNil(t, result)
		})
	}
}

func TestFinder_PreActionHandler(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context

		globalOpContext *operation.OpContext
		opContext       *finder.OpContext[TestUser]
		opTypes         []operation.OpType
		wantErr         error
	}{
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().PreActionHandler(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return mockCollection
			},
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &finder.OpContext[TestUser]{},
			opTypes:         []operation.OpType{operation.OpTypeBeforeFind},
			wantErr:         nil,
		},
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().PreActionHandler(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("pre action error")).Times(1)
				return mockCollection
			},
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &finder.OpContext[TestUser]{},
			opTypes:         []operation.OpType{operation.OpTypeBeforeFind},
			wantErr:         errors.New("pre action error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			err := finder.PreActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opTypes...)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestFinder_PostActionHandler(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser]
		ctx  context.Context

		globalOpContext *operation.OpContext
		opContext       *finder.OpContext[TestUser]
		opTypes         []operation.OpType
		wantErr         error
	}{
		{
			name: "success",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().PostActionHandler(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return mockCollection
			},
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &finder.OpContext[TestUser]{},
			opTypes:         []operation.OpType{operation.OpTypeAfterFind},
			wantErr:         nil,
		},
		{
			name: "error",
			mock: func(ctx context.Context, ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				mockCollection.EXPECT().PostActionHandler(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("post action error")).Times(1)
				return mockCollection
			},
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &finder.OpContext[TestUser]{},
			opTypes:         []operation.OpType{operation.OpTypeAfterFind},
			wantErr:         errors.New("post action error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(tc.ctx, ctl)

			err := finder.PostActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opTypes...)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestFinder_GetCollection(t *testing.T) {
	type TestUser struct {
		ID           bson.ObjectID `bson:"_id,omitempty"`
		Name         string        `bson:"name"`
		Age          int64
		UnknownField string    `bson:"-"`
		CreatedAt    time.Time `bson:"created_at"`
		UpdatedAt    time.Time `bson:"updated_at"`
	}
	testCases := []struct {
		name string
		mock func(ctl *gomock.Controller) finder.IFinder[TestUser]

		want *mongo.Collection
	}{
		{
			name: "get collection",
			mock: func(ctl *gomock.Controller) finder.IFinder[TestUser] {
				mockCollection := mocks.NewMockIFinder[TestUser](ctl)
				expectedCollection := &mongo.Collection{}
				mockCollection.EXPECT().GetCollection().Return(expectedCollection).Times(1)
				return mockCollection
			},
			want: &mongo.Collection{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()
			finder := tc.mock(ctl)

			result := finder.GetCollection()
			assert.Equal(t, tc.want, result)
		})
	}
}
