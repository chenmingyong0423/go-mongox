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

type logicalBuilder struct {
	parent *Builder
}

func (b *logicalBuilder) And(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationAnd, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) AndWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationAnd, Value: expressions})
	return b.parent
}

func (b *logicalBuilder) Not(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationNot, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) NotWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationNot, Value: expressions})
	return b.parent
}

func (b *logicalBuilder) Or(key string, expressions ...any) *Builder {
	e := bson.E{Key: types.AggregationOr, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *logicalBuilder) OrWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationOr, Value: expressions})
	return b.parent
}
