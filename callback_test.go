// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongox

import (
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRegisterPlugin_Insert(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before insert
	t.Run("before insert", func(t *testing.T) {
		RegisterPlugin("before insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			user := opCtx.Doc.(*User)
			require.NotNil(t, user)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeInsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before insert", operation.OpTypeBeforeInsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after insert
	t.Run("after insert", func(t *testing.T) {
		RegisterPlugin("after insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterInsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after insert", operation.OpTypeAfterInsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Find(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before find
	t.Run("before find", func(t *testing.T) {
		RegisterPlugin("before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeFind)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before find", operation.OpTypeBeforeFind)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after find
	t.Run("after find", func(t *testing.T) {
		RegisterPlugin("after find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterFind)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after find", operation.OpTypeAfterFind)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Delete(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before delete
	t.Run("before delete", func(t *testing.T) {
		RegisterPlugin("before delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeDelete)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before delete", operation.OpTypeBeforeDelete)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after delete
	t.Run("after delete", func(t *testing.T) {
		RegisterPlugin("after delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeAfterDelete)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after delete", operation.OpTypeAfterDelete)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Update(t *testing.T) {
	isCalled := false
	// before update
	t.Run("before update", func(t *testing.T) {
		RegisterPlugin("before update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpdate)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdate(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before update", operation.OpTypeBeforeUpdate)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdate(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after update
	t.Run("after update", func(t *testing.T) {
		RegisterPlugin("after update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpdate)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdate(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after update", operation.OpTypeAfterUpdate)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdate(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Upsert(t *testing.T) {
	isCalled := false
	// before upsert
	t.Run("before upsert", func(t *testing.T) {
		RegisterPlugin("before upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Replacement)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("before upsert", operation.OpTypeBeforeUpsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after upsert
	t.Run("after upsert", func(t *testing.T) {
		RegisterPlugin("after upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Replacement)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpsert)
		err := callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		RemovePlugin("after upsert", operation.OpTypeAfterUpsert)
		err = callback.Callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithReplacement(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}
