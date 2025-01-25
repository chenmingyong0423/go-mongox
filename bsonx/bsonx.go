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

package bsonx

import (
	"bytes"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func M(key string, value any) bson.M {
	return bson.M{key: value}
}

func E(key string, value any) bson.E {
	return bson.E{Key: key, Value: value}
}

func A(values ...any) bson.A {
	value := make(bson.A, 0, len(values))
	for _, v := range values {
		value = append(value, v)
	}
	return value
}

func D(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: value}}
}

func Id(value any) bson.M {
	return M("_id", value)
}

func ToBsonM(data any) bson.M {
	if data == nil {
		return nil
	}
	if d, ok := data.(bson.M); ok {
		return d
	}

	if d, ok := data.(bson.D); ok {
		return dToM(d)
	}

	if d, ok := data.(map[string]any); ok {
		return MapToBsonM(d)
	}

	if d, ok := data.(*map[string]any); ok && d != nil {
		return MapToBsonM(*d)
	}

	return nil
}

func MapToBsonM(data map[string]any) bson.M {
	m := bson.M{}
	for k, v := range data {
		m[k] = v
	}
	return m
}

func dToM(d bson.D) bson.M {
	marshal, err := bson.Marshal(d)
	if err != nil {
		return nil
	}
	var m bson.M
	decoder := bson.NewDecoder(bson.NewDocumentReader(bytes.NewReader(marshal)))
	decoder.DefaultDocumentM()
	err = decoder.Decode(&m)
	if err != nil {
		return nil
	}
	return m
}

// StringSortToBsonD transform string sort to bson D
// "-created_at" => bson.D{{"created_at", -1}}
// []string{"age", "-created_at"} => bson.D{{"age", 1}, {"created_at", -1}}
func StringSortToBsonD(sorts ...string) bson.D {
	var res bson.D
	for _, sort := range sorts {
		if len(sort) == 0 || sort == "-" {
			continue
		}

		if sort[0] == '-' {
			res = append(res, bson.E{Key: sort[1:], Value: -1})
		} else {
			res = append(res, bson.E{Key: sort, Value: 1})
		}
	}

	return res
}
