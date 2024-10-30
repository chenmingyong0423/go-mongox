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
	"go.mongodb.org/mongo-driver/v2/bson"
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

func Test_arithmeticBuilder_Ceil(t *testing.T) {
	t.Run("test ceil", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "ceilingValue", Value: bson.D{bson.E{Key: "$ceil", Value: "$value"}}}},
			NewBuilder().Ceil("ceilingValue", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_CeilWithoutKey(t *testing.T) {
	t.Run("test ceil without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$ceil", Value: "$value"}}, NewBuilder().CeilWithoutKey("$value").Build())
	})
}

func Test_arithmeticBuilder_Exp(t *testing.T) {
	t.Run("test exp", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "expValue", Value: bson.D{bson.E{Key: "$exp", Value: "$value"}}}},
			NewBuilder().Exp("expValue", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_ExpWithoutKey(t *testing.T) {
	t.Run("test exp without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$exp", Value: "$interestRate"}}, NewBuilder().ExpWithoutKey("$interestRate").Build())
	})
}

func Test_arithmeticBuilder_Floor(t *testing.T) {
	t.Run("test floor", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "floorValue", Value: bson.D{bson.E{Key: "$floor", Value: "$value"}}}},
			NewBuilder().Floor("floorValue", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_FloorWithoutKey(t *testing.T) {
	t.Run("test floor without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$floor", Value: "$value"}}, NewBuilder().FloorWithoutKey("$value").Build())
	})
}

func Test_arithmeticBuilder_Ln(t *testing.T) {
	t.Run("test ln", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "lnValue", Value: bson.D{bson.E{Key: "$ln", Value: "$value"}}}},
			NewBuilder().Ln("lnValue", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_LnWithoutKey(t *testing.T) {
	t.Run("test ln without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$ln", Value: "$value"}}, NewBuilder().LnWithoutKey("$value").Build())
	})
}

func Test_arithmeticBuilder_Log(t *testing.T) {
	t.Run("test log", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "logValue", Value: bson.D{bson.E{Key: "$log", Value: bson.A{"$value", 2}}}}},
			NewBuilder().Log("logValue", "$value", 2).Build(),
		)
	})
}

func Test_arithmeticBuilder_LogWithoutKey(t *testing.T) {
	t.Run("test log without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$log", Value: bson.A{"$value", 2}}}, NewBuilder().LogWithoutKey("$value", 2).Build())
	})
}

func Test_arithmeticBuilder_Log10(t *testing.T) {
	t.Run("test log10", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "log10Value", Value: bson.D{bson.E{Key: "$log10", Value: bson.A{"$value", 10}}}}},
			NewBuilder().Log10("log10Value", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_Log10WithoutKey(t *testing.T) {
	t.Run("test log10 without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$log10", Value: bson.A{"$value", 10}}}, NewBuilder().Log10WithoutKey("$value").Build())
	})
}

func Test_arithmeticBuilder_Pow(t *testing.T) {
	t.Run("test pow", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "powValue", Value: bson.D{bson.E{Key: "$pow", Value: bson.A{"$value", 2}}}}},
			NewBuilder().Pow("powValue", "$value", 2).Build(),
		)
	})
}

func Test_arithmeticBuilder_PowWithoutKey(t *testing.T) {
	t.Run("test pow without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$pow", Value: bson.A{"$value", 2}}}, NewBuilder().PowWithoutKey("$value", 2).Build())
	})
}

func Test_arithmeticBuilder_Round(t *testing.T) {
	t.Run("test round", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "roundValue", Value: bson.D{bson.E{Key: "$round", Value: bson.A{"$value", 2}}}}},
			NewBuilder().Round("roundValue", "$value", 2).Build(),
		)
	})
}

func Test_arithmeticBuilder_RoundWithoutKey(t *testing.T) {
	t.Run("test round without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$round", Value: bson.A{"$value", 2}}}, NewBuilder().RoundWithoutKey("$value", 2).Build())
	})
}

func Test_arithmeticBuilder_Sqrt(t *testing.T) {
	t.Run("test sqrt", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "sqrtValue", Value: bson.D{bson.E{Key: "$sqrt", Value: "$value"}}}},
			NewBuilder().Sqrt("sqrtValue", "$value").Build(),
		)
	})
}

func Test_arithmeticBuilder_SqrtWithoutKey(t *testing.T) {
	t.Run("test sqrt without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$sqrt", Value: "$value"}}, NewBuilder().SqrtWithoutKey("$value").Build())
	})
}

func Test_arithmeticBuilder_Trunc(t *testing.T) {
	t.Run("test trunc", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "truncValue", Value: bson.D{bson.E{Key: "$trunc", Value: bson.A{"$value", 0}}}}},
			NewBuilder().Trunc("truncValue", "$value", 0).Build(),
		)
	})
}

func Test_arithmeticBuilder_TruncWithoutKey(t *testing.T) {
	t.Run("test trunc without key", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$trunc", Value: bson.A{"$value", 0}}}, NewBuilder().TruncWithoutKey("$value", 0).Build())
	})
}
