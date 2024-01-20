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
	"testing"

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func TestId(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test id",
			value: "123",
			want:  bson.D{bson.E{Key: "_id", Value: "123"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Id(tc.value))
		})
	}
}

func TestAll(t *testing.T) {
	testCases := []struct {
		name   string
		values []int
		want   bson.D
	}{
		{
			name:   "test all",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "$all", Value: []int{1, 2, 3}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, All(tc.values...))
		})
	}
}

func TestElemMatch(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		cond any
		want bson.D
	}{
		{
			name: "test elem match",
			key:  "key",
			cond: bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}},
			want: bson.D{bson.E{Key: "key", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ElemMatch(tc.key, tc.cond))
		})
	}
}

func TestSize(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		size int
		want bson.D
	}{
		{
			name: "test size",
			key:  "apples",
			size: 10,
			want: bson.D{bson.E{Key: "apples", Value: bson.D{{Key: "$size", Value: 10}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Size(tc.key, tc.size))
		})
	}
}

func TestEq(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test eq",
			key:   "name",
			value: "cmy",
			want:  bson.D{bson.E{Key: "name", Value: bson.D{{Key: "$eq", Value: "cmy"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Eq(tc.key, tc.value))
		})
	}
}

func TestGt(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test gt",
			key:   "age",
			value: 18,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Gt(tc.key, tc.value))
		})
	}
}

func TestGte(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test gte",
			key:   "age",
			value: 18,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gte", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Gte(tc.key, tc.value))
		})
	}
}

func TestIn(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name:   "test in",
			key:    "age",
			values: []int{18, 19, 20},
			want:   bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$in", Value: []int{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, In(tc.key, tc.values...))
		})
	}
}

func TestNIn(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name:   "test nin",
			key:    "age",
			values: []int{18, 19, 20},
			want:   bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$nin", Value: []int{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NIn(tc.key, tc.values...))
		})
	}
}

func TestLt(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test lt",
			key:   "age",
			value: 18,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Lt(tc.key, tc.value))
		})
	}
}

func TestLte(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test lte",
			key:   "age",
			value: 18,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lte", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Lte(tc.key, tc.value))
		})
	}
}

func TestNe(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test ne",
			key:   "age",
			value: 18,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$ne", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Ne(tc.key, tc.value))
		})
	}
}

func TestExists(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value bool
		want  bson.D
	}{
		{
			name:  "test exists",
			key:   "age",
			value: true,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$exists", Value: true}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Exists(tc.key, tc.value))
		})
	}
}

func TestType(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value bsontype.Type
		want  bson.D
	}{
		{
			name:  "test type",
			key:   "age",
			value: bson.TypeInt32,
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$type", Value: bson.TypeInt32}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Type(tc.key, tc.value))
		})
	}
}

func TestTypeAlias(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value string
		want  bson.D
	}{
		{
			name:  "test type alias",
			key:   "age",
			value: "int",
			want:  bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$type", Value: "int"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, TypeAlias(tc.key, tc.value))
		})
	}
}

func TestTypeArray(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []bsontype.Type
		want   bson.D
	}{
		{
			name:   "test type array",
			key:    "age",
			values: []bsontype.Type{bson.TypeInt32, bson.TypeInt64},
			want:   bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$type", Value: []bsontype.Type{bson.TypeInt32, bson.TypeInt64}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, TypeArray(tc.key, tc.values...))
		})
	}
}

func TestTypeArrayAlias(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []string
		want   bson.D
	}{
		{
			name:   "test type array alias",
			key:    "age",
			values: []string{"int", "long"},
			want:   bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$type", Value: []string{"int", "long"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, TypeArrayAlias(tc.key, tc.values...))
		})
	}
}

func TestExpr(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test expr",
			value: bson.D{bson.E{Key: "$gt", Value: 18}},
			want:  bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$gt", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Expr(tc.value))
		})
	}
}

func TestJsonSchema(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test json schema",
			value: bson.D{bson.E{Key: "$gt", Value: 18}},
			want:  bson.D{bson.E{Key: "$jsonSchema", Value: bson.D{bson.E{Key: "$gt", Value: 18}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, JsonSchema(tc.value))
		})
	}
}

func TestMod(t *testing.T) {
	testCases := []struct {
		name      string
		key       string
		divisor   any
		remainder int
		want      bson.D
	}{
		{
			name:      "test mod",
			key:       "age",
			divisor:   10,
			remainder: 1,
			want:      bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$mod", Value: bson.A{10, 1}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Mod(tc.key, tc.divisor, tc.remainder))
		})
	}
}

