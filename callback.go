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

	validator2 "github.com/go-playground/validator/v10"

	"github.com/chenmingyong0423/go-mongox/v2/hook/validator"

	"github.com/chenmingyong0423/go-mongox/v2/hook/model"

	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/hook/field"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
)

func RegisterPlugin(name string, cb callback.CbFn, opType operation.OpType) {
	callback.Callbacks.Register(opType, name, cb)
}

func RemovePlugin(name string, opType operation.OpType) {
	callback.Callbacks.Remove(opType, name)
}

type PluginConfig struct {
	EnableDefaultFieldHook bool
	EnableModelHook        bool
	EnableValidationHook   bool
	// use to replace to the default validate instance
	Validate *validator2.Validate
}

func InitPlugin(config *PluginConfig) {
	if config.EnableDefaultFieldHook {
		opTypes := []operation.OpType{operation.OpTypeBeforeInsert, operation.OpTypeBeforeUpdate, operation.OpTypeBeforeUpsert}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:default_field", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return field.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
	if config.EnableModelHook {
		opTypes := []operation.OpType{
			operation.OpTypeBeforeInsert, operation.OpTypeAfterInsert,
			operation.OpTypeBeforeDelete, operation.OpTypeAfterDelete,
			operation.OpTypeBeforeUpdate, operation.OpTypeAfterUpdate,
			operation.OpTypeBeforeUpsert, operation.OpTypeAfterUpsert,
			operation.OpTypeBeforeFind, operation.OpTypeAfterFind,
		}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:model", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return model.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
	if config.EnableValidationHook {
		validator.SetValidate(config.Validate)
		opTypes := []operation.OpType{operation.OpTypeBeforeInsert, operation.OpTypeBeforeUpsert}
		for _, opType := range opTypes {
			typ := opType
			RegisterPlugin("mongox:validation", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
				return validator.Execute(ctx, opCtx, typ, opts...)
			}, typ)
		}
	}
}
