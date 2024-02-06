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
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func Id(value any) bson.D {
	return bson.D{bson.E{Key: types.Id, Value: value}}
}

func All[T any](values ...T) bson.D {
	return bson.D{bson.E{Key: types.All, Value: values}}
}

func ElemMatch(key string, cond any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.ElemMatch, Value: cond}}}}
}

func Size(key string, value int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Size, Value: value}}}}
}

func Eq(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Eq, Value: value}}}}
}

func Gt(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Gt, Value: value}}}}
}

func Gte(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Gte, Value: value}}}}
}

func In[T any](key string, values ...T) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.In, Value: values}}}}
}

func NIn[T any](key string, values ...T) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Nin, Value: values}}}}
}

func Lt(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Lt, Value: value}}}}
}

func Lte(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Lte, Value: value}}}}
}

func Ne(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Ne, Value: value}}}}
}

func Exists(key string, value bool) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Exists, Value: value}}}}
}

func Type(key string, value bsontype.Type) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Type, Value: value}}}}
}

func TypeAlias(key string, value string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Type, Value: value}}}}
}

func TypeArray(key string, values ...bsontype.Type) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Type, Value: values}}}}
}

func TypeArrayAlias(key string, values ...string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Type, Value: values}}}}
}

func Expr(value any) bson.D {
	return bson.D{bson.E{Key: types.Expr, Value: value}}
}

func JsonSchema(value any) bson.D {
	return bson.D{bson.E{Key: types.JsonSchema, Value: value}}
}

func Mod(key string, divisor any, remainder int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Mod, Value: bson.A{divisor, remainder}}}}}
}

func Regex(key, value string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Regex, Value: value}}}}
}

func RegexOptions(key, value, options string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Regex, Value: value}, {Key: types.Options, Value: options}}}}
}

func Text(search string, opt *types.TextOptions) bson.D {
	d := bson.D{bson.E{Key: types.Search, Value: search}}
	if opt != nil {
		if opt.Language != "" {
			d = append(d, bson.E{Key: types.Language, Value: opt.Language})
		}
		if opt.CaseSensitive {
			d = append(d, bson.E{Key: types.CaseSensitive, Value: opt.CaseSensitive})
		}
		if opt.DiacriticSensitive {
			d = append(d, bson.E{Key: types.DiacriticSensitive, Value: opt.DiacriticSensitive})
		}
	}

	return bson.D{bson.E{Key: types.Text, Value: d}}
}

func Where(value string) bson.D {
	return bson.D{bson.E{Key: types.Where, Value: value}}
}

func And(conditions ...any) bson.D {
	return bson.D{bson.E{Key: types.And, Value: conditions}}
}

func Not(cond any) bson.D {
	return bson.D{bson.E{Key: types.Not, Value: cond}}
}

func Nor(conditions ...any) bson.D {
	return bson.D{bson.E{Key: types.Nor, Value: conditions}}
}

func Or(conditions ...any) bson.D {
	return bson.D{bson.E{Key: types.Or, Value: conditions}}
}

func Slice(key string, number int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Slice, Value: number}}}}
}

func SliceRanger(key string, start, end int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: types.Slice, Value: []int{start, end}}}}}
}
