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
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

func M(key string, value any) bson.M {
	return bson.M{key: value}
}

func D(bsonElements ...types.KeyValue) bson.D {
	value := bson.D{}
	for _, element := range bsonElements {
		value = append(value, bson.E{Key: element.Key, Value: element.Value})
	}
	return value
}

func KV(key string, value any) types.KeyValue {
	return types.KeyValue{Key: key, Value: value}
}

func Id(value any) bson.M {
	return M("_id", value)
}

func KVsToBson(bsonElements ...types.KeyValue) bson.D {
	value := bson.D{}
	for _, element := range bsonElements {
		value = append(value, bson.E{Key: element.Key, Value: element.Value})
	}
	return value
}
