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

type comparisonBuilder struct {
	parent *Builder
}

func (b *comparisonBuilder) Eq(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationEq, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) EqWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationEq, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Ne(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationNe, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) NeWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationNe, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Gt(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationGt, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) GtWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationGt, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Gte(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationGte, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) GteWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationGte, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Lt(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationLt, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) LtWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationLt, Value: expressions})
	return b.parent
}

func (b *comparisonBuilder) Lte(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationLte, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *comparisonBuilder) LteWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationLte, Value: expressions})
	return b.parent
}
