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

func BsonBuilder() *Builder {
	b := &Builder{d: bson.D{}}

	b.arithmeticBuilder = arithmeticBuilder{parent: b}
	b.comparisonBuilder = comparisonBuilder{parent: b}
	b.logicalBuilder = logicalBuilder{parent: b}
	b.stringBuilder = stringBuilder{parent: b}
	b.arrayBuilder = arrayBuilder{parent: b}
	b.dateBuilder = dateBuilder{parent: b}

	return b
}

type Builder struct {
	arithmeticBuilder
	comparisonBuilder
	logicalBuilder
	stringBuilder
	arrayBuilder
	dateBuilder

	d bson.D
}

func (b *Builder) Build() bson.D {
	return b.d
}

func (b *Builder) AddKeyValues(keyValues ...any) *Builder {
	if len(keyValues)%2 == 0 {
		for i := 0; i < len(keyValues); i += 2 {
			key, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			b.d = append(b.d, bson.E{Key: key, Value: keyValues[i+1]})
		}
	}
	return b
}

func (b *Builder) Sum(expression any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationSum, Value: expression})
	return b
}

func (b *Builder) SumMany(expressions ...any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationSum, Value: expressions})
	return b
}

func (b *Builder) Contact(expressions ...any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationContact, Value: expressions})
	return b
}

func (b *Builder) Push(expression any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationPush, Value: expression})
	return b
}

func (b *Builder) Avg(expression any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationAvg, Value: expression})
	return b
}

func (b *Builder) AvgMany(expressions ...any) *Builder {
	b.d = append(b.d, bson.E{Key: types.AggregationAvg, Value: expressions})
	return b
}
