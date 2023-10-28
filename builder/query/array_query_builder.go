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

type arrayQueryBuilder struct {
	parent *QueryBuilder
}

func (b *arrayQueryBuilder) All(key string, values ...any) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.All: values}})
	return b.parent
}

func (b *arrayQueryBuilder) AllUint(key string, values ...uint) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllUint8(key string, values ...uint8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllUint16(key string, values ...uint16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllUint32(key string, values ...uint32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllUint64(key string, values ...uint64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllInt(key string, values ...int) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllInt8(key string, values ...int8) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllInt16(key string, values ...int16) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllInt32(key string, values ...int32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllInt64(key string, values ...int64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllString(key string, values ...string) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllFloat32(key string, values ...float32) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) AllFloat64(key string, values ...float64) *QueryBuilder {
	valuesAny := pkg.ToAnySlice(values...)
	b.All(key, valuesAny...)
	return b.parent
}

func (b *arrayQueryBuilder) ElemMatch(key string, condition bson.D) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.ElemMatch: condition}})
	return b.parent
}

func (b *arrayQueryBuilder) Size(key string, size int) *QueryBuilder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.M{types.Size: size}})
	return b.parent
}
