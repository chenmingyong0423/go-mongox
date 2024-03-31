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

package validator

import (
	"context"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/chenmingyong0423/go-mongox/operation"
)

var validate = validator.New()

func getPayload(opCtx *operation.OpContext, opType operation.OpType) any {
	switch opType {
	case operation.OpTypeBeforeInsert:
		return opCtx.Doc
	case operation.OpTypeBeforeUpsert:
		return opCtx.Replacement
	default:
		return nil
	}
}

func Execute(ctx context.Context, opCtx *operation.OpContext, opType operation.OpType, opts ...any) error {
	payLoad := getPayload(opCtx, opType)
	if payLoad == nil {
		return nil
	}
	value := reflect.ValueOf(payLoad)
	if value.IsZero() {
		return nil
	}
	switch value.Type().Kind() {
	case reflect.Slice:
		return executeSlice(ctx, value, opts...)
	case reflect.Ptr:
		return execute(ctx, value, opts...)
	default:
		return nil
	}
}

func executeSlice(ctx context.Context, docs reflect.Value, opts ...any) error {
	for i := 0; i < docs.Len(); i++ {
		doc := docs.Index(i)
		if err := execute(ctx, doc, opts...); err != nil {
			return err
		}
	}
	return nil
}

func execute(ctx context.Context, value reflect.Value, _ ...any) error {
	doc := validateStruct(value)
	if doc == nil {
		return nil
	}
	return validate.StructCtx(ctx, doc)
}

func validateStruct(doc reflect.Value) any {
	if doc.Kind() == reflect.Pointer && !doc.IsNil() {
		doc = doc.Elem()
	}
	if doc.Kind() != reflect.Struct || doc.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
		return nil
	}
	return doc.Interface()
}
