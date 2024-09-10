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

package finder

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate optioner -type OpContext
type OpContext struct {
	Col          *mongo.Collection `opt:"-"`
	Filter       any               `opt:"-"`
	Updates      any
	MongoOptions any
	ModelHook    any
}

//go:generate optioner -type AfterOpContext
type AfterOpContext[T any] struct {
	*OpContext `opt:"-"`
	Doc        *T
	Docs       []*T
}

type (
	beforeHookFn       func(ctx context.Context, opContext *OpContext, opts ...any) error
	afterHookFn[T any] func(ctx context.Context, opContext *AfterOpContext[T], opts ...any) error
)

type TestUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
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

type TestTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type IllegalUser struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  string
}

type UpdatedUser struct {
	Name string `bson:"name"`
	Age  int64
}

type UserName struct {
	Name string `bson:"name"`
}
