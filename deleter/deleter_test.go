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
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/deleter"
	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/drivertest"
)

type testDeleteUser struct {
	Name string `bson:"name"`
}

func newMockDeleteCollection(t *testing.T) (*mongo.Collection, *drivertest.MockDeployment, *[]bson.Raw) {
	t.Helper()
	deployment := drivertest.NewMockDeployment()
	var commands []bson.Raw
	opts := options.Client().SetRetryWrites(false).SetMonitor(&event.CommandMonitor{
		Started: func(_ context.Context, ev *event.CommandStartedEvent) {
			commands = append(commands, append(bson.Raw(nil), ev.Command...))
		},
	})
	opts.Deployment = deployment
	client, err := mongo.Connect(opts)
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, client.Disconnect(context.Background())) })
	return client.Database("test").Collection("users"), deployment, &commands
}

type TestUser struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	DeletedAt    time.Time `bson:"deleted_at"`
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
		name      string
		ctx       context.Context
		filter    any
		opts      []options.Lister[options.DeleteOneOptions]
		response  bson.D
		beforeErr error
		afterErr  error
		want      *mongo.DeleteResult
		wantAfter bool
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			ctx:  context.Background(),
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.ErrorIs(t, err, mongo.ErrNilDocument)
			},
		},
		{
			name:      "deleted count: 0",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			response:  bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(0)}},
			want:      &mongo.DeleteResult{Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
		{
			name:      "delete success",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			want:      &mongo.DeleteResult{DeletedCount: 1, Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
		{
			name:      "delete success with options",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			opts:      []options.Lister[options.DeleteOneOptions]{options.DeleteOne().SetComment("physical-delete")},
			want:      &mongo.DeleteResult{DeletedCount: 1, Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
		{
			name:      "before hook error",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			beforeErr: assert.AnError,
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.ErrorIs(t, err, assert.AnError)
			},
		},
		{
			name:      "after hook error",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			afterErr:  assert.AnError,
			wantAfter: true,
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.ErrorIs(t, err, assert.AnError)
			},
		},
		{
			name:     "driver error",
			ctx:      context.Background(),
			filter:   bson.M{"name": "test"},
			response: bson.D{{Key: "ok", Value: 0}, {Key: "code", Value: 121}, {Key: "errmsg", Value: "validation failed"}},
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				var commandErr mongo.CommandError
				return assert.ErrorAs(t, err, &commandErr) && assert.Equal(t, int32(121), commandErr.Code)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			collection, deployment, commands := newMockDeleteCollection(t)
			if tc.response == nil {
				deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(1)}})
			} else {
				deployment.AddResponses(tc.response)
			}
			d := deleter.NewDeleter[testDeleteUser](collection, callback.InitializeCallbacks(), field.ParseFields(testDeleteUser{}))
			afterCalled := false
			d.Filter(tc.filter).RegisterBeforeHooks(func(context.Context, *deleter.OpContext, ...any) error {
				return tc.beforeErr
			}).RegisterAfterHooks(func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
				afterCalled = true
				if tc.want != nil {
					assert.Equal(t, tc.want, opCtx.Result)
				}
				return tc.afterErr
			})

			got, err := d.DeleteOne(tc.ctx, tc.opts...)
			tc.wantErr(t, err)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantAfter, afterCalled)
			if tc.filter == nil || tc.beforeErr != nil {
				assert.Empty(t, *commands)
			}
			if tc.want != nil {
				require.Len(t, *commands, 1)
				command := (*commands)[0]
				assert.Equal(t, "users", command.Lookup("delete").StringValue())
				deletion := command.Lookup("deletes").Array().Index(0).Document()
				assert.Equal(t, int32(1), deletion.Lookup("limit").Int32())
				assert.Equal(t, "test", deletion.Lookup("q", "name").StringValue())
				if len(tc.opts) > 0 {
					assert.Equal(t, "physical-delete", command.Lookup("comment").StringValue())
				}
			}
		})
	}

	t.Run("soft delete", func(t *testing.T) {
		collection, deployment, commands := newMockDeleteCollection(t)
		deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(1)}, {Key: "nModified", Value: int32(1)}})
		d := deleter.NewDeleter[TestUser](collection, callback.InitializeCallbacks(), field.ParseFields(TestUser{}))
		d.Filter(bson.M{"name": "test"}).UpdateOneOptions(options.UpdateOne().SetComment("soft-delete"))

		got, err := d.DeleteOne(context.Background(), options.DeleteOne().SetComment("physical-delete"))
		require.NoError(t, err)
		assert.Equal(t, &mongo.DeleteResult{DeletedCount: 1, Acknowledged: true}, got)
		require.Len(t, *commands, 1)
		command := (*commands)[0]
		assert.Equal(t, "users", command.Lookup("update").StringValue())
		assert.Equal(t, "soft-delete", command.Lookup("comment").StringValue())
		update := command.Lookup("updates").Array().Index(0).Document()
		multi, _ := update.Lookup("multi").BooleanOK()
		assert.Equal(t, false, multi)
		assert.Equal(t, bson.TypeDateTime, update.Lookup("u", "$set", "deleted_at").Type)
		assert.Zero(t, update.Lookup("u", "$set", "updated_at").Type)
	})
}

