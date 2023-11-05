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

type accumulatorsBuilder struct {
	parent *Builder
}

func (b *accumulatorsBuilder) Sum(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSum, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) SumMany(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSum, Value: expressions})
	return b.parent
}

func (b *accumulatorsBuilder) Push(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationPush, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Avg(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationAvg, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) AvgMany(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationAvg, Value: expressions})
	return b.parent
}

func (b *accumulatorsBuilder) First(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationFirst, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Last(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationLast, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) Min(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMin, Value: expression})
	return b.parent
}
func (b *accumulatorsBuilder) MinMany(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMin, Value: expressions})
	return b.parent
}

func (b *accumulatorsBuilder) Max(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMax, Value: expression})
	return b.parent
}

func (b *accumulatorsBuilder) MaxMany(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMax, Value: expressions})
	return b.parent
}
