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
	"go.mongodb.org/mongo-driver/bson"
)

const (
	id  = "_id"
	set = "$set"
)

type BsonBuilder struct {
	data bson.D
}

func NewBsonBuilder() *BsonBuilder {
	return &BsonBuilder{data: bson.D{}}
}

func (b *BsonBuilder) Id(v any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: id, Value: v})
	return b
}

func (b *BsonBuilder) Add(k string, v any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: k, Value: v})
	return b
}

func (b *BsonBuilder) Set(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: set, Value: bson.D{bson.E{Key: key, Value: value}}})
	return b
}

func (b *BsonBuilder) SetForMap(data map[string]any) *BsonBuilder {
	if d := MapToBson(data); len(d) != 0 {
		b.data = append(b.data, bson.E{Key: set, Value: d})
	}
	return b
}

func (b *BsonBuilder) SetForStruct(data any) *BsonBuilder {
	if d := StructToBson(data); len(d) != 0 {
		b.data = append(b.data, bson.E{Key: set, Value: d})
	}
	return b
}

func (b *BsonBuilder) Build() bson.D {
	return b.data
}