func TestDeleter_DeleteMany(t *testing.T) {
	testCases := []struct {
		name      string
		ctx       context.Context
		filter    any
		opts      []options.Lister[options.DeleteManyOptions]
		response  bson.D
		want      *mongo.DeleteResult
		wantAfter bool
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "error: nil filter",
			ctx:  context.Background(),
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.ErrorIs(t, err, mongo.ErrNilDocument)
			},
		},
		{
			name:      "deleted count: 0",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			response:  bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(0)}},
			want:      &mongo.DeleteResult{Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
		{
			name:      "delete success",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			want:      &mongo.DeleteResult{DeletedCount: 2, Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
		{
			name:      "delete success with options",
			ctx:       context.Background(),
			filter:    bson.M{"name": "test"},
			opts:      []options.Lister[options.DeleteManyOptions]{options.DeleteMany().SetComment("physical-delete")},
			want:      &mongo.DeleteResult{DeletedCount: 2, Acknowledged: true},
			wantAfter: true,
			wantErr:   assert.NoError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			collection, deployment, commands := newMockDeleteCollection(t)
			if tc.response == nil {
				deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(2)}})
			} else {
				deployment.AddResponses(tc.response)
			}
			d := deleter.NewDeleter[testDeleteUser](collection, callback.InitializeCallbacks(), field.ParseFields(testDeleteUser{}))
			afterCalled := false
			d.Filter(tc.filter).RegisterAfterHooks(func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
				afterCalled = true
				if tc.want != nil {
					assert.Equal(t, tc.want, opCtx.Result)
				}
				return nil
			})

			got, err := d.DeleteMany(tc.ctx, tc.opts...)
			tc.wantErr(t, err)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantAfter, afterCalled)
			if tc.want != nil {
				require.Len(t, *commands, 1)
				command := (*commands)[0]
				assert.Equal(t, "users", command.Lookup("delete").StringValue())
				deletion := command.Lookup("deletes").Array().Index(0).Document()
				assert.Equal(t, int32(0), deletion.Lookup("limit").Int32())
				assert.Equal(t, "test", deletion.Lookup("q", "name").StringValue())
				if len(tc.opts) > 0 {
					assert.Equal(t, "physical-delete", command.Lookup("comment").StringValue())
				}
			}
		})
	}

	t.Run("soft delete", func(t *testing.T) {
		collection, deployment, commands := newMockDeleteCollection(t)
		deployment.AddResponses(bson.D{{Key: "ok", Value: 1}, {Key: "n", Value: int32(4)}, {Key: "nModified", Value: int32(2)}})
		d := deleter.NewDeleter[TestUser](collection, callback.InitializeCallbacks(), field.ParseFields(TestUser{}))
		d.Filter(bson.M{"name": "test"}).UpdateManyOptions(options.UpdateMany().SetComment("soft-delete"))

		got, err := d.DeleteMany(context.Background(), options.DeleteMany().SetComment("physical-delete"))
		require.NoError(t, err)
		assert.Equal(t, &mongo.DeleteResult{DeletedCount: 2, Acknowledged: true}, got)
		require.Len(t, *commands, 1)
		command := (*commands)[0]
		assert.Equal(t, "users", command.Lookup("update").StringValue())
		assert.Equal(t, "soft-delete", command.Lookup("comment").StringValue())
		update := command.Lookup("updates").Array().Index(0).Document()
		multi, _ := update.Lookup("multi").BooleanOK()
		assert.Equal(t, true, multi)
		assert.Equal(t, bson.TypeDateTime, update.Lookup("u", "$set", "deleted_at").Type)
		assert.Zero(t, update.Lookup("u", "$set", "updated_at").Type)
	})
}

