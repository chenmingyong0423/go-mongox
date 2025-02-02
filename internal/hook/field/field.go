// Copyright 2025 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package field

import (
	"context"
	"reflect"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/operation"
)

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	switch opType {
	case operation.OpTypeBeforeInsert:
		valueOf := opCtx.ReflectValue

		if !valueOf.IsValid() {
			return nil
		}

		switch valueOf.Type().Kind() {
		case reflect.Slice:
			return executeSlice(ctx, valueOf, opType, opCtx.StartTime, opCtx.Fields, opts...)
		case reflect.Ptr:
			if valueOf.IsZero() {
				return nil
			}
			return execute(ctx, valueOf, opType, opCtx.StartTime, opCtx.Fields, opts...)
		default:
			return nil
		}
	case operation.OpTypeBeforeUpdate, operation.OpTypeBeforeUpsert:
		return execute(ctx, opCtx.Updates, opType, opCtx.StartTime, opCtx.Fields, opts...)
	}
	return nil
}

func executeSlice(ctx context.Context, docs reflect.Value, opType operation.OpType, currentTime time.Time, fields []*field.Filed, opts ...any) error {
	for i := 0; i < docs.Len(); i++ {
		doc := docs.Index(i)
		if err := execute(ctx, doc, opType, currentTime, fields, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(_ context.Context, dest any, opType operation.OpType, currentTime time.Time, fields []*field.Filed, opts ...any) error {
	return strategies[opType](dest, currentTime, fields, opts...)
}
