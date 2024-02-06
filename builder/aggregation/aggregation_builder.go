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
	"go.mongodb.org/mongo-driver/bson"
)

func BsonBuilder() *Builder {
	b := &Builder{d: bson.D{}}

	b.arithmeticBuilder = arithmeticBuilder{parent: b}
	b.comparisonBuilder = comparisonBuilder{parent: b}
	b.logicalBuilder = logicalBuilder{parent: b}
	b.stringBuilder = stringBuilder{parent: b}
	b.arrayBuilder = arrayBuilder{parent: b}
	b.dateBuilder = dateBuilder{parent: b}
	b.condBuilder = condBuilder{parent: b}
	b.accumulatorsBuilder = accumulatorsBuilder{parent: b}

	return b
}

type Builder struct {
	arithmeticBuilder
	comparisonBuilder
	logicalBuilder
	stringBuilder
	arrayBuilder
	dateBuilder
	condBuilder
	accumulatorsBuilder

	d bson.D
}

func (b *Builder) Build() bson.D {
	return b.d
}

func (b *Builder) AddKeyValues(key string, value any) *Builder {
	b.d = append(b.d, bson.E{Key: key, Value: value})
	return b
}

// tryMergeValue attempts to merge the provided bson.E elements into an existing bson.D element
// in the builder's data slice, identified by the specified key.
func (b *Builder) tryMergeValue(key string, e ...bson.E) bool {
	for idx, datum := range b.d {
		if datum.Key == key {
			if m, ok := datum.Value.(bson.D); ok {
				m = append(m, e...)
				b.d[idx].Value = m
				return true
			}
		}
	}
	return false
}
