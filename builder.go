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
	"github.com/chenmingyong0423/go-mongox/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func NewBsonBuilder() *BsonBuilder {
	return &BsonBuilder{data: bson.D{}}
}

type BsonBuilder struct {
	data bson.D
}

func (b *BsonBuilder) Build() bson.D {
	return b.data
}

func (b *BsonBuilder) Id(v any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.Id, Value: v})
	return b
}

func (b *BsonBuilder) Add(k string, v any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: k, Value: v})
	return b
}

func (b *BsonBuilder) Set(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.Set, Value: bson.D{bson.E{Key: key, Value: value}}})
	return b
}

func (b *BsonBuilder) SetForMap(data map[string]any) *BsonBuilder {
	if d := MapToBson(data); len(d) != 0 {
		b.data = append(b.data, bson.E{Key: types.Set, Value: d})
	}
	return b
}

func (b *BsonBuilder) SetForStruct(data any) *BsonBuilder {
	if d := StructToBson(data); len(d) != 0 {
		b.data = append(b.data, bson.E{Key: types.Set, Value: d})
	}
	return b
}

func (b *BsonBuilder) In(key string, values ...any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.In: values}})
	return b
}

func (b *BsonBuilder) InUint(key string, values ...uint) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InUint8(key string, values ...uint8) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InUint16(key string, values ...uint16) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InUint32(key string, values ...uint32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InUint64(key string, values ...uint64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InInt(key string, values ...int) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InInt8(key string, values ...int8) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InInt16(key string, values ...int16) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InInt32(key string, values ...int32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InInt64(key string, values ...int64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InString(key string, values ...string) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InFloat32(key string, values ...float32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) InFloat64(key string, values ...float64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.In(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinUint(key string, values ...uint) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinUint8(key string, values ...uint8) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinUint16(key string, values ...uint16) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinUint32(key string, values ...uint32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinUint64(key string, values ...uint64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinInt(key string, values ...int) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinInt8(key string, values ...int8) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinInt16(key string, values ...int16) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinInt32(key string, values ...int32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinInt64(key string, values ...int64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinString(key string, values ...string) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinFloat32(key string, values ...float32) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func (b *BsonBuilder) NinFloat64(key string, values ...float64) *BsonBuilder {
	valuesAny := toAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b
}

func toAnySlice[T any](values ...T) []any {
	if values == nil {
		return nil
	}
	valuesAny := make([]any, len(values))
	for i, v := range values {
		valuesAny[i] = v
	}
	return valuesAny
}

func (b *BsonBuilder) Eq(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Eq: value}})
	return b
}

func (b *BsonBuilder) Gt(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Gt: value}})
	return b
}

func (b *BsonBuilder) Gte(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Gte: value}})
	return b
}

func (b *BsonBuilder) Lt(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Lt: value}})
	return b
}

func (b *BsonBuilder) Lte(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Lte: value}})
	return b
}

func (b *BsonBuilder) Ne(key string, value any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Ne: value}})
	return b
}

func (b *BsonBuilder) Nin(key string, values ...any) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Nin: values}})
	return b
}

// And
// 对于 conditions 参数，你同样可以使用 BsonBuilder 去生成
func (b *BsonBuilder) And(conditions ...bson.D) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.And, Value: conditions})
	return b
}

func (b *BsonBuilder) Not(condition bson.D) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.Not, Value: condition})
	return b
}

// Nor
// 对于 conditions 参数，你同样可以使用 BsonBuilder 去生成
func (b *BsonBuilder) Nor(conditions ...bson.D) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.Nor, Value: conditions})
	return b
}

// Or
// 对于 conditions 参数，你同样可以使用 BsonBuilder 去生成
func (b *BsonBuilder) Or(conditions ...bson.D) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: types.Or, Value: conditions})
	return b
}

func (b *BsonBuilder) Exists(key string, exists bool) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Exists: exists}})
	return b
}

func (b *BsonBuilder) Type(key string, t bsontype.Type) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Type: t}})
	return b
}

func (b *BsonBuilder) TypeAlias(key string, alias string) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Type: alias}})
	return b
}

func (b *BsonBuilder) TypeArray(key string, ts ...bsontype.Type) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Type: ts}})
	return b
}

func (b *BsonBuilder) TypeArrayAlias(key string, aliases ...string) *BsonBuilder {
	b.data = append(b.data, bson.E{Key: key, Value: bson.M{types.Type: aliases}})
	return b
}
