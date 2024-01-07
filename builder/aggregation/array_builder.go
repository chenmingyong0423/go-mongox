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

package aggregation

import (
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type arrayBuilder struct {
	parent *Builder
}

func (b *arrayBuilder) ArrayElemAt(key string, expression any, index int64) *Builder {
	e := bson.E{Key: types.AggregationArrayElemAt, Value: []any{expression, index}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) ArrayElemAtWithoutKey(expression any, index int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationArrayElemAt, Value: []any{expression, index}})
	return b.parent
}

func (b *arrayBuilder) ConcatArrays(key string, arrays ...any) *Builder {
	e := bson.E{Key: types.AggregationConcatArrays, Value: arrays}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) ConcatArraysWithoutKey(arrays ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationConcatArrays, Value: arrays})
	return b.parent
}

func (b *arrayBuilder) ArrayToObject(key string, expression any) *Builder {
	e := bson.E{Key: types.AggregationArrayToObject, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) ArrayToObjectWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationArrayToObject, Value: expression})
	return b.parent
}

func (b *arrayBuilder) Size(key string, expression any) *Builder {
	e := bson.E{Key: types.AggregationSize, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) SizeWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSize, Value: expression})
	return b.parent
}

func (b *arrayBuilder) Slice(key string, array any, nElements int64) *Builder {
	e := bson.E{Key: types.AggregationSlice, Value: []any{array, nElements}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) SliceWithoutKey(array any, nElements int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSlice, Value: []any{array, nElements}})
	return b.parent
}

func (b *arrayBuilder) SliceWithPosition(key string, array any, position, nElements int64) *Builder {
	e := bson.E{Key: types.AggregationSlice, Value: []any{array, position, nElements}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) SliceWithPositionWithoutKey(array any, position, nElements int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSlice, Value: []any{array, position, nElements}})
	return b.parent
}

func (b *arrayBuilder) Map(key string, inputArray any, as string, in any) *Builder {
	e := bson.E{Key: types.AggregationMap, Value: bson.D{
		{Key: types.AggregationInput, Value: inputArray},
		{Key: types.AggregationAs, Value: as},
		{Key: types.AggregationIn, Value: in},
	}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) MapWithoutKey(inputArray any, as string, in any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMap, Value: bson.D{
		{Key: types.AggregationInput, Value: inputArray},
		{Key: types.AggregationAs, Value: as},
		{Key: types.AggregationIn, Value: in},
	}})
	return b.parent
}

func (b *arrayBuilder) Filter(key string, inputArray any, cond any, opt *types.FilterOptions) *Builder {
	d := bson.D{{Key: types.AggregationInput, Value: inputArray}, {Key: types.AggregationCondWithoutOperator, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: types.AggregationAs, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: types.AggregationLimit, Value: opt.Limit})
		}
	}
	e := bson.E{Key: types.AggregationFilter, Value: d}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayBuilder) FilterWithoutKey(inputArray any, cond any, opt *types.FilterOptions) *Builder {
	d := bson.D{{Key: types.AggregationInput, Value: inputArray}, {Key: types.AggregationCondWithoutOperator, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: types.AggregationAs, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: types.AggregationLimit, Value: opt.Limit})
		}
	}
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationFilter, Value: d})
	return b.parent
}
