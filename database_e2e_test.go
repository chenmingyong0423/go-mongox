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

package mongox

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Test_e2e_newDatabase(t *testing.T) {
	c := getMongoClient(t)

	db := newDatabase(NewClient(c, &Config{}), "db-test")
	require.Equal(t, db.Database().Name(), "db-test")

	db.RegisterPlugin("global before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
		return nil
	}, operation.OpTypeBeforeFind)

	db.RemovePlugin("global before find", operation.OpTypeBeforeFind)
}

func getMongoClient(t *testing.T) *mongo.Client {
	c, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	require.NoError(t, err)

	return c
}

func TestRegisterPlugin_Insert(t *testing.T) {
	c := getMongoClient(t)

	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before insert
	t.Run("before insert", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			user := opCtx.Doc.(*User)
			require.NotNil(t, user)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeInsert)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before insert", operation.OpTypeBeforeInsert)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after insert
	t.Run("after insert", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")

		db.RegisterPlugin("after insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterInsert)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("after insert", operation.OpTypeAfterInsert)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Find(t *testing.T) {
	c := getMongoClient(t)
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before find
	t.Run("before find", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeFind)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before find", operation.OpTypeBeforeFind)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after find
	t.Run("after find", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("after find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			users := opCtx.Doc.([]*User)
			require.NotNil(t, users)
			isCalled = true
			return nil
		}, operation.OpTypeAfterFind)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("after find", operation.OpTypeAfterFind)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Delete(t *testing.T) {
	c := getMongoClient(t)
	type User struct {
		Name string
		Age  int
	}
	isCalled := false
	// before delete
	t.Run("before delete", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeDelete)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before delete", operation.OpTypeBeforeDelete)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after delete
	t.Run("after delete", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("after delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeAfterDelete)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("after delete", operation.OpTypeAfterDelete)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc([]*User{{Name: "Mingyong Chen", Age: 18}})), operation.OpTypeAfterDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Update(t *testing.T) {
	c := getMongoClient(t)
	isCalled := false
	// before update
	t.Run("before update", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpdate)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before update", operation.OpTypeBeforeUpdate)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after update
	t.Run("after update", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("after update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpdate)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("after update", operation.OpTypeAfterUpdate)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeAfterUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_Upsert(t *testing.T) {
	c := getMongoClient(t)
	isCalled := false
	// before upsert
	t.Run("before upsert", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeUpsert)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before upsert", operation.OpTypeBeforeUpsert)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// after upsert
	t.Run("after upsert", func(t *testing.T) {
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("after upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeAfterUpsert)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("after upsert", operation.OpTypeAfterUpsert)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeAfterUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}

func TestRegisterPlugin_BeforeAny(t *testing.T) {
	c := getMongoClient(t)

	// before find
	t.Run("before find", func(t *testing.T) {
		var isCalled bool
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeAny)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before find", operation.OpTypeBeforeAny)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeFind)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// before insert
	t.Run("before insert", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		var isCalled bool
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before insert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			user := opCtx.Doc.(*User)
			require.NotNil(t, user)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeAny)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before insert", operation.OpTypeBeforeAny)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})), operation.OpTypeBeforeInsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// before delete
	t.Run("before delete", func(t *testing.T) {
		var isCalled bool
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before delete", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeAny)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before delete", operation.OpTypeBeforeAny)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"})), operation.OpTypeBeforeDelete)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// before update
	t.Run("before update", func(t *testing.T) {
		var isCalled bool
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before update", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeAny)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before update", operation.OpTypeBeforeAny)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"$set": bson.M{"name": "Burt"}})), operation.OpTypeBeforeUpdate)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})

	// before upsert
	t.Run("before upsert", func(t *testing.T) {
		var isCalled bool
		db := newDatabase(NewClient(c, &Config{}), "db-test")
		db.RegisterPlugin("before upsert", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
			require.NotNil(t, opCtx.Filter)
			require.NotNil(t, opCtx.Updates)
			isCalled = true
			return nil
		}, operation.OpTypeBeforeAny)
		err := db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.True(t, isCalled)
		isCalled = false
		db.RemovePlugin("before upsert", operation.OpTypeBeforeAny)
		err = db.callbacks.Execute(context.Background(), operation.NewOpContext(nil, operation.WithFilter(bson.M{"name": "Mingyong Chen"}), operation.WithUpdates(bson.M{"name": "Burt"})), operation.OpTypeBeforeUpsert)
		require.Nil(t, err)
		assert.False(t, isCalled)
	})
}
