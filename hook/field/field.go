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

	"github.com/chenmingyong0423/go-mongox/operation"
)

func Execute(ctx context.Context, doc any, opType operation.OpType, opts ...any) error {
	typeOf := reflect.TypeOf(doc)
	if typeOf == nil {
		return nil
	}
	switch typeOf.Kind() {
	case reflect.Slice:
		return executeSlice(ctx, doc, opType, opts...)
	case reflect.Ptr:
		return execute(ctx, doc, opType, opts...)
	}
	return nil
}

func executeSlice(ctx context.Context, docs any, opType operation.OpType, opts ...any) error {
	sliceValue := reflect.ValueOf(docs)
	for i := 0; i < sliceValue.Len(); i++ {
		doc := sliceValue.Index(i)
		if err := execute(ctx, doc.Interface(), opType, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(_ context.Context, doc any, opType operation.OpType, _ ...any) error {
	if strategy, ok := strategies[opType]; ok {
		return strategy(doc)
	}
	return nil
}
