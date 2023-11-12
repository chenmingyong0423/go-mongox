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

type logicalQueryBuilder struct {
	parent *Builder
}

// And
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) And(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.And, Value: conditions})
	return b.parent
}

func (b *logicalQueryBuilder) Not(condition any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Not, Value: condition})
	return b.parent
}

// Nor
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) Nor(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Nor, Value: conditions})
	return b.parent
}

// Or
// 对于 conditions 参数，你同样可以使用 QueryBuilder 去生成
func (b *logicalQueryBuilder) Or(conditions ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Or, Value: conditions})
	return b.parent
}