func TestDeleter_UpdateOneOptions(t *testing.T) {
	d := deleter.NewDeleter[TestUser](&mongo.Collection{}, nil, field.ParseFields(TestUser{}))
	calls := 0
	testCases := []struct {
		name string
		opts []options.Lister[options.UpdateOneOptions]
		want deleter.IDeleter[TestUser]
	}{
		{
			name: "empty options",
			want: d,
		},
		{
			name: "multiple options",
			opts: []options.Lister[options.UpdateOneOptions]{options.UpdateOne().SetComment("test"), options.UpdateOne().SetHint("_id_")},
			want: d,
		},
		{
			name: "deferred option setter",
			opts: []options.Lister[options.UpdateOneOptions]{&options.UpdateOneOptionsBuilder{Opts: []func(*options.UpdateOneOptions) error{
				func(*options.UpdateOneOptions) error {
					calls++
					return assert.AnError
				},
			}}},
			want: d,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := d.UpdateOneOptions(tc.opts...)
			assert.Same(t, tc.want, result)
			assert.Zero(t, calls)
		})
	}
}

func TestDeleter_UpdateManyOptions(t *testing.T) {
	d := deleter.NewDeleter[TestUser](&mongo.Collection{}, nil, field.ParseFields(TestUser{}))
	calls := 0
	testCases := []struct {
		name string
		opts []options.Lister[options.UpdateManyOptions]
		want deleter.IDeleter[TestUser]
	}{
		{
			name: "empty options",
			want: d,
		},
		{
			name: "multiple options",
			opts: []options.Lister[options.UpdateManyOptions]{options.UpdateMany().SetComment("test"), options.UpdateMany().SetHint("_id_")},
			want: d,
		},
		{
			name: "deferred option setter",
			opts: []options.Lister[options.UpdateManyOptions]{&options.UpdateManyOptionsBuilder{Opts: []func(*options.UpdateManyOptions) error{
				func(*options.UpdateManyOptions) error {
					calls++
					return assert.AnError
				},
			}}},
			want: d,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := d.UpdateManyOptions(tc.opts...)
			assert.Same(t, tc.want, result)
			assert.Zero(t, calls)
		})
	}
}

func TestDeleter_Filter(t *testing.T) {
	testCases := []struct {
		name   string
		filter any
	}{
		{
			name:   "set filter successfully",
			filter: bson.M{"name": "test"},
		},
		{
			name: "clear filter",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			callbacks := callback.InitializeCallbacks()
			callbacks.Register(operation.OpTypeBeforeDelete, "inspect context", func(_ context.Context, opCtx *operation.OpContext, _ ...any) error {
				assert.Equal(t, tc.filter, opCtx.Filter)
				return assert.AnError
			})
			d := deleter.NewDeleter[testDeleteUser](nil, callbacks, field.ParseFields(testDeleteUser{}))
			d.Filter(bson.M{"name": "old"})

			result := d.Filter(tc.filter)
			assert.Same(t, d, result)
			got, err := d.DeleteOne(context.Background())
			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, got)
		})
	}
}

func TestDeleter_ModelHook(t *testing.T) {
	testCases := []struct {
		name      string
		modelHook any
	}{
		{
			name:      "set model hook successfully",
			modelHook: &TestUser{Name: "hook"},
		},
		{
			name: "clear model hook",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			callbacks := callback.InitializeCallbacks()
			callbacks.Register(operation.OpTypeBeforeDelete, "inspect context", func(_ context.Context, opCtx *operation.OpContext, _ ...any) error {
				assert.Equal(t, tc.modelHook, opCtx.ModelHook)
				return assert.AnError
			})
			d := deleter.NewDeleter[testDeleteUser](nil, callbacks, field.ParseFields(testDeleteUser{}))
			d.ModelHook(&TestUser{Name: "old"})

			result := d.ModelHook(tc.modelHook)
			assert.Same(t, d, result)
			got, err := d.DeleteOne(context.Background())
			assert.ErrorIs(t, err, assert.AnError)
			assert.Nil(t, got)
		})
	}
}

