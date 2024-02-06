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
	"github.com/chenmingyong0423/go-mongox/pkg/utils"
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type evaluationQueryBuilder struct {
	parent *Builder
}

func (b *evaluationQueryBuilder) Expr(d bson.D) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Expr, Value: d})
	return b.parent
}

func (b *evaluationQueryBuilder) JsonSchema(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.JsonSchema, Value: value})
	return b.parent
}

func (b *evaluationQueryBuilder) Mod(key string, divisor any, remainder int) *Builder {
	if utils.IsNumeric(divisor) {
		e := bson.E{Key: types.Mod, Value: bson.A{divisor, remainder}}
		if !b.parent.tryMergeValue(key, e) {
			b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
		}
	}
	return b.parent
}

func (b *evaluationQueryBuilder) Regex(key, value string) *Builder {
	e := bson.E{Key: types.Regex, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *evaluationQueryBuilder) RegexOptions(key, value, options string) *Builder {
	if !b.parent.tryMergeValue(key, bson.E{Key: types.Regex, Value: value}, bson.E{Key: types.Options, Value: options}) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Regex, Value: value}, bson.E{Key: types.Options, Value: options}}})
	}
	return b.parent
}

// Text
// 如果 language 的值为零值，则不作为查询条件 If the value of language is zero, it is not used as a query condition
// 如果 caseSensitive 的值为零值，则不作为查询条件 If the value of caseSensitive is zero, it is not used as a query condition
// 如果 diacriticSensitive 的值为零值，则不作为查询条件 If the value of diacriticSensitive is zero, it is not used as a query condition
func (b *evaluationQueryBuilder) Text(value, language string, caseSensitive, diacriticSensitive bool) *Builder {
	d := bson.D{bson.E{Key: types.Search, Value: value}}
	if language != "" {
		d = append(d, bson.E{Key: types.Language, Value: language})
	}
	if caseSensitive {
		d = append(d, bson.E{Key: types.CaseSensitive, Value: caseSensitive})
	}
	if diacriticSensitive {
		d = append(d, bson.E{Key: types.DiacriticSensitive, Value: diacriticSensitive})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Text, Value: d})
	return b.parent
}

func (b *evaluationQueryBuilder) Where(value string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Where, Value: value})
	return b.parent
}
