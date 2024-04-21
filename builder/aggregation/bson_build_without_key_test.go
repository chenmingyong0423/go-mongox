// Copyright 2024 chenmingyong0423

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

func TestSumWithoutKey(t *testing.T) {
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
			got := SumWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPushWithoutKey(t *testing.T) {
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
			got := PushWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAvgWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       bson.D
	}{
		{
			name:       "test avg",
			expression: MultiplyWithoutKey("$price", "$quantity"),
			want:       bson.D{{Key: "$avg", Value: bson.D{{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AvgWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFirstWithoutKey(t *testing.T) {
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
			got := FirstWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLastWithoutKey(t *testing.T) {
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
			got := LastWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMinWithoutKey(t *testing.T) {
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
			got := MinWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMaxWithoutKey(t *testing.T) {
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
			got := MaxWithoutKey(tc.expression)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAddWithoutKey(t *testing.T) {
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
			got := AddWithoutKey(tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMultiplyWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, MultiplyWithoutKey(tc.expressions...))
		})
	}
}

func TestSubtractWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$date", 5 * 60 * 1000},
			expected:    bson.D{bson.E{Key: "$subtract", Value: []any{"$date", 5 * 60 * 1000}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, SubtractWithoutKey(tc.expressions...))
		})
	}
}

func Test_DivideWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DivideWithoutKey(tc.expressions...))
		})
	}
}

func Test_ModWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ModWithoutKey(tc.expressions...))
		})
	}
}

func Test_ArrayElemAtWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ArrayElemAtWithoutKey(tc.expression, tc.index))
		})
	}
}

func Test_ConcatArraysWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ConcatArraysWithoutKey(tc.arrays...))
		})
	}
}

func Test_ArrayToObjectWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ArrayToObjectWithoutKey(tc.expression))
		})
	}
}

func Test_SizeWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, SizeWithoutKey(tc.expression))
		})
	}
}

func Test_SliceWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, SliceWithoutKey(tc.array, tc.nElements))
		})
	}
}

func Test_SliceWithPositionWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, SliceWithPositionWithoutKey(tc.array, tc.position, tc.nElements))
		})
	}
}

func Test_MapWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, MapWithoutKey(tc.inputArray, tc.as, tc.in))
		})
	}
}

func Test_FilterWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, FilterWithoutKey(tc.inputArray, tc.cond, tc.opt))
		})
	}
}

func Test_EqWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, EqWithoutKey(tc.expressions...))
		})
	}
}

func Test_NeWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, NeWithoutKey(tc.expressions...))
		})
	}
}

func Test_GtWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, GtWithoutKey(tc.expressions...))
		})
	}
}

func Test_GteWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, GteWithoutKey(tc.expressions...))
		})
	}
}

func Test_LtWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, LtWithoutKey(tc.expressions...))
		})
	}
}

func Test_LteWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, LteWithoutKey(tc.expressions...))
		})
	}
}

func Test_CondWithoutKey(t *testing.T) {
	testCases := []struct {
		name      string
		boolExpr  any
		trueExpr  any
		falseExpr any
		want      bson.D
	}{
		{
			name:      "normal",
			boolExpr:  GteWithoutKey("$qty", 250),
			trueExpr:  30,
			falseExpr: 20,
			want:      bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, CondWithoutKey(tc.boolExpr, tc.trueExpr, tc.falseExpr))
		})
	}
}

func Test_IfNullWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, IfNullWithoutKey(tc.expr, tc.replacement))
		})
	}
}

func Test_SwitchWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		cases       []types.CaseThen
		defaultCase any
		want        bson.D
	}{
		{
			name:        "nil cases",
			cases:       nil,
			defaultCase: "Did not match",
			want: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
		{
			name:        "empty cases",
			cases:       []types.CaseThen{},
			defaultCase: "Did not match",
			want: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
		{
			name: "normal",
			cases: []types.CaseThen{
				{
					Case: EqWithoutKey(0, 5),
					Then: "equals",
				},
				{
					Case: GtWithoutKey(0, 5),
					Then: "greater than",
				},
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
			assert.Equal(t, tc.want, SwitchWithoutKey(tc.cases, tc.defaultCase))
		})
	}
}

func Test_DateOfMonthWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfMonthWithoutKey(tc.date))
		})
	}
}

func Test_DayOfMonthWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfMonthWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_DayOfWeekWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfWeekWithoutKey(tc.date))
		})
	}
}

func Test_DayOfWeekWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfWeekWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_DayOfYearWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfYearWithoutKey(tc.date))
		})
	}
}

func Test_DayOfYearWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DayOfYearWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_YearWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, YearWithoutKey(tc.date))
		})
	}
}

func Test_YearWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, YearWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_MonthWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, MonthWithoutKey(tc.date))
		})
	}
}

func Test_MonthWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, MonthWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_WeekWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, WeekWithoutKey(tc.date))
		})
	}
}

func Test_WeekWithTimezoneWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, WeekWithTimezoneWithoutKey(tc.date, tc.timezone))
		})
	}
}

func Test_DateToStringWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, DateToStringWithoutKey(tc.date, tc.opt))
		})
	}
}

func Test_AndWithoutKey(t *testing.T) {
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
			expressions: []any{GtWithoutKey("$qty", 100), LtWithoutKey("$qty", 250)},
			want:        bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, AndWithoutKey(tc.expressions...))
		})
	}
}

func Test_NotWithoutKey(t *testing.T) {
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
			expressions: []any{GtWithoutKey("$qty", 250)},
			want:        bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NotWithoutKey(tc.expressions...))
		})
	}
}

func Test_OrWithoutKey(t *testing.T) {
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
			expressions: []any{query.Eq("x", 0), query.Expr(EqWithoutKey(DivideWithoutKey(1, "$x"), 3))},
			want:        bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "x", Value: bson.D{bson.E{Key: "$eq", Value: 0}}}}, bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$eq", Value: []any{bson.D{bson.E{Key: "$divide", Value: []any{1, "$x"}}}, 3}}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, OrWithoutKey(tc.expressions...))
		})
	}
}

func Test_ConcatWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ConcatWithoutKey(tc.expressions...))
		})
	}
}

func Test_SubstrBytesWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, SubstrBytesWithoutKey(tc.stringExpression, tc.byteIndex, tc.byteCount))
		})
	}
}

func Test_ToLowerWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ToLowerWithoutKey(tc.expression))
		})
	}
}

func Test_ToUpperWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ToUpperWithoutKey(tc.expression))
		})
	}
}

func Test_ContactWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.want, ContactWithoutKey(tc.expressions...))
		})
	}
}
