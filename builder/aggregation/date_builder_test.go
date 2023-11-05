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

	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_dateBuilder_DateOfMonth(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$dayOfMonth", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfMonth(tc.date).Build())
		})
	}
}

func Test_dateBuilder_DayOfMonthWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$dayOfMonth", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfMonthWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_DayOfWeek(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$dayOfWeek", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfWeek(tc.date).Build())
		})
	}
}

func Test_dateBuilder_DayOfWeekWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$dayOfWeek", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfWeekWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_DayOfYear(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$dayOfYear", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfYear(tc.date).Build())
		})
	}
}

func Test_dateBuilder_DayOfYearWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$dayOfYear", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DayOfYearWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_Year(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$year", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Year(tc.date).Build())
		})
	}
}

func Test_dateBuilder_YearWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$year", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().YearWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_Month(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$month", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Month(tc.date).Build())
		})
	}
}

func Test_dateBuilder_MonthWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$month", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().MonthWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_Week(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			expected: bson.D{bson.E{Key: "$week", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Week(tc.date).Build())
		})
	}
}

func Test_dateBuilder_WeekWithTimezone(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		timezone string
		expected bson.D
	}{
		{
			name:     "normal date",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			timezone: "Asia/Shanghai",
			expected: bson.D{bson.E{Key: "$week", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().WeekWithTimezone(tc.date, tc.timezone).Build())
		})
	}
}

func Test_dateBuilder_DateToString(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		opt      *types.DateToStringOptions
		expected bson.D
	}{
		{
			name:     "nil opt",
			date:     time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt:      nil,
			expected: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}},
		},
		{
			name: "empty format",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "",
				Timezone: "Asia/Shanghai",
				OnNull:   nil,
			},
			expected: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
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
			expected: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)},
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
			expected: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}}}},
		},
		{
			name: "normal",
			date: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC),
			opt: &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "Asia/Shanghai",
				OnNull:   "null",
			},
			expected: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "format", Value: "%Y-%m-%d"}, bson.E{Key: "timezone", Value: "Asia/Shanghai"}, bson.E{Key: "onNull", Value: "null"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().DateToString(tc.date, tc.opt).Build())
		})
	}
}
