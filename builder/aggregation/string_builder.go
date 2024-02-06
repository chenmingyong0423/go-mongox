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

type stringBuilder struct {
	parent *Builder
}

func (b *stringBuilder) Concat(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationConcat, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ConcatWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationConcat, Value: expressions})
	return b.parent
}

func (b *stringBuilder) SubstrBytes(key string, stringExpression string, byteIndex int64, byteCount int64) *Builder {
	e := bson.E{Key: types.AggregationSubstrBytes, Value: []any{stringExpression, byteIndex, byteCount}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) SubstrBytesWithoutKey(stringExpression string, byteIndex int64, byteCount int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSubstrBytes, Value: []any{stringExpression, byteIndex, byteCount}})
	return b.parent
}

func (b *stringBuilder) ToLower(key string, expression any) *Builder {
	e := bson.E{Key: types.AggregationToLower, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ToLowerWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationToLower, Value: expression})
	return b.parent
}

func (b *stringBuilder) ToUpper(key string, expression any) *Builder {
	e := bson.E{Key: types.AggregationToUpper, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ToUpperWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationToUpper, Value: expression})
	return b.parent
}

func (b *stringBuilder) Contact(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationContact, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *stringBuilder) ContactWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationContact, Value: expressions})
	return b.parent
}
