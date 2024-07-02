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

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_arithmeticBuilder_Add(t *testing.T) {
	t.Run("test add", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			NewBuilder().Add("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}

func Test_arithmeticBuilder_AddWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$add", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$add", Value: []any{}}},
		},
		{
			name:        "single type",
			expressions: []any{1, 2, 3, 4},
			expected:    bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, 4}}},
		},
		{
			name:        "multiple types",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			expected:    bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, "$a", "$b", "$c"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().AddWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Multiply(t *testing.T) {
	t.Run("test multiply", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			NewBuilder().Multiply("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}

func Test_arithmeticBuilder_MultiplyWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$multiply", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$multiply", Value: []any{}}},
		},
		{
			name:        "single type",
			expressions: []any{1, 2, 3, 4},
			expected:    bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, 4}}},
		},
		{
			name:        "multiple types",
			expressions: []any{1, 2, 3, "$a", "$b", "$c"},
			expected:    bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, "$a", "$b", "$c"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().MultiplyWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Subtract(t *testing.T) {
	t.Run("test subtract", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "dateDifference", Value: bson.D{bson.E{Key: "$subtract", Value: []any{"$date", 5 * 60 * 1000}}}}},
			NewBuilder().Subtract("dateDifference", []any{"$date", 5 * 60 * 1000}...).Build(),
		)
	})

}

func Test_arithmeticBuilder_SubtractWithoutKey(t *testing.T) {
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
			assert.Equal(t, tc.expected, NewBuilder().SubtractWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Divide(t *testing.T) {
	t.Run("test divide", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$divide", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			NewBuilder().Divide("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}
func Test_arithmeticBuilder_DivideWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"hours", 8},
			expected:    bson.D{bson.E{Key: "$divide", Value: []any{"hours", 8}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().DivideWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Mod(t *testing.T) {
	t.Run("test mod", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$mod", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			NewBuilder().Mod("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}
func Test_arithmeticBuilder_ModWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$hours", "$tasks"},
			expected:    bson.D{bson.E{Key: "$mod", Value: []any{"$hours", "$tasks"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().ModWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Abs(t *testing.T) {
	testCases := []struct {
		name       string
		key        string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			key:        "hours",
			expression: "$hours",
			expected:   bson.D{bson.E{Key: "hours", Value: bson.D{bson.E{Key: "$abs", Value: "$hours"}}}},
		},
		{
			name:       "nested expression",
			key:        "tempDiff",
			expression: bson.M{"$subtract": []any{"$startTemp", "$endTemp"}},
			expected:   bson.D{bson.E{Key: "tempDiff", Value: bson.D{bson.E{Key: "$abs", Value: bson.M{"$subtract": []any{"$startTemp", "$endTemp"}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().Abs(tc.key, tc.expression).Build())
		})
	}
}

func Test_arithmeticBuilder_AbsWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$hours",
			expected:   bson.D{bson.E{Key: "$abs", Value: "$hours"}},
		},
		{
			name:       "nested expression",
			expression: bson.M{"$subtract": []any{"$startTemp", "$endTemp"}},
			expected:   bson.D{bson.E{Key: "$abs", Value: bson.M{"$subtract": []any{"$startTemp", "$endTemp"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().AbsWithoutKey(tc.expression).Build())
		})
	}
}
