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

	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate optioner -type CondContext
type CondContext struct {
	Filter      any `opt:"-"`
	Updates     any
	Replacement any
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
