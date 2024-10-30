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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func All[T any](values ...T) bson.D {
	return bson.D{bson.E{Key: AllOp, Value: values}}
}

func And(conditions ...any) bson.D {
	return bson.D{bson.E{Key: AndOp, Value: conditions}}
}

func ElemMatch(key string, cond any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: ElemMatchOp, Value: cond}}}}
}

func Eq(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: EqOp, Value: value}}}}
}

func Exists(key string, value bool) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: ExistsOp, Value: value}}}}
}

func Expr(value any) bson.D {
	return bson.D{bson.E{Key: ExprOp, Value: value}}
}

func Gt(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: GtOp, Value: value}}}}
}

func Gte(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: GteOp, Value: value}}}}
}

func Id(value any) bson.D {
	return bson.D{bson.E{Key: IdOp, Value: value}}
}

func In[T any](key string, values ...T) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: InOp, Value: values}}}}
}

func JsonSchema(value any) bson.D {
	return bson.D{bson.E{Key: JsonSchemaOp, Value: value}}
}

func Lt(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: LtOp, Value: value}}}}
}

func Lte(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: LteOp, Value: value}}}}
}

func Mod(key string, divisor any, remainder int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: ModOp, Value: bson.A{divisor, remainder}}}}}
}

func NIn[T any](key string, values ...T) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: NinOp, Value: values}}}}
}

func Ne(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: NeOp, Value: value}}}}
}

func Nor(conditions ...any) bson.D {
	return bson.D{bson.E{Key: NorOp, Value: conditions}}
}

func Not(cond any) bson.D {
	return bson.D{bson.E{Key: NotOp, Value: cond}}
}

func Or(conditions ...any) bson.D {
	return bson.D{bson.E{Key: OrOp, Value: conditions}}
}

func Regex(key, value string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: RegexOp, Value: value}}}}
}

func RegexOptions(key, value, options string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: RegexOp, Value: value}, {Key: OptionsOp, Value: options}}}}
}

func Size(key string, value int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: SizeOp, Value: value}}}}
}

func Slice(key string, number int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: SliceOp, Value: number}}}}
}

func SliceRanger(key string, start, end int) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: SliceOp, Value: []int{start, end}}}}}
}

func Text(search string, opt *TextOptions) bson.D {
	d := bson.D{bson.E{Key: SearchOp, Value: search}}
	if opt != nil {
		if opt.Language != "" {
			d = append(d, bson.E{Key: LanguageOp, Value: opt.Language})
		}
		if opt.CaseSensitive {
			d = append(d, bson.E{Key: CaseSensitiveOp, Value: opt.CaseSensitive})
		}
		if opt.DiacriticSensitive {
			d = append(d, bson.E{Key: DiacriticSensitiveOp, Value: opt.DiacriticSensitive})
		}
	}

	return bson.D{bson.E{Key: TextOp, Value: d}}
}

func Type(key string, value bson.Type) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: TypeOp, Value: value}}}}
}

func TypeAlias(key string, value string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: TypeOp, Value: value}}}}
}

func TypeArray(key string, values ...bson.Type) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: TypeOp, Value: values}}}}
}

func TypeArrayAlias(key string, values ...string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{{Key: TypeOp, Value: values}}}}
}

func Where(value string) bson.D {
	return bson.D{bson.E{Key: WhereOp, Value: value}}
}
