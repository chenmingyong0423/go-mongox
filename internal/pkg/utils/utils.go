// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func ToAnySlice[T any](values ...T) []any {
	if values == nil {
		return nil
	}
	valuesAny := make([]any, len(values))
	for i, v := range values {
		valuesAny[i] = v
	}
	return valuesAny
}

// EqualBSONDElements 比较两个 bson.D 结构的元素是否一致，不考虑顺序
func EqualBSONDElements(d1, d2 bson.D) bool {
	// 如果长度不同，它们不相等
	if len(d1) != len(d2) {
		fmt.Printf("Not equal: \n"+
			"expected: %#v\n"+
			"actual  : %#v\n", d1, d2)
		return false
	}

	// 创建 map 用于存储元素的键值对
	elementsMap1 := make(map[string]interface{})

	// 将元素存储在 map 中
	for _, e := range d1 {
		elementsMap1[e.Key] = e.Value
	}

	for _, e := range d2 {
		v, ok := elementsMap1[e.Key]
		if !ok {
			fmt.Printf("Not equal: \n"+
				"expected: %#v\n"+
				"actual  : %#v\n", d1, d2)
			return false
		}
		if bv, ok := v.(bson.D); ok {
			return EqualBSONDElements(bv, e.Value.(bson.D))
		}
		if !reflect.DeepEqual(e.Value, v) {
			fmt.Printf("Not equal: \n"+
				"expected: %#v\n"+
				"actual  : %#v\n", d1, d2)
			return false
		}
	}
	return true
}

func EqualPipelineElements(p1, p2 mongo.Pipeline) bool {
	// 如果长度不同，它们不相等
	if len(p1) != len(p2) {
		return false
	}

	for i := range p1 {
		if !EqualBSONDElements(p1[i], p2[i]) {
			return false
		}
	}
	return true
}

func IsNumeric(value any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}
