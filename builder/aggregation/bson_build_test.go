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
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSum(t *testing.T) {
	t.Run("test sum", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "totalAmount", Value: bson.D{bson.E{Key: "$sum", Value: "$qty"}}}}, Sum("totalAmount", "$qty"))
	})
}

func TestPush(t *testing.T) {
	t.Run("test push", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "allItems", Value: bson.D{bson.E{Key: "$push", Value: "$item"}}}}, Push("allItems", "$item"))
	})
}

func TestAvg(t *testing.T) {
	t.Run("test avg", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "avgAmount", Value: bson.D{bson.E{Key: "$avg", Value: "$qty"}}}}, Avg("avgAmount", "$qty"))
	})
}

func TestFirst(t *testing.T) {
	t.Run("test first", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "firstSale", Value: bson.D{bson.E{Key: "$first", Value: "$date"}}}}, First("firstSale", "$date"))
	})
}

func TestLast(t *testing.T) {
	t.Run("test last", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "lastSale", Value: bson.D{bson.E{Key: "$last", Value: "$date"}}}}, Last("lastSale", "$date"))
	})
}

func TestMin(t *testing.T) {
	t.Run("test min", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "minQuantity", Value: bson.D{bson.E{Key: "$min", Value: "$quantity"}}}}, Min("minQuantity", "$quantity"))
	})
}

func TestMax(t *testing.T) {
	t.Run("test max", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "maxQuantity", Value: bson.D{bson.E{Key: "$max", Value: "$quantity"}}}}, Max("maxQuantity", "$quantity"))
	})
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		key         string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil value",
			key:         "total",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{nil}}}}},
		},
		{
			name:        "empty",
			key:         "total",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{}}}}},
		},
		{
			name:        "single type",
			key:         "total",
			expressions: []any{1, 2, 3, 4},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, 4}}}}},
		},
		{
			name:        "multiple types",
			key:         "total",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.key, tc.expressions...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name        string
		key         string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil",
			key:         "total",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{nil}}}}},
		},
		{
			name:        "empty",
			key:         "total",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{}}}}},
		},
		{
			name:        "single type",
			key:         "total",
			expressions: []any{1, 2, 3, 4},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, 4}}}}},
		},
		{
			name:        "multiple types",
			key:         "total",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			want:        bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Multiply(tc.key, tc.expressions...))
		})
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		name        string
		key         string
		expressions []any
		want        bson.D
	}{
		{
			name:        "nil value",
			key:         "dateDifference",
			expressions: []any{nil},
			want:        bson.D{bson.E{Key: "dateDifference", Value: bson.D{bson.E{Key: "$subtract", Value: []any{nil}}}}},
		},
		{
			name:        "empty",
			key:         "dateDifference",
			expressions: []any{},
			want:        bson.D{bson.E{Key: "dateDifference", Value: bson.D{bson.E{Key: "$subtract", Value: []any{}}}}},
		},
		{
			name:        "single type",
			key:         "dateDifference",
			expressions: []any{1, 2, 3, 4},
			want:        bson.D{bson.E{Key: "dateDifference", Value: bson.D{bson.E{Key: "$subtract", Value: []any{1, 2, 3, 4}}}}},
		},
		{
			name:        "multiple types",
			key:         "dateDifference",
			expressions: []any{"$date", 5 * 60 * 1000},
			want:        bson.D{bson.E{Key: "dateDifference", Value: bson.D{bson.E{Key: "$subtract", Value: []any{"$date", 5 * 60 * 1000}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Subtract(tc.key, tc.expressions...))
		})
	}
}

func Test_Divide(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "workdays", Value: bson.D{bson.E{Key: "$divide", Value: []any{"$hours", 8}}}}}, Divide("workdays", "$hours", 8))
}

func Test_Mod(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "remainder", Value: bson.D{bson.E{Key: "$mod", Value: []any{"$distance", 5}}}}}, Mod("remainder", "$distance", 5))
}

