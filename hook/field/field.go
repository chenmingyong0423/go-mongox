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

package field

import (
	"context"
	"reflect"

	"github.com/chenmingyong0423/go-mongox/v2/operation"
)

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	doc := opCtx.Doc
	if doc == nil {
		return nil
	}
	valueOf := reflect.ValueOf(doc)
	opts = append([]any{opCtx.Updates, opCtx.Replacement, opCtx.MongoOptions}, opts...)
	switch valueOf.Type().Kind() {
	case reflect.Slice:
		return executeSlice(ctx, valueOf, opType, opts...)
	case reflect.Ptr:
		if valueOf.IsZero() {
			return nil
		}
		return execute(ctx, doc, opType, opts...)
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

func execute(_ context.Context, doc any, opType operation.OpType, opts ...any) error {
	if strategy, ok := strategies[opType]; ok {
		return strategy(doc, opts...)
	}
	return nil
}
