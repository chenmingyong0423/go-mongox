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

package converter

import (
	"reflect"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/types"

	"go.mongodb.org/mongo-driver/bson"
)

func ToBson(data any) bson.D {
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
		return structToBson(data)
	case reflect.Ptr:
		elemVal := val.Elem()
		if elemVal.Kind() == reflect.Struct {
			return structToBson(elemVal.Interface())
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

func MapToBson[T any](data map[string]T) bson.D {
	if data == nil {
		return nil
	}
	d := bson.D{}
	for k, v := range data {
		isMap := utils.IsMap(v)
		if !isMap {
			d = append(d, bson.E{Key: k, Value: v})
		} else {
			d = append(d, bson.E{Key: k, Value: ToBson(v)})
		}
	}
	return d
}

func structToBson(val any) bson.D {
	marshal, err := bson.Marshal(val)
	if err != nil {
		return nil
	}
	var d bson.D
	_ = bson.Unmarshal(marshal, &d)
	return d
}

func StructToBson(data any) (d bson.D) {
	if data == nil {
		return
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Struct {
		return structToBson(data)
	} else if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		return structToBson(val.Elem().Interface())
	}
	return
}

func MapToSetBson(data map[string]any) bson.D {
	if data == nil {
		return nil
	}
	return bson.D{bson.E{Key: types.Set, Value: MapToBson(data)}}
}

func StructToSetBson(data any) (d bson.D) {
	if data == nil {
		return
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Struct {
		d = structToBson(data)
	} else if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		d = structToBson(val.Elem().Interface())
	}
	if len(d) == 0 {
		return nil
	}
	return bson.D{bson.E{Key: types.Set, Value: d}}
}

func ToSetBson(updates any) bson.D {
	val := reflect.ValueOf(updates)
	kind := val.Kind()
	switch kind {
	case reflect.Map:
		if val.Type().Key().Kind() == reflect.String && val.Type().Elem().Kind() == reflect.Interface {
			return MapToSetBson(updates.(map[string]any))
		}
	case reflect.Struct:
		d := structToBson(updates)
		if len(d) != 0 {
			return bson.D{bson.E{Key: types.Set, Value: d}}
		}
		return d
	case reflect.Ptr:
		elemVal := val.Elem()
		if elemVal.Kind() == reflect.Struct {
			d := structToBson(elemVal.Interface())
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

func KeyValue(key string, value any) types.KeyValue {
	return types.KeyValue{Key: key, Value: value}
}
