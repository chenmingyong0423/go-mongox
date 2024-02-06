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

package query

import (
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type comparisonQueryBuilder struct {
	parent *Builder
}

func (b *comparisonQueryBuilder) Eq(key string, value any) *Builder {
	e := bson.E{Key: types.Eq, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Gt(key string, value any) *Builder {
	e := bson.E{Key: types.Gt, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Gte(key string, value any) *Builder {
	e := bson.E{Key: types.Gte, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) In(key string, values ...any) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InUint(key string, values ...uint) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InUint8(key string, values ...uint8) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InUint16(key string, values ...uint16) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InUint32(key string, values ...uint32) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InUint64(key string, values ...uint64) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InInt(key string, values ...int) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InInt8(key string, values ...int8) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InInt16(key string, values ...int16) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InInt32(key string, values ...int32) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InInt64(key string, values ...int64) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InString(key string, values ...string) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InFloat32(key string, values ...float32) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) InFloat64(key string, values ...float64) *Builder {
	e := bson.E{Key: types.In, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Nin(key string, values ...any) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint(key string, values ...uint) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint8(key string, values ...uint8) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint16(key string, values ...uint16) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint32(key string, values ...uint32) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinUint64(key string, values ...uint64) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt(key string, values ...int) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt8(key string, values ...int8) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt16(key string, values ...int16) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt32(key string, values ...int32) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinInt64(key string, values ...int64) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinString(key string, values ...string) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinFloat32(key string, values ...float32) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) NinFloat64(key string, values ...float64) *Builder {
	e := bson.E{Key: types.Nin, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Lt(key string, value any) *Builder {
	e := bson.E{Key: types.Lt, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Lte(key string, value any) *Builder {
	e := bson.E{Key: types.Lte, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonQueryBuilder) Ne(key string, value any) *Builder {
	e := bson.E{Key: types.Ne, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
