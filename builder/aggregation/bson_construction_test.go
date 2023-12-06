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
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test sum",
			expression: "$price",
			want:       bson.D{{Key: "$sum", Value: "$price"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSumMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "test sum many",
			expressions: []any{"$price", "$fee"},
			want:        bson.D{{Key: "$sum", Value: []any{"$price", "$fee"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumMany(tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPush(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test push",
			expression: bsonx.NewD().Add("item", "$item").Add("quantity", "$quantity").Build(),
			want:       bson.D{{Key: "$push", Value: bson.D{{Key: "item", Value: "$item"}, {Key: "quantity", Value: "$quantity"}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Push(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAvg(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test avg",
			expression: Multiply("$price", "$quantity"),
			want:       bson.D{{Key: "$avg", Value: bson.D{{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Avg(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAvgMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "test avg many",
			expressions: []any{"$price", "$fee"},
			want:        bson.D{{Key: "$avg", Value: []any{"$price", "$fee"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AvgMany(tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFirst(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test first",
			expression: "$type",
			want:       bson.D{{Key: "$first", Value: "$type"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := First(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLast(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test last",
			expression: "$type",
			want:       bson.D{{Key: "$last", Value: "$type"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Last(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test min",
			expression: "$price",
			want:       bson.D{{Key: "$min", Value: "$price"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Min(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMinMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "test min many",
			expressions: []any{"$price", "$fee"},
			want:        bson.D{{Key: "$min", Value: []any{"$price", "$fee"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MinMany(tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test max",
			expression: "$price",
			want:       bson.D{{Key: "$max", Value: "$price"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Max(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMaxMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "test max many",
			expressions: []any{"$price", "$fee"},
			want:        bson.D{{Key: "$max", Value: []any{"$price", "$fee"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxMany(tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$add", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$add", Value: []any{}}},
		},
		{
			name:        "single type",
			expressions: []any{1, 2, 3, 4},
			want:        bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, 4}}},
		},
		{
			name:        "multiple types",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			want:        bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, "$a", "$b", "$c"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.expressions)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$multiply", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$multiply", Value: []any{}}},
		},
		{
			name:        "single type",
			expressions: []any{1, 2, 3, 4},
			want:        bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, 4}}},
		},
		{
			name:        "multiple types",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			want:        bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, "$a", "$b", "$c"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Multiply(tc.expressions...))
		})
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		start  int64
		length int64
		want   bson.D
	}{
		{
			name:   "normal",
			s:      "$quarter",
			start:  0,
			length: 2,
			want:   bson.D{bson.E{Key: "$subtract", Value: []any{"$quarter", int64(0), int64(2)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Subtract(tc.s, tc.start, tc.length))
		})
	}
}

func Test_Divide(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"hours", 8},
			want:        bson.D{bson.E{Key: "$divide", Value: []any{"hours", 8}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Divide(tc.expressions...))
		})
	}
}

func Test_Mod(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$hours", "$tasks"},
			want:        bson.D{bson.E{Key: "$mod", Value: []any{"$hours", "$tasks"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Mod(tc.expressions...))
		})
	}
}

func Test_ArrayElemAt(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		index      int64
		want       bson.D
	}{
		{
			name:       "valid expression",
			expression: "$favorites",
			index:      0,
			want:       bson.D{{Key: "$arrayElemAt", Value: []any{"$favorites", int64(0)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ArrayElemAt(tc.expression, tc.index))
		})
	}
}

func Test_ConcatArrays(t *testing.T) {
	testCases := []struct {
		name   string
		arrays []any
		want   bson.D
	}{
		{
			name:   "valid arrays",
			arrays: []any{"$instock", "$ordered"},
			want:   bson.D{{Key: "$concatArrays", Value: []any{"$instock", "$ordered"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ConcatArrays(tc.arrays...))
		})
	}
}

func Test_ArrayToObject(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "string expression",
			expression: "$dimensions",
			want:       bson.D{{Key: "$arrayToObject", Value: "$dimensions"}},
		},
		{
			name:       "array expression",
			expression: []any{bsonx.NewD().Add("k", "item").Add("v", "abc123").Build()},
			want:       bson.D{{Key: "$arrayToObject", Value: []any{bson.D{{Key: "k", Value: "item"}, {Key: "v", Value: "abc123"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ArrayToObject(tc.expression))
		})
	}
}

func Test_Size(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "valid expression",
			expression: "$items",
			want:       bson.D{{Key: "$size", Value: "$items"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Size(tc.expression))
		})
	}
}

func Test_Slice(t *testing.T) {
	testCases := []struct {
		name      string
		array     any
		nElements int64
		want      bson.D
	}{
		{
			name:      "valid expression",
			array:     "$items",
			nElements: 5,
			want:      bson.D{{Key: "$slice", Value: []any{"$items", int64(5)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Slice(tc.array, tc.nElements))
		})
	}
}

func Test_SliceWithPosition(t *testing.T) {
	testCases := []struct {
		name      string
		array     any
		position  int64
		nElements int64
		want      bson.D
	}{
		{
			name:      "valid expression",
			array:     "$items",
			position:  20,
			nElements: 5,
			want:      bson.D{{Key: "$slice", Value: []any{"$items", int64(20), int64(5)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, SliceWithPosition(tc.array, tc.position, tc.nElements))
		})
	}
}

func Test_Map(t *testing.T) {
	testCases := []struct {
		name       string
		inputArray any
		as         string
		in         any
		want       bson.D
	}{
		{
			name:       "valid expression",
			inputArray: "$items",
			as:         "item",
			in:         "$$item.price * 1.25",
			want:       bson.D{{Key: "$map", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "as", Value: "item"}, {Key: "in", Value: "$$item.price * 1.25"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Map(tc.inputArray, tc.as, tc.in))
		})
	}
}

func Test_Filter(t *testing.T) {
	testCases := []struct {
		name       string
		inputArray any
		cond       any
		opt        *types.FilterOptions
		want       bson.D
	}{
		{
			name:       "valid expression",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        nil,
			want:       bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}}}},
		},
		{
			name:       "valid expression with options",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        &types.FilterOptions{As: "item", Limit: 5},
			want:       bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}, {Key: "as", Value: "item"}, {Key: "limit", Value: int64(5)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Filter(tc.inputArray, tc.cond, tc.opt))
		})
	}
}

func Test_Eq(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$eq", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$eq", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$eq", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Eq(tc.expressions...))
		})
	}
}

func Test_Ne(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$ne", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$ne", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$ne", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Ne(tc.expressions...))
		})
	}
}

func Test_Gt(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$gt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$gt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Gt(tc.expressions...))
		})
	}
}

func Test_Gte(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$gte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$gte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$gte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Gte(tc.expressions...))
		})
	}
}

func Test_Lt(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$lt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$lt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Lt(tc.expressions...))
		})
	}
}

func Test_Lte(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$lte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$lte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			want:        bson.D{bson.E{Key: "$lte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Lte(tc.expressions...))
		})
	}
}

func Test_Cond(t *testing.T) {
	testCases := []struct {
		name      string
		boolExpr  any
		trueExpr  any
		falseExpr any
		want      bson.D
	}{
		{
			name:      "normal",
			boolExpr:  Gte("$qty", 250),
			trueExpr:  30,
			falseExpr: 20,
			want:      bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Cond(tc.boolExpr, tc.trueExpr, tc.falseExpr))
		})
	}
}

