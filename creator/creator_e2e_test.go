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

//go:build e2e

package creator

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/builder/query"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (u *User) DefaultId() {
	if u.ID.IsZero() {
		u.ID = primitive.NewObjectID()
	}
}

func (u *User) DefaultCreatedAt() {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().Local()
	}
}

func (u *User) DefaultUpdatedAt() {
	u.UpdatedAt = time.Now().Local()
}

func newCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := client.Database("db-test").Collection("test_user")
	return collection
}

func TestCreator_e2e_One(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[User](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name       string
		before     func(ctx context.Context, t *testing.T)
		after      func(ctx context.Context, t *testing.T)
		opts       []*options.InsertOneOptions
		ctx        context.Context
		doc        *User
		globalHook []globalHook
		beforeHook []hookFn[User]
		afterHook  []hookFn[User]

		wantError assert.ErrorAssertionFunc
	}{
		{
			name:   "nil doc",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc:       nil,
			wantError: assert.Error,
		},
		{
			name:   "insert one successfully",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeInsert,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name:   "global after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterInsert,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name:   "global before and after hook ",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeInsert,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if user, ok := opCtx.Doc.(*User); !ok || user == nil {
							return errors.New("before hook error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterInsert,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx == nil {
							return errors.New("after hook error")
						}
						return nil
					},
				},
			},
			wantError: assert.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: nil,
			beforeHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name:   "after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			afterHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name:   "before and after hook ",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertOneOptions{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			beforeHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					if opContext.Doc == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					if opContext == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			wantError: assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			insertOneResult, err := creator.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).InsertOne(tc.ctx, tc.doc, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if err == nil {
				require.NotNil(t, insertOneResult.InsertedID)
				require.NotZero(t, tc.doc.CreatedAt)
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			creator.beforeHooks = nil
			creator.afterHooks = nil
		})
	}
}

func TestCreator_e2e_Many(t *testing.T) {
	collection := newCollection(t)
	creator := NewCreator[User](collection)

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name   string
		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx        context.Context
		docs       []*User
		opts       []*options.InsertManyOptions
		globalHook []globalHook
		beforeHook []hookFn[User]
		afterHook  []hookFn[User]

		wantIdsLength int
		wantError     assert.ErrorAssertionFunc
	}{
		{
			name:   "nil docs",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			ctx:       context.Background(),
			docs:      nil,
			wantError: assert.Error,
		},
		{
			name:   "insert many successfully",
			before: func(_ context.Context, _ *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			ctx: context.Background(),
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
			wantIdsLength: 2,
			wantError:     assert.NoError,
		},
		{
			name:   "global before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: nil,
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeInsert,
					name:   "before hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("before hook error")
					},
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name:   "global after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeAfterInsert,
					name:   "after hook error",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						return errors.New("after hook error")
					},
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name:   "global before and after hook ",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
			globalHook: []globalHook{
				{
					opType: operation.OpTypeBeforeInsert,
					name:   "before hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if users, ok := opCtx.Doc.([]*User); !ok || len(users) != 2 {
							return errors.New("before hook error")
						}
						return nil
					},
				},
				{
					opType: operation.OpTypeAfterInsert,
					name:   "after hook",
					fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
						if opCtx == nil {
							return errors.New("after hook error")
						}
						return nil
					},
				},
			},
			wantIdsLength: 2,
			wantError:     assert.NoError,
		},
		{
			name:   "before hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: nil,
			beforeHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					return errors.New("before hook error")
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("before hook error"), err)
			},
		},
		{
			name:   "after hook error",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
			afterHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					return errors.New("after hook error")
				},
			},
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Equal(t, errors.New("after hook error"), err)
			},
		},
		{
			name:   "before and after hook ",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				assert.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			opts: []*options.InsertManyOptions{
				options.InsertMany().SetComment("test"),
			},
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  19,
				},
			},
			beforeHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					if len(opContext.Docs) != 2 {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []hookFn[User]{
				func(ctx context.Context, opContext *OpContext[User], opts ...any) error {
					if opContext == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			wantError:     assert.NoError,
			wantIdsLength: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			for _, hook := range tc.globalHook {
				callback.GetCallback().Register(hook.opType, hook.name, hook.fn)
			}
			insertManyResult, err := creator.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).InsertMany(tc.ctx, tc.docs, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if err == nil {
				require.NotNil(t, insertManyResult)
				require.Len(t, insertManyResult.InsertedIDs, tc.wantIdsLength)
				for _, doc := range tc.docs {
					require.NotZero(t, doc.CreatedAt)
				}
			}
			for _, hook := range tc.globalHook {
				callback.GetCallback().Remove(hook.opType, hook.name)
			}
			creator.beforeHooks = nil
			creator.afterHooks = nil
		})
	}
}
