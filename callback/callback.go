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

	"github.com/chenmingyong0423/go-mongox/hook/validator"

	"github.com/chenmingyong0423/go-mongox/hook/model"

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
			{
				name: "mongox:model",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return model.Execute(ctx, opCtx, operation.OpTypeBeforeInsert, opts...)
				},
			},
			{
				name: "mongox:validation",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return validator.Execute(ctx, opCtx, operation.OpTypeBeforeInsert, opts...)
				},
			},
		},
		afterInsert: []callbackHandler{
			{
				name: "mongox:model",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return model.Execute(ctx, opCtx, operation.OpTypeAfterInsert, opts...)
				},
			},
		},
		beforeUpdate: make([]callbackHandler, 0),
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
			{
				name: "mongox:model",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return model.Execute(ctx, opCtx, operation.OpTypeBeforeUpsert, opts...)
				},
			},
			{
				name: "mongox:validation",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return validator.Execute(ctx, opCtx, operation.OpTypeBeforeUpsert, opts...)
				},
			},
		},
		afterUpsert: []callbackHandler{
			{
				name: "mongox:model",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return model.Execute(ctx, opCtx, operation.OpTypeAfterUpsert, opts...)
				},
			},
		},
		beforeFind: make([]callbackHandler, 0),
		afterFind: []callbackHandler{
			{
				name: "mongox:model",
				fn: func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
					return model.Execute(ctx, opCtx, operation.OpTypeAfterFind, opts...)
				},
			},
		},
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

func (c *Callback) Remove(opType operation.OpType, name string) {
	switch opType {
	case operation.OpTypeBeforeInsert:
		c.beforeInsert = c.remove(c.beforeInsert, name)
	case operation.OpTypeAfterInsert:
		c.afterInsert = c.remove(c.afterInsert, name)
	case operation.OpTypeBeforeUpdate:
		c.beforeUpdate = c.remove(c.beforeUpdate, name)
	case operation.OpTypeAfterUpdate:
		c.afterUpdate = c.remove(c.afterUpdate, name)
	case operation.OpTypeBeforeDelete:
		c.beforeDelete = c.remove(c.beforeDelete, name)
	case operation.OpTypeAfterDelete:
		c.afterDelete = c.remove(c.afterDelete, name)
	case operation.OpTypeBeforeUpsert:
		c.beforeUpsert = c.remove(c.beforeUpsert, name)
	case operation.OpTypeAfterUpsert:
		c.afterUpsert = c.remove(c.afterUpsert, name)
	case operation.OpTypeBeforeFind:
		c.beforeFind = c.remove(c.beforeFind, name)
	case operation.OpTypeAfterFind:
		c.afterFind = c.remove(c.afterFind, name)
	}
}

func (c *Callback) remove(callbackHandlers []callbackHandler, name string) []callbackHandler {
	for i, handler := range callbackHandlers {
		if handler.name == name {
			callbackHandlers = append(callbackHandlers[:i], callbackHandlers[i+1:]...)
			break
		}
	}
	return callbackHandlers
}

type callbackHandler struct {
	name string
	fn   CbFn
}