func Test_ArrayElemAt(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "first", Value: bson.D{bson.E{Key: "$arrayElemAt", Value: []any{"$favorites", int64(0)}}}}}, ArrayElemAt("first", "$favorites", int64(0)))
}

func Test_ConcatArrays(t *testing.T) {
	// { items: { $concatArrays: [ "$instock", "$ordered" ] }
	assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{bson.E{Key: "$concatArrays", Value: []any{"$instock", "$ordered"}}}}}, ConcatArrays("items", "$instock", "$ordered"))
}

func Test_ArrayToObject(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$arrayToObject", Value: "$item"}}}}, ArrayToObject("item", "$item"))
}

func Test_Size(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "numOfScores", Value: bson.D{bson.E{Key: "$size", Value: "$scores"}}}}, Size("numOfScores", "$scores"))
}

func Test_Slice(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "firstThree", Value: bson.D{bson.E{Key: "$slice", Value: []any{"$array", int64(3)}}}}}, Slice("firstThree", "$array", int64(3)))
}

func Test_SliceWithPosition(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "firstThree", Value: bson.D{bson.E{Key: "$slice", Value: []any{"$array", int64(1), int64(3)}}}}}, SliceWithPosition("firstThree", "$array", int64(1), int64(3)))
}

func Test_Map(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "result", Value: bson.D{bson.E{Key: "$map", Value: bson.D{bson.E{Key: "input", Value: "$quizzes"}, {Key: "as", Value: "grade"}, {Key: "in", Value: bson.D{bson.E{Key: "adjustedGrade", Value: bson.D{bson.E{Key: "$add", Value: []any{"$$grade", 2}}}}}}}}}}}, Map("result", "$quizzes", "grade", bson.D{bson.E{Key: "adjustedGrade", Value: bson.D{bson.E{Key: "$add", Value: []any{"$$grade", 2}}}}}))
}

func Test_Filter(t *testing.T) {
	testCases := []struct {
		name       string
		key        string
		inputArray any
		cond       any
		opt        *types.FilterOptions
		want       bson.D
	}{
		{
			name:       "nil options",
			key:        "items",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        nil,
			want:       bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}}}}}},
		},
		{
			name:       "with options",
			key:        "items",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        &types.FilterOptions{As: "item", Limit: int64(5)},
			want:       bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}, {Key: "as", Value: "item"}, {Key: "limit", Value: int64(5)}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Filter(tc.key, tc.inputArray, tc.cond, tc.opt))
		})
	}
}

func Test_Eq(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$eq", Value: []any{"$item", 1}}}}}, Eq("item", "$item", 1))
}

func Test_Ne(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$ne", Value: []any{"$item", 1}}}}}, Ne("item", "$item", 1))
}

func Test_Gt(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$gt", Value: []any{"$item", 1}}}}}, Gt("item", "$item", 1))
}

func Test_Gte(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$gte", Value: []any{"$item", 1}}}}}, Gte("item", "$item", 1))
}

func Test_Lt(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$lt", Value: []any{"$item", 1}}}}}, Lt("item", "$item", 1))
}

func Test_Lte(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$lte", Value: []any{"$item", 1}}}}}, Lte("item", "$item", 1))
}

func Test_Cond(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}}}}, Cond("discount", bsonx.D("$gte", []any{"$qty", 250}), 30, 20))
}

func Test_IfNull(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$ifNull", Value: []any{"$coupon", int64(0)}}}}}, IfNull("discount", "$coupon", int64(0)))

}

func Test_Switch(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "summary", Value: bson.D{
		{Key: "$switch", Value: bson.D{
			{Key: "branches", Value: bson.A{
				bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
				bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
			}},
			{Key: "default", Value: "Did not match"},
		}},
	}}},
		Switch("summary", []types.CaseThen{
			{
				Case: bsonx.D("$eq", []any{0, 5}), Then: "equals",
			},
			{
				Case: bsonx.D("$gt", []any{0, 5}), Then: "greater than",
			},
		}, "Did not match"),
	)
}