func TestDeleter_RegisterBeforeHooks(t *testing.T) {
	var calls []int
	testCases := []struct {
		name  string
		hooks []deleter.BeforeHookFn
		want  []int
	}{
		{
			name: "register empty hooks",
			want: []int{1},
		},
		{
			name: "register single before hook successfully",
			hooks: []deleter.BeforeHookFn{
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 2)
					return nil
				},
			},
			want: []int{1, 2},
		},
		{
			name: "register multiple before hooks successfully",
			hooks: []deleter.BeforeHookFn{
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 2)
					return nil
				},
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 3)
					return nil
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls = nil
			d := deleter.NewDeleter[testDeleteUser](nil, callback.InitializeCallbacks(), field.ParseFields(testDeleteUser{}))
			d.RegisterBeforeHooks(func(context.Context, *deleter.OpContext, ...any) error {
				calls = append(calls, 1)
				return nil
			})

			result := d.RegisterBeforeHooks(tc.hooks...)
			assert.Same(t, d, result)
			err := d.PreActionHandler(context.Background(), &operation.OpContext{}, &deleter.OpContext{}, operation.OpTypeBeforeDelete)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, calls)
		})
	}
}

func TestDeleter_RegisterAfterHooks(t *testing.T) {
	var calls []int
	testCases := []struct {
		name  string
		hooks []deleter.AfterHookFn
		want  []int
	}{
		{
			name: "register empty hooks",
			want: []int{1},
		},
		{
			name: "register single after hook successfully",
			hooks: []deleter.AfterHookFn{
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 2)
					return nil
				},
			},
			want: []int{1, 2},
		},
		{
			name: "register multiple after hooks successfully",
			hooks: []deleter.AfterHookFn{
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 2)
					return nil
				},
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, 3)
					return nil
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls = nil
			d := deleter.NewDeleter[testDeleteUser](nil, callback.InitializeCallbacks(), field.ParseFields(testDeleteUser{}))
			d.RegisterAfterHooks(func(context.Context, *deleter.OpContext, ...any) error {
				calls = append(calls, 1)
				return nil
			})

			result := d.RegisterAfterHooks(tc.hooks...)
			assert.Same(t, d, result)
			err := d.PostActionHandler(context.Background(), &operation.OpContext{}, &deleter.OpContext{}, operation.OpTypeAfterDelete)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, calls)
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
			d := deleter.NewDeleter[TestUser](tc.want, nil, field.ParseFields(TestUser{}))

			result := d.GetCollection()
			assert.Same(t, tc.want, result)
		})
	}
}

