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

package updater

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate optioner -type CondContext
type CondContext struct {
	Filter       any `opt:"-"`
	Updates      any
	Replacement  any
	MongoOptions any
	ModelHook    any
}

//go:generate optioner -type BeforeOpContext
type BeforeOpContext struct {
	Col          *mongo.Collection `opt:"-"`
	*CondContext `opt:"-"`
}

//go:generate optioner -type AfterOpContext
type AfterOpContext struct {
	Col          *mongo.Collection `opt:"-"`
	*CondContext `opt:"-"`
}

type (
	beforeHookFn func(ctx context.Context, opContext *BeforeOpContext, opts ...any) error
	afterHookFn  func(ctx context.Context, opContext *AfterOpContext, opts ...any) error
)

type User struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type UserTestField struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64  `bson:"age,omitempty"`
	UnknownField string `bson:"-"`
}

type TestUser struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (m *TestUser) DefaultId() bson.ObjectID {
	if m.ID.IsZero() {
		m.ID = bson.NewObjectID()
	}
	return m.ID
}

func (m *TestUser) DefaultCreatedAt() time.Time {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
	return m.CreatedAt
}

func (m *TestUser) DefaultUpdatedAt() time.Time {
	m.UpdatedAt = time.Now().Local()
	return m.UpdatedAt
}

type TestUser2 struct {
	ID           string `bson:"_id,omitempty"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

func (m *TestUser2) CustomID() (string, any) {
	if m.ID == "" {
		m.ID = bson.NewObjectID().Hex()
	}
	return "_id", m.ID
}

func (m *TestUser2) CustomCreatedAt() (string, any) {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
	return "createdAt", m.CreatedAt
}

func (m *TestUser2) CustomUpdatedAt() (string, any) {
	m.UpdatedAt = time.Now().Local()
	return "updatedAt", m.UpdatedAt
}
