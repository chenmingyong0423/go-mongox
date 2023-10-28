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

package builder

import (
	"github.com/chenmingyong0423/go-mongox/pkg"
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type comparisonQueryBuilder struct {
	parent *QueryBuilder
}

func (b *comparisonQueryBuilder) Eq(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Eq: value}})
	return b.parent
}

func (b *comparisonQueryBuilder) Gt(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Gt: value}})
	return b.parent
}

func (b *comparisonQueryBuilder) Gte(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Gte: value}})
	return b.parent
}

func (b *comparisonQueryBuilder) In(key string, values ...any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.In: values}})
	return b.parent
}

func (b *comparisonQueryBuilder) InUint(key string, values ...uint) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InUint8(key string, values ...uint8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InUint16(key string, values ...uint16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InUint32(key string, values ...uint32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InUint64(key string, values ...uint64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InInt(key string, values ...int) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InInt8(key string, values ...int8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InInt16(key string, values ...int16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InInt32(key string, values ...int32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InInt64(key string, values ...int64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InString(key string, values ...string) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InFloat32(key string, values ...float32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) InFloat64(key string, values ...float64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.In(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) Nin(key string, values ...any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Nin: values}})
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint(key string, values ...uint) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint8(key string, values ...uint8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint16(key string, values ...uint16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint32(key string, values ...uint32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint64(key string, values ...uint64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt(key string, values ...int) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt8(key string, values ...int8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt16(key string, values ...int16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt32(key string, values ...int32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt64(key string, values ...int64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinString(key string, values ...string) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinFloat32(key string, values ...float32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) NinFloat64(key string, values ...float64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.Nin(key, valuesAny...)
	return b.parent
}

func (b *comparisonQueryBuilder) Lt(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Lt: value}})
	return b.parent
}

func (b *comparisonQueryBuilder) Lte(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Lte: value}})
	return b.parent
}

func (b *comparisonQueryBuilder) Ne(key string, value any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Ne: value}})
	return b.parent
}