func Test_IfNull(t *testing.T) {
	testCases := []struct {
		name        string
		expr        any
		replacement any
		want        bson.D
	}{
		{
			name:        "nil expr",
			expr:        nil,
			replacement: "Unspecified",
			want:        bson.D{{Key: "$ifNull", Value: []any{nil, "Unspecified"}}},
		},
		{
			name:        "nil replacement",
			expr:        "$description",
			replacement: nil,
			want:        bson.D{{Key: "$ifNull", Value: []any{"$description", nil}}},
		},
		{
			name:        "normal",
			expr:        "$description",
			replacement: "Unspecified",
			want:        bson.D{{Key: "$ifNull", Value: []any{"$description", "Unspecified"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IfNull(tc.expr, tc.replacement))
		})
	}
}

func Test_Switch(t *testing.T) {
	testCases := []struct {
		name        string
		cases       []any
		defaultCase any
		want        bson.D
	}{
		{
			name:        "nil cases",
			cases:       nil,
			defaultCase: "Did not match",
			want:        bson.D{},
		},
		{
			name:        "empty cases",
			cases:       []any{},
			defaultCase: "Did not match",
			want:        bson.D{},
		},
		{
			name: "normal",
			cases: []any{
				Eq(0, 5), "equals",
				Gt(0, 5), "greater than",
			},
			defaultCase: "Did not match",
			want: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{
						bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
						bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
					}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Switch(tc.cases, tc.defaultCase))
		})
	}
}