func TestRegex(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value string
		want  bson.D
	}{
		{
			name:  "test regex",
			key:   "name",
			value: ".*cmy.*",
			want:  bson.D{bson.E{Key: "name", Value: bson.D{{Key: "$regex", Value: ".*cmy.*"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Regex(tc.key, tc.value))
		})
	}
}

func TestRegexOptions(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		value   string
		options string
		want    bson.D
	}{
		{
			name:    "test regex options",
			key:     "name",
			value:   ".*cmy.*",
			options: "i",
			want:    bson.D{bson.E{Key: "name", Value: bson.D{{Key: "$regex", Value: ".*cmy.*"}, {Key: "$options", Value: "i"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, RegexOptions(tc.key, tc.value, tc.options))
		})
	}
}

func TestText(t *testing.T) {
	testCases := []struct {
		name   string
		search string
		opt    *types.TextOptions
		want   bson.D
	}{
		{
			name:   "nil opt",
			search: "cmy",
			opt:    nil,
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}}}},
		},
		{
			name:   "empty opt",
			search: "cmy",
			opt:    &types.TextOptions{},
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}}}},
		},
		{
			name:   "nil language",
			search: "cmy",
			opt:    &types.TextOptions{CaseSensitive: true, DiacriticSensitive: true},
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}, {Key: "$caseSensitive", Value: true}, {Key: "$diacriticSensitive", Value: true}}}},
		},
		{
			name:   "nil caseSensitive",
			search: "cmy",
			opt:    &types.TextOptions{Language: "en", DiacriticSensitive: true},
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}, {Key: "$language", Value: "en"}, {Key: "$diacriticSensitive", Value: true}}}},
		},
		{
			name:   "nil diacriticSensitive",
			search: "cmy",
			opt:    &types.TextOptions{Language: "en", CaseSensitive: true},
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}, {Key: "$language", Value: "en"}, {Key: "$caseSensitive", Value: true}}}},
		},
		{
			name:   "all not nil",
			search: "cmy",
			opt:    &types.TextOptions{Language: "en", CaseSensitive: true, DiacriticSensitive: true},
			want:   bson.D{bson.E{Key: "$text", Value: bson.D{bson.E{Key: "$search", Value: "cmy"}, {Key: "$language", Value: "en"}, {Key: "$caseSensitive", Value: true}, {Key: "$diacriticSensitive", Value: true}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Text(tc.search, tc.opt))
		})
	}
}

func TestWhere(t *testing.T) {
	testCases := []struct {
		name  string
		value string
		want  bson.D
	}{
		{
			name:  "test where",
			value: "this.age > 18",
			want:  bson.D{bson.E{Key: "$where", Value: "this.age > 18"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Where(tc.value))
		})
	}
}

func TestAnd(t *testing.T) {
	testCases := []struct {
		name       string
		conditions []any
		want       bson.D
	}{
		{
			name:       "test and",
			conditions: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}},
			want:       bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, And(tc.conditions...))
		})
	}
}

func TestNot(t *testing.T) {
	testCases := []struct {
		name string
		cond bson.D
		want bson.D
	}{
		{
			name: "test not",
			cond: bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}},
			want: bson.D{bson.E{Key: "$not", Value: bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Not(tc.cond))
		})
	}
}

func TestNor(t *testing.T) {
	testCases := []struct {
		name       string
		conditions []any
		want       bson.D
	}{
		{
			name:       "test nor",
			conditions: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}},
			want:       bson.D{bson.E{Key: "$nor", Value: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Nor(tc.conditions...))
		})
	}
}

func TestOr(t *testing.T) {
	testCases := []struct {
		name       string
		conditions []any
		want       bson.D
	}{
		{
			name:       "test or",
			conditions: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}},
			want:       bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$gt", Value: 18}}}}, bson.D{bson.E{Key: "age", Value: bson.D{{Key: "$lt", Value: 30}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Or(tc.conditions...))
		})
	}
}

func TestSlice(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		number int
		want   bson.D
	}{
		{
			name:   "test slice",
			key:    "details.colors",
			number: 1,
			want:   bson.D{bson.E{Key: "details.colors", Value: bson.D{{Key: "$slice", Value: 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, Slice(tc.key, tc.number))
		})
	}
}

func TestSliceRanger(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		start  int
		number int
		want   bson.D
	}{
		{
			name:   "test slice ranger",
			key:    "details.colors",
			start:  1,
			number: 1,
			want:   bson.D{bson.E{Key: "details.colors", Value: bson.D{{Key: "$slice", Value: []int{1, 1}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.want, SliceRanger(tc.key, tc.start, tc.number))
		})
	}
}
