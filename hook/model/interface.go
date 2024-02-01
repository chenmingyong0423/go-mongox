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

package model

import "context"

type BeforeInsert interface {
	BeforeInsert(ctx context.Context) error
}

type AfterInsert interface {
	AfterInsert(ctx context.Context) error
}

type BeforeUpdate interface {
	BeforeUpdate(ctx context.Context) error
}

type AfterUpdate interface {
	AfterUpdate(ctx context.Context) error
}

type BeforeUpsert interface {
	BeforeUpsert(ctx context.Context) error
}

type AfterUpsert interface {
	AfterUpsert(ctx context.Context) error
}

type BeforeDelete interface {
	BeforeDelete(ctx context.Context) error
}

type AfterDelete interface {
	AfterDelete(ctx context.Context) error
}

type BeforeFind interface {
	BeforeFind(ctx context.Context) error
}

type AfterFind interface {
	AfterFind(ctx context.Context) error
}