func Test_DayOfMonth(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfMonth", Value: bson.D{bson.E{Key: "$dayOfMonth", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, DayOfMonth("dayOfMonth", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_DayOfMonthWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfMonth", Value: bson.D{bson.E{Key: "$dayOfMonth", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, DayOfMonthWithTimezone("dayOfMonth", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_DayOfWeek(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfWeek", Value: bson.D{bson.E{Key: "$dayOfWeek", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, DayOfWeek("dayOfWeek", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_DayOfWeekWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfWeek", Value: bson.D{bson.E{Key: "$dayOfWeek", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, DayOfWeekWithTimezone("dayOfWeek", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_DayOfYear(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfYear", Value: bson.D{bson.E{Key: "$dayOfYear", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, DayOfYear("dayOfYear", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_DayOfYearWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "dayOfYear", Value: bson.D{bson.E{Key: "$dayOfYear", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, DayOfYearWithTimezone("dayOfYear", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_Year(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "year", Value: bson.D{bson.E{Key: "$year", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, Year("year", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_YearWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "year", Value: bson.D{bson.E{Key: "$year", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, YearWithTimezone("year", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_Month(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "month", Value: bson.D{bson.E{Key: "$month", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, Month("month", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_MonthWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "month", Value: bson.D{bson.E{Key: "$month", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, MonthWithTimezone("month", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_Week(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "week", Value: bson.D{bson.E{Key: "$week", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, Week("week", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)))
}

func Test_WeekWithTimezone(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "week", Value: bson.D{bson.E{Key: "$week", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}}, WeekWithTimezone("week", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "Asia/Shanghai"))
}

func Test_DateToString(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		date time.Time
		opt  *types.DateToStringOptions
		want bson.D
	}{
		{
			name: "nil opt",
			key:  "date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt:  nil,
			want: bson.D{bson.E{Key: "date", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}}},
		},
		{
			name: "empty format",
			key:  "date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "",
				Timezone: "Asia/Shanghai",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "date", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
				bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}},
		},
		{
			name: "empty timezone",
			key:  "date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "date", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
				bson.E{Key: "format", Value: "%Y-%m-%d"}}}}}},
		},
		{
			name: "nil onNull",
			key:  "date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "Asia/Shanghai",
				OnNull:   nil,
			},
			want: bson.D{bson.E{Key: "date", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}}}},
		},
		{
			name: "normal",
			key:  "date",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "Asia/Shanghai",
				OnNull:   "null",
			},
			want: bson.D{bson.E{Key: "date", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}, bson.E{Key: "onNull", Value: "null"}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, DateToString(tc.key, tc.date, tc.opt))
		})
	}
}

func Test_And(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}}}}, And("item", bsonx.D("$gt", []any{"$qty", 100}), bsonx.D("$lt", []any{"$qty", 250})))
}

func Test_Not(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}}}}, Not("item", bsonx.D("$gt", []any{"$qty", 250})))
}

func Test_Or(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 50}}}}}}}}, Or("item", bsonx.D("$gt", []any{"$qty", 250}), bsonx.D("$lt", []any{"$qty", 50})))
}

func Test_Concat(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$item", " - ", "$description"}}}}}, Concat("item", "$item", " - ", "$description"))
}

func Test_SubstrBytes(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "quarterSubtring", Value: bson.D{{Key: "$substrBytes", Value: []any{"$quarter", int64(0), int64(2)}}}}}, SubstrBytes("quarterSubtring", "$quarter", int64(0), int64(2)))
}

func Test_ToLower(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{{Key: "$toLower", Value: "$item"}}}}, ToLower("item", "$item"))
}

func Test_ToUpper(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{{Key: "$toUpper", Value: "$item"}}}}, ToUpper("item", "$item"))
}

func Test_Contact(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$item", " - ", "$description"}}}}}, Contact("item", "$item", " - ", "$description"))
}
