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

package mongox

import (
	"reflect"
	"strings"

	"github.com/chenmingyong0423/go-mongox/internal/types"

	"go.mongodb.org/mongo-driver/bson"
)

func toBson(data any) bson.D {
	if data == nil {
		return nil
	}
	val := reflect.ValueOf(data)
	kind := val.Kind()
	switch kind {
	case reflect.Map:
		if val.Type().Key().Kind() == reflect.String && val.Type().Elem().Kind() == reflect.Interface {
			return MapToBson(data.(map[string]any))
		}
	case reflect.Struct:
		return structToBson(val)
	case reflect.Ptr:
		elemVal := val.Elem()
		if elemVal.Kind() == reflect.Struct {
			return structToBson(elemVal)
		} else if elemVal.Kind() == reflect.Map && elemVal.Type().Key().Kind() == reflect.String && elemVal.Type().Elem().Kind() == reflect.Interface {
			return MapToBson(elemVal.Interface().(map[string]any))
		}
	default:
		if d, ok := data.(bson.D); ok {
			return d
		}
	}
	return nil
}

func MapToBson(data map[string]any) bson.D {
	d := bson.D{}
	for k, v := range data {
		d = append(d, bson.E{Key: k, Value: v})
	}
	return d
}

func structToBson(val reflect.Value) (d bson.D) {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		valueField := val.Field(i)

		bsonTag := field.Tag.Get("bson")
		if bsonTag == "-" {
			continue // Ignore fields with bson tag "-"
		}

		fieldName := bsonTag
		if fieldName == "" {
			fieldName = strings.ToLower(field.Name)
		}
		d = append(d, bson.E{Key: fieldName, Value: valueField.Interface()})
	}
	return
}

func StructToBson(data any) (d bson.D) {
	if data == nil {
		return
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}
	return structToBson(val)
}

func MapToSetBson(data map[string]any) (d bson.D) {
	if d = MapToBson(data); len(d) != 0 {
		return bson.D{bson.E{Key: types.Set, Value: d}}
	}
	return
}

func StructToSetBson(data any) bson.D {
	if data == nil {
		return nil
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	if d := structToBson(val); len(d) != 0 {
		return bson.D{bson.E{Key: types.Set, Value: d}}
	}
	return nil
}

func toSetBson(updates any) bson.D {
	val := reflect.ValueOf(updates)
	kind := val.Kind()
	switch kind {
	case reflect.Map:
		if val.Type().Key().Kind() == reflect.String && val.Type().Elem().Kind() == reflect.Interface {
			return MapToSetBson(updates.(map[string]any))
		}
	case reflect.Struct:
		d := structToBson(val)
		if len(d) != 0 {
			return bson.D{bson.E{Key: types.Set, Value: d}}
		}
		return d
	case reflect.Ptr:
		elemVal := val.Elem()
		if elemVal.Kind() == reflect.Struct {
			d := structToBson(elemVal)
			if len(d) != 0 {
				return bson.D{bson.E{Key: types.Set, Value: d}}
			}
		} else if elemVal.Kind() == reflect.Map && elemVal.Type().Key().Kind() == reflect.String && elemVal.Type().Elem().Kind() == reflect.Interface {
			return MapToSetBson(elemVal.Interface().(map[string]any))
		}
	default:
		if d, ok := updates.(bson.D); ok {
			return d
		}
	}
	return nil
}
