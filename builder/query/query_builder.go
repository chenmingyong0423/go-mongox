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

func BsonBuilder() *Builder {
	query := &Builder{
		data: bson.D{},
		err:  make([]error, 0),
	}
	query.comparisonQueryBuilder = comparisonQueryBuilder{parent: query}
	query.logicalQueryBuilder = logicalQueryBuilder{parent: query}
	query.elementQueryBuilder = elementQueryBuilder{parent: query}
	query.arrayQueryBuilder = arrayQueryBuilder{parent: query}
	query.evaluationQueryBuilder = evaluationQueryBuilder{parent: query}
	query.projectionQueryBuilder = projectionQueryBuilder{parent: query}
	return query
}

type Builder struct {
	data bson.D
	comparisonQueryBuilder
	logicalQueryBuilder
	elementQueryBuilder
	arrayQueryBuilder
	evaluationQueryBuilder
	projectionQueryBuilder

	err []error
}

func (b *Builder) Build() bson.D {
	return b.data
}

// Id appends an element with '_id' key and given value to the builder's data slice.
func (b *Builder) Id(v any) *Builder {
	b.data = append(b.data, bson.E{Key: types.Id, Value: v})
	return b
}

// Add appends given types.KeyValue elements to the builder's data slice. Each types.KeyValue
// is converted into a bson.E before appending.
func (b *Builder) Add(key string, value any) *Builder {
	b.data = append(b.data, bson.E{Key: key, Value: value})
	return b
}

// tryMergeValue attempts to merge the provided bson.E elements into an existing bson.D element
// in the builder's data slice, identified by the specified key.
func (b *Builder) tryMergeValue(key string, e ...bson.E) bool {
	for idx, datum := range b.data {
		if datum.Key == key {
			if m, ok := datum.Value.(bson.D); ok {
				m = append(m, e...)
				b.data[idx].Value = m
				return true
			}
		}
	}
	return false
}