func TestDeleter_PreActionHandler(t *testing.T) {
	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	filter := bson.M{"name": "alice"}
	fields := field.ParseFields(TestUser{})
	scopedFilter := bson.D{{Key: "$and", Value: bson.A{filter, bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$in", Value: bson.A{nil, time.Time{}}}}}}}}}
	testCases := []struct {
		name            string
		ctx             context.Context
		globalOpContext *operation.OpContext
		opContext       *deleter.OpContext
		opType          operation.OpType
		globalHook      callback.CbFn
		beforeHook      []deleter.BeforeHookFn
		after           func(t *testing.T, globalOpContext *operation.OpContext, opContext *deleter.OpContext)
		wantErr         bool
		expectedError   error
	}{
		{
			name:            "pre action handler success",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
		},
		{
			name:            "global before hook error",
			ctx:             context.Background(),
			globalOpContext: &operation.OpContext{},
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			globalHook: func(context.Context, *operation.OpContext, ...any) error {
				return assert.AnError
			},
			wantErr:       true,
			expectedError: assert.AnError,
		},
		{
			name:            "soft delete filter and updates",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			beforeHook: []deleter.BeforeHookFn{
				func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
					assert.Equal(t, scopedFilter, opCtx.Filter)
					assert.Equal(t, bson.M{"$set": bson.M{"deleted_at": now}}, opCtx.Updates)
					opCtx.Filter = bson.M{"name": "bob"}
					opCtx.Updates = bson.M{"$set": bson.M{"deleted_at": now, "deleted_by": "hook"}}
					return nil
				},
			},
			after: func(t *testing.T, globalOpContext *operation.OpContext, opContext *deleter.OpContext) {
				assert.True(t, globalOpContext.SoftDelete)
				assert.Equal(t, bson.M{"name": "bob"}, globalOpContext.Filter)
				assert.Equal(t, bson.M{"$set": bson.M{"deleted_at": now, "deleted_by": "hook"}}, globalOpContext.Updates)
				assert.Equal(t, globalOpContext.Filter, opContext.Filter)
				assert.Equal(t, globalOpContext.Updates, opContext.Updates)
				assert.Equal(t, bson.M{"name": "alice"}, filter)
			},
		},
		{
			name:            "soft delete with cleared updates",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			beforeHook: []deleter.BeforeHookFn{
				func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
					opCtx.Updates = nil
					return nil
				},
			},
			after: func(t *testing.T, globalOpContext *operation.OpContext, opContext *deleter.OpContext) {
				assert.True(t, globalOpContext.SoftDelete)
				assert.Nil(t, globalOpContext.Updates)
				assert.Nil(t, opContext.Updates)
			},
		},
		{
			name:            "physical delete with updates",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter)),
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			beforeHook: []deleter.BeforeHookFn{
				func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
					assert.Equal(t, filter, opCtx.Filter)
					assert.Nil(t, opCtx.Updates)
					opCtx.Updates = bson.M{"$set": bson.M{"deleted_by": "hook"}}
					return nil
				},
			},
			after: func(t *testing.T, globalOpContext *operation.OpContext, opContext *deleter.OpContext) {
				assert.False(t, globalOpContext.SoftDelete)
				assert.Equal(t, bson.M{"$set": bson.M{"deleted_by": "hook"}}, globalOpContext.Updates)
				assert.Equal(t, globalOpContext.Updates, opContext.Updates)
			},
		},
		{
			name:            "local before hook error",
			ctx:             context.Background(),
			globalOpContext: operation.NewOpContext(nil, operation.WithFilter(filter), operation.WithFields(fields), operation.WithStartTime(now)),
			opContext:       &deleter.OpContext{},
			opType:          operation.OpTypeBeforeDelete,
			beforeHook: []deleter.BeforeHookFn{
				func(context.Context, *deleter.OpContext, ...any) error { return assert.AnError },
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
			d := deleter.NewDeleter[TestUser](nil, callbacks, fields)
			d.RegisterBeforeHooks(tc.beforeHook...)
			err := d.PreActionHandler(tc.ctx, tc.globalOpContext, tc.opContext, tc.opType)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
			if tc.after != nil {
				tc.after(t, tc.globalOpContext, tc.opContext)
			}
		})
	}
}

func TestDeleter_PostActionHandler(t *testing.T) {
	testCases := []struct {
		name          string
		ctx           context.Context
		globalErr     error
		localErr      error
		want          []string
		wantErr       bool
		expectedError error
	}{
		{
			name: "post action handler success",
			ctx:  context.Background(),
			want: []string{"global", "local", "second local"},
		},
		{
			name:          "global after hook error",
			ctx:           context.Background(),
			globalErr:     assert.AnError,
			want:          []string{"global"},
			wantErr:       true,
			expectedError: assert.AnError,
		},
		{
			name:          "local after hook error",
			ctx:           context.Background(),
			localErr:      assert.AnError,
			want:          []string{"global", "local"},
			wantErr:       true,
			expectedError: assert.AnError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := &mongo.DeleteResult{DeletedCount: 1, Acknowledged: true}
			var calls []string
			callbacks := callback.InitializeCallbacks()
			callbacks.Register(operation.OpTypeAfterDelete, tc.name, func(_ context.Context, opCtx *operation.OpContext, _ ...any) error {
				calls = append(calls, "global")
				assert.Same(t, result, opCtx.Result)
				return tc.globalErr
			})
			d := deleter.NewDeleter[TestUser](nil, callbacks, field.ParseFields(TestUser{}))
			d.RegisterAfterHooks(
				func(_ context.Context, opCtx *deleter.OpContext, _ ...any) error {
					calls = append(calls, "local")
					assert.Same(t, result, opCtx.Result)
					return tc.localErr
				},
				func(context.Context, *deleter.OpContext, ...any) error {
					calls = append(calls, "second local")
					return nil
				},
			)
			globalOpContext := operation.NewOpContext(nil, operation.WithResult(result))
			opContext := deleter.NewOpContext(nil, nil, deleter.WithResult(result))

			err := d.PostActionHandler(tc.ctx, globalOpContext, opContext, operation.OpTypeAfterDelete)
			if tc.wantErr {
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, calls)
		})
	}
}
