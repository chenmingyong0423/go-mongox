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

func (b *arrayBuilder) ArrayElemAt(expression any, index int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationArrayElemAt, Value: []any{expression, index}})
	return b.parent
}

func (b *arrayBuilder) ConcatArrays(arrays ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationConcatArrays, Value: arrays})
	return b.parent
}

func (b *arrayBuilder) ArrayToObject(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationArrayToObject, Value: expression})
	return b.parent
}

func (b *arrayBuilder) Size(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSize, Value: expression})
	return b.parent
}

func (b *arrayBuilder) Slice(array any, nElements int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSlice, Value: []any{array, nElements}})
	return b.parent
}

func (b *arrayBuilder) SliceWithPosition(array any, position, nElements int64) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSlice, Value: []any{array, position, nElements}})
	return b.parent
}
