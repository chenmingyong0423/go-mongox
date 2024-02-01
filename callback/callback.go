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

package callback

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/hook/field"
	"github.com/chenmingyong0423/go-mongox/operation"
)

type CbFn func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error

var Callbacks = initializeCallbacks()

func initializeCallbacks() *Callback {
	return &Callback{
		beforeInsert: []callbackHandler{
			{
				name: "mongox:default_field",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, operation.OpTypeBeforeInsert, opts...)
				},
			},
		},
		afterInsert: make([]callbackHandler, 0),
		beforeUpdate: []callbackHandler{
			{
				name: "mongox:default_field",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, operation.OpTypeBeforeUpdate, opts...)
				},
			},
		},
		afterUpdate:  make([]callbackHandler, 0),
		beforeDelete: make([]callbackHandler, 0),
		afterDelete:  make([]callbackHandler, 0),
		beforeUpsert: []callbackHandler{
			{
				name: "mongox:default_field",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return field.Execute(ctx, opCtx, operation.OpTypeBeforeUpsert, opts...)
				},
			},
		},
		afterUpsert: make([]callbackHandler, 0),
		beforeFind:  make([]callbackHandler, 0),
		afterFind:   make([]callbackHandler, 0),
	}
}

func GetCallback() *Callback {
	return Callbacks
}

type Callback struct {
	beforeInsert []callbackHandler
	afterInsert  []callbackHandler
	beforeUpdate []callbackHandler
	afterUpdate  []callbackHandler
	beforeDelete []callbackHandler
	afterDelete  []callbackHandler
	beforeUpsert []callbackHandler
	afterUpsert  []callbackHandler
	beforeFind   []callbackHandler
	afterFind    []callbackHandler
}

func (c *Callback) Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	switch opType {
	case operation.OpTypeBeforeInsert:
		return c.execute(ctx, opCtx, c.beforeInsert, opts...)
	case operation.OpTypeAfterInsert:
		return c.execute(ctx, opCtx, c.afterInsert, opts...)
	case operation.OpTypeBeforeUpdate:
		return c.execute(ctx, opCtx, c.beforeUpdate, opts...)
	case operation.OpTypeAfterUpdate:
		return c.execute(ctx, opCtx, c.afterUpdate, opts...)
	case operation.OpTypeBeforeDelete:
		return c.execute(ctx, opCtx, c.beforeDelete, opts...)
	case operation.OpTypeAfterDelete:
		return c.execute(ctx, opCtx, c.afterDelete, opts...)
	case operation.OpTypeBeforeUpsert:
		return c.execute(ctx, opCtx, c.beforeUpsert, opts...)
	case operation.OpTypeAfterUpsert:
		return c.execute(ctx, opCtx, c.afterUpsert, opts...)
	case operation.OpTypeBeforeFind:
		return c.execute(ctx, opCtx, c.beforeFind, opts...)
	case operation.OpTypeAfterFind:
		return c.execute(ctx, opCtx, c.afterFind, opts...)
	}
	return nil
}

func (c *Callback) execute(ctx context.Context, opCtx *operation.OpContext, handlers []callbackHandler, opts ...any) error {
	for _, handler := range handlers {
		if err := handler.fn(ctx, opCtx, opts...); err != nil {
			return err
		}
	}
	return nil
}

func (c *Callback) Register(opType operation.OpType, name string, fn CbFn) {
	switch opType {
	case operation.OpTypeBeforeInsert:
		c.beforeInsert = append(c.beforeInsert, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeAfterInsert:
		c.afterInsert = append(c.afterInsert, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeBeforeUpdate:
		c.beforeUpdate = append(c.beforeUpdate, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeAfterUpdate:
		c.afterUpdate = append(c.afterUpdate, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeBeforeDelete:
		c.beforeDelete = append(c.beforeDelete, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeAfterDelete:
		c.afterDelete = append(c.afterDelete, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeBeforeUpsert:
		c.beforeUpsert = append(c.beforeUpsert, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeAfterUpsert:
		c.afterUpsert = append(c.afterUpsert, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeBeforeFind:
		c.beforeFind = append(c.beforeFind, callbackHandler{
			name: name,
			fn:   fn,
		})
	case operation.OpTypeAfterFind:
		c.afterFind = append(c.afterFind, callbackHandler{
			name: name,
			fn:   fn,
		})
	}
}

type callbackHandler struct {
	name string
	fn   CbFn
}