func Test_DateOfMonth(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$dayOfMonth", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfMonth(tc.date))
		})
	}
}

func Test_DayOfMonthWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$dayOfMonth", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfMonthWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_DayOfWeek(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$dayOfWeek", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfWeek(tc.date))
		})
	}
}

func Test_DayOfWeekWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$dayOfWeek", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfWeekWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_DayOfYear(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$dayOfYear", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfYear(tc.date))
		})
	}
}

func Test_DayOfYearWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$dayOfYear", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DayOfYearWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_Year(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$year", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Year(tc.date))
		})
	}
}

func Test_YearWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$year", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, YearWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_Month(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$month", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Month(tc.date))
		})
	}
}

func Test_MonthWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$month", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, MonthWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_Week(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		want bson.D
	}{
		{
			name: "normal date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			want: bson.D{bson.E{Key: "$week", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Week(tc.date))
		})
	}
}

func Test_WeekWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		want     bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			want:     bson.D{bson.E{Key: "$week", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, WeekWithTimezone(tc.date, tc.timezone))
		})
	}
}

func Test_DateToString(t *testing.T) {
	testCases := []struct {
		name string
		date time.Time
		opt  *types.DateToStringOptions
		want bson.D
	}{
		{
			name: "nil opt",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt:  nil,
			want: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}},
		},
		{
			name: "empty format",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "",
				Timezone: "Asia/Shanghai",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
				bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
		{
			name: "empty timezone",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
				bson.E{Key: "format", Value: "%Y-%m-%d"}}}},
		},
		{
			name: "nil onNull",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "Asia/Shanghai",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
		{
			name: "normal",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "Asia/Shanghai",
				OnNull:   "null",
			},
			want: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}, bson.E{Key: "onNull", Value: "null"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DateToString(tc.date, tc.opt))
		})
	}
}

func Test_And(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$and", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$and", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{Gt("$qty", 100), Lt("$qty", 250)},
			want:        bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, And(tc.expressions...))
		})
	}
}

func Test_Not(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$not", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$not", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{Gt("$qty", 250)},
			want:        bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Not(tc.expressions...))
		})
	}
}

func Test_Or(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "$or", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "$or", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{query.Eq("x", 0), query.Expr(Eq(Divide(1, "$x"), 3))},
			want:        bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "x", Value: bson.D{bson.E{Key: "$eq", Value: 0}}}}, bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$eq", Value: []any{bson.D{bson.E{Key: "$divide", Value: []any{1, "$x"}}}, 3}}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Or(tc.expressions...))
		})
	}
}

func Test_Concat(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			want:        bson.D{{Key: "$concat", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			want:        bson.D{{Key: "$concat", Value: []any{}}},
		},
		{
			name:        "normal expression",
			expressions: []any{"$item", " - ", "$description"},
			want:        bson.D{{Key: "$concat", Value: []any{"$item", " - ", "$description"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Concat(tc.expressions...))
		})
	}
}

func Test_SubstrBytes(t *testing.T) {
	testCases := []struct {
		name             string
		stringExpression string
		byteIndex        int64
		byteCount        int64
		want             bson.D
	}{
		{
			name:             "normal expression",
			stringExpression: "$quarter",
			byteIndex:        0,
			byteCount:        2,
			want:             bson.D{{Key: "$substrBytes", Value: []any{"$quarter", int64(0), int64(2)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, SubstrBytes(tc.stringExpression, tc.byteIndex, tc.byteCount))
		})
	}
}

func Test_ToLower(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "normal expression",
			expression: "$item",
			want:       bson.D{{Key: "$toLower", Value: "$item"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ToLower(tc.expression))
		})
	}
}

func Test_ToUpper(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "normal expression",
			expression: "$item",
			want:       bson.D{{Key: "$toUpper", Value: "$item"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ToUpper(tc.expression))
		})
	}
}

func Test_Contact(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		want        bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$item", " - ", "$description"},
			want:        bson.D{{Key: "$concat", Value: []any{"$item", " - ", "$description"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Contact(tc.expressions...))
		})
	}
}
