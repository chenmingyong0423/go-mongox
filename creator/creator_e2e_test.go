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

package creator_test

import (
	"context"
	"errors"
	"testing"
	"time"

	xcreator "github.com/chenmingyong0423/go-mongox/v2/creator"
	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type User struct {
	ID               bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
	Name             string        `bson:"name"`
	Age              int64
	UnknownField     string    `bson:"-"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
	DeletedAt        time.Time `bson:"deleted_at,omitempty"`
	CreateSecondTime int64     `bson:"create_second_time" mongox:"autoCreateTime:second"`
	UpdateSecondTime int64     `bson:"update_second_time" mongox:"autoUpdateTime:second"`
	CreateMilliTime  int64     `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
	UpdateMilliTime  int64     `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
	CreateNanoTime   int64     `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
	UpdateNanoTime   int64     `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
}

func newCollection(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
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
	creator := xcreator.NewCreator[User](collection, callback.InitializeCallbacks(), field.ParseFields(User{}))

	type globalHook struct {
		opType operation.OpType
		name   string
		fn     callback.CbFn
	}

	testCases := []struct {
		name       string
		before     func(ctx context.Context, t *testing.T)
		after      func(ctx context.Context, t *testing.T)
		opts       []options.Lister[options.InsertOneOptions]
		ctx        context.Context
		doc        *User
		globalHook []globalHook
		beforeHook []xcreator.HookFn[User]
		afterHook  []xcreator.HookFn[User]

		wantError assert.ErrorAssertionFunc
	}{
		{
			name:   "nil doc",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			ctx:    context.Background(),
			opts: []options.Lister[options.InsertOneOptions]{
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
			opts: []options.Lister[options.InsertOneOptions]{
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
			opts: []options.Lister[options.InsertOneOptions]{
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
			opts: []options.Lister[options.InsertOneOptions]{
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
			opts: []options.Lister[options.InsertOneOptions]{
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
			opts: []options.Lister[options.InsertOneOptions]{
				options.InsertOne().SetComment("test"),
			},
			doc: nil,
			beforeHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
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
			opts: []options.Lister[options.InsertOneOptions]{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
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
			opts: []options.Lister[options.InsertOneOptions]{
				options.InsertOne().SetComment("test"),
			},
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			beforeHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					if opContext.Doc == nil {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					if opContext == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			wantError: assert.NoError,
		},
		{
			name:   "validate field hook",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				user := &User{}
				err := collection.FindOne(ctx, query.Eq("name", "Mingyong Chen")).Decode(&user)
				if err != nil {
					require.NoError(t, err)
				}

				assert.NotZero(t, user.CreatedAt)
				assert.NotZero(t, user.UpdatedAt)
				assert.NotZero(t, user.CreateSecondTime)
				assert.NotZero(t, user.UpdateSecondTime)
				assert.NotZero(t, user.CreateMilliTime)
				assert.NotZero(t, user.UpdateMilliTime)
				assert.NotZero(t, user.CreateNanoTime)
				assert.NotZero(t, user.UpdateNanoTime)

				deleteResult, err := collection.DeleteOne(ctx, query.Eq("name", "Mingyong Chen"))
				require.NoError(t, err)
				require.Equal(t, int64(1), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			doc: &User{
				Name: "Mingyong Chen",
				Age:  18,
			},
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					user := opContext.Doc
					if user == nil {
						return errors.New("user is nil")
					}
					if user.Name != "Mingyong Chen" {
						return errors.New("name is not Mingyong Chen")
					}
					if user.Age != 18 {
						return errors.New("age is not 18")
					}
					if user.ID.IsZero() {
						return errors.New("id is zero")
					}
					if user.CreatedAt.IsZero() {
						return errors.New("created at is zero")
					}
					if user.UpdatedAt.IsZero() {
						return errors.New("updated at is zero")
					}
					if user.CreateSecondTime == 0 {
						return errors.New("create second time is zero")
					}
					if user.UpdateSecondTime == 0 {
						return errors.New("update second time is zero")
					}
					if user.CreateMilliTime == 0 {
						return errors.New("create milli time is zero")
					}
					if user.UpdateMilliTime == 0 {
						return errors.New("update milli time is zero")
					}
					if user.CreateNanoTime == 0 {
						return errors.New("create nano time is zero")
					}
					if user.UpdateNanoTime == 0 {
						return errors.New("update nano time is zero")
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
				creator.DBCallbacks.Register(hook.opType, hook.name, hook.fn)
			}
			insertOneResult, err := creator.RegisterBeforeHooks(tc.beforeHook...).
				RegisterAfterHooks(tc.afterHook...).InsertOne(tc.ctx, tc.doc, tc.opts...)
			tc.after(tc.ctx, t)
			if !tc.wantError(t, err) {
				return
			}
			if err == nil {
				require.NotNil(t, insertOneResult.InsertedID)
			}
			for _, hook := range tc.globalHook {
				creator.DBCallbacks.Remove(hook.opType, hook.name)
			}
			creator.BeforeHooks = nil
			creator.AfterHooks = nil
		})
	}
}

func TestCreator_e2e_Many(t *testing.T) {
	collection := newCollection(t)
	creator := xcreator.NewCreator[User](collection, callback.InitializeCallbacks(), field.ParseFields(User{}))

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
		opts       []options.Lister[options.InsertManyOptions]
		globalHook []globalHook
		beforeHook []xcreator.HookFn[User]
		afterHook  []xcreator.HookFn[User]

		wantIdsLength int
		wantError     assert.ErrorAssertionFunc
	}{
		{
			name:   "nil docs",
			before: func(ctx context.Context, t *testing.T) {},
			after:  func(ctx context.Context, t *testing.T) {},
			opts: []options.Lister[options.InsertManyOptions]{
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			opts: []options.Lister[options.InsertManyOptions]{
				options.InsertMany().SetComment("test"),
			},
			docs: nil,
			beforeHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
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
			opts: []options.Lister[options.InsertManyOptions]{
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
			beforeHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					if len(opContext.Docs) != 2 {
						return errors.New("before hook error")
					}
					return nil
				},
			},
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					if opContext == nil {
						return errors.New("after hook error")
					}
					return nil
				},
			},
			wantError:     assert.NoError,
			wantIdsLength: 2,
		},
		{
			name:   "validate field hook",
			before: func(ctx context.Context, t *testing.T) {},
			after: func(ctx context.Context, t *testing.T) {
				var users []*User
				cursor, err := collection.Find(ctx, query.In("name", "Mingyong Chen", "burt"))
				assert.NoError(t, err)
				defer cursor.Close(ctx)
				err = cursor.All(ctx, &users)
				assert.NoError(t, err)
				for _, user := range users {
					assert.NotZero(t, user.CreatedAt)
					assert.NotZero(t, user.UpdatedAt)
					assert.NotZero(t, user.CreateSecondTime)
					assert.NotZero(t, user.UpdateSecondTime)
					assert.NotZero(t, user.CreateMilliTime)
					assert.NotZero(t, user.UpdateMilliTime)
					assert.NotZero(t, user.CreateNanoTime)
					assert.NotZero(t, user.UpdateNanoTime)
				}

				deleteResult, err := collection.DeleteMany(ctx, query.In("name", "Mingyong Chen", "burt"))
				require.NoError(t, err)
				require.Equal(t, int64(2), deleteResult.DeletedCount)
			},
			ctx: context.Background(),
			docs: []*User{
				{
					Name: "Mingyong Chen",
					Age:  18,
				},
				{
					Name: "burt",
					Age:  18,
				},
			},
			afterHook: []xcreator.HookFn[User]{
				func(ctx context.Context, opContext *xcreator.OpContext[User], opts ...any) error {
					users := opContext.Docs
					if users == nil {
						return errors.New("users is nil")
					}
					for _, user := range users {
						if user.Name != "Mingyong Chen" && user.Name != "burt" {
							return errors.New("name is not Mingyong Chen or burt")
						}
						if user.Age != 18 {
							return errors.New("age is not 18")
						}
						if user.ID.IsZero() {
							return errors.New("id is zero")
						}
						if user.CreatedAt.IsZero() {
							return errors.New("created at is zero")
						}
						if user.UpdatedAt.IsZero() {
							return errors.New("updated at is zero")
						}
						if user.CreateSecondTime == 0 {
							return errors.New("create second time is zero")
						}
						if user.UpdateSecondTime == 0 {
							return errors.New("update second time is zero")
						}
						if user.CreateMilliTime == 0 {
							return errors.New("create milli time is zero")
						}
						if user.UpdateMilliTime == 0 {
							return errors.New("update milli time is zero")
						}
						if user.CreateNanoTime == 0 {
							return errors.New("create nano time is zero")
						}
						if user.UpdateNanoTime == 0 {
							return errors.New("update nano time is zero")
						}
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
				creator.DBCallbacks.Register(hook.opType, hook.name, hook.fn)
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
			}
			for _, hook := range tc.globalHook {
				creator.DBCallbacks.Remove(hook.opType, hook.name)
			}
			creator.BeforeHooks = nil
			creator.AfterHooks = nil
		})
	}
}
