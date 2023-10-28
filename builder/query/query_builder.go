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
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

func Query() *QueryBuilder {
	query := &QueryBuilder{
		data: bson.D{},
	}
	query.comparisonQueryBuilder = comparisonQueryBuilder{parent: query}
	query.logicalQueryBuilder = logicalQueryBuilder{parent: query}
	query.elementQueryBuilder = elementQueryBuilder{parent: query}
	query.arrayQueryBuilder = arrayQueryBuilder{parent: query}
	return query
}

type QueryBuilder struct {
	data bson.D
	comparisonQueryBuilder
	logicalQueryBuilder
	elementQueryBuilder
	arrayQueryBuilder
}

func (b *QueryBuilder) Build() bson.D {
	return b.data
}

func (b *QueryBuilder) Id(v any) *QueryBuilder {
	b.data = append(b.data, bson.E{Key: types.Id, Value: v})
	return b
}

func (b *QueryBuilder) Add(k string, v any) *QueryBuilder {
	b.data = append(b.data, bson.E{Key: k, Value: v})
	return b
}
