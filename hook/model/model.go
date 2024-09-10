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

import (
	"context"
	"reflect"

	"github.com/chenmingyong0423/go-mongox/operation"
)

func getPayload(opCtx *operation.OpContext, opType operation.OpType) any {
	if opCtx == nil {
		return nil
	}
	if opCtx.ModelHook != nil {
		return opCtx.ModelHook
	}

	switch opType {
	case operation.OpTypeBeforeInsert, operation.OpTypeAfterInsert, operation.OpTypeAfterFind:
		return opCtx.Doc
	case operation.OpTypeBeforeUpdate, operation.OpTypeAfterUpdate, operation.OpTypeBeforeUpsert, operation.OpTypeAfterUpsert:
		return opCtx.Updates
	default:
		return opCtx.ModelHook
	}
}

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	payLoad := getPayload(opCtx, opType)
	if payLoad == nil {
		return nil
	}
	valueOf := reflect.ValueOf(payLoad)

	switch valueOf.Type().Kind() {
	case reflect.Slice:
		return executeSlice(ctx, valueOf, opType, opts...)
	case reflect.Ptr:
		if valueOf.IsZero() {
			return nil
		}
		return execute(ctx, payLoad, opType, opts...)
	default:
		return nil
	}
}

func executeSlice(ctx context.Context, docs reflect.Value, opType operation.OpType, opts ...any) error {
	for i := 0; i < docs.Len(); i++ {
		doc := docs.Index(i)
		if err := execute(ctx, doc.Interface(), opType, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(ctx context.Context, doc any, opType operation.OpType, _ ...any) error {
	if doc == nil {
		return nil
	}
	switch opType {
	case operation.OpTypeBeforeInsert:
		if m, ok := doc.(BeforeInsert); ok {
			return m.BeforeInsert(ctx)
		}
	case operation.OpTypeAfterInsert:
		if m, ok := doc.(AfterInsert); ok {
			return m.AfterInsert(ctx)
		}
	case operation.OpTypeBeforeDelete:
		if m, ok := doc.(BeforeDelete); ok {
			return m.BeforeDelete(ctx)
		}
	case operation.OpTypeAfterDelete:
		if m, ok := doc.(AfterDelete); ok {
			return m.AfterDelete(ctx)
		}
	case operation.OpTypeBeforeUpdate:
		if m, ok := doc.(BeforeUpdate); ok {
			return m.BeforeUpdate(ctx)
		}
	case operation.OpTypeAfterUpdate:
		if m, ok := doc.(AfterUpdate); ok {
			return m.AfterUpdate(ctx)
		}
	case operation.OpTypeBeforeUpsert:
		if m, ok := doc.(BeforeUpsert); ok {
			return m.BeforeUpsert(ctx)
		}
	case operation.OpTypeAfterUpsert:
		if m, ok := doc.(AfterUpsert); ok {
			return m.AfterUpsert(ctx)
		}
	case operation.OpTypeBeforeFind:
		if m, ok := doc.(BeforeFind); ok {
			return m.BeforeFind(ctx)
		}
	case operation.OpTypeAfterFind:
		if m, ok := doc.(AfterFind); ok {
			return m.AfterFind(ctx)
		}
	}
	return nil
}
