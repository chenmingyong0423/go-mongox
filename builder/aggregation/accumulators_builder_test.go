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

func Test_accumulatorsBuilder_Sum(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$price"}}}}, NewBuilder().Sum("totalAmount", "$price").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$price"}}}, bson.E{Key: "totalFee", Value: bson.D{{Key: "$sum", Value: "$fee"}}}}, NewBuilder().Sum("totalAmount", "$price").Sum("totalFee", "$fee").Build())
	})
}

func Test_accumulatorsBuilder_SumWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$price",
			expected:   bson.D{{Key: "$sum", Value: "$price"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().SumWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Push(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$push", Value: "$item"}}}}, NewBuilder().Push("items", "$item").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$push", Value: "$item"}}}, bson.E{Key: "types", Value: bson.D{{Key: "$push", Value: "$type"}}}}, NewBuilder().Push("items", "$item").Push("types", "$type").Build())
	})
}

func Test_accumulatorsBuilder_PushWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name: "normal",
			// { item: "$item", quantity: "$quantity" }
			expression: NewBuilder().KeyValue("item", "$item").KeyValue("quantity", "$quantity").Build(),
			expected:   bson.D{{Key: "$push", Value: bson.D{{Key: "item", Value: "$item"}, {Key: "quantity", Value: "$quantity"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().PushWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Avg(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "avgAmount", Value: bson.D{{Key: "$avg", Value: "$price"}}}}, NewBuilder().Avg("avgAmount", "$price").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "avgAmount", Value: bson.D{{Key: "$avg", Value: "$price"}}}, bson.E{Key: "avgFee", Value: bson.D{{Key: "$avg", Value: "$fee"}}}}, NewBuilder().Avg("avgAmount", "$price").Avg("avgFee", "$fee").Build())
	})
}

func Test_accumulatorsBuilder_AvgWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name: "normal",
			// { $multiply: [ "$price", "$quantity" ] }
			expression: NewBuilder().MultiplyWithoutKey("$price", "$quantity").Build(),
			expected:   bson.D{{Key: "$avg", Value: bson.D{{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().AvgWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_First(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "firstType", Value: bson.D{{Key: "$first", Value: "$type"}}}}, NewBuilder().First("firstType", "$type").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "firstType", Value: bson.D{{Key: "$first", Value: "$type"}}}, bson.E{Key: "firstPrice", Value: bson.D{{Key: "$first", Value: "$price"}}}}, NewBuilder().First("firstType", "$type").First("firstPrice", "$price").Build())
	})
}

func Test_accumulatorsBuilder_FirstWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$type",
			expected:   bson.D{{Key: "$first", Value: "$type"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().FirstWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Last(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "lastType", Value: bson.D{{Key: "$last", Value: "$type"}}}}, NewBuilder().Last("lastType", "$type").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "lastType", Value: bson.D{{Key: "$last", Value: "$type"}}}, bson.E{Key: "lastPrice", Value: bson.D{{Key: "$last", Value: "$price"}}}}, NewBuilder().Last("lastType", "$type").Last("lastPrice", "$price").Build())
	})
}

func Test_accumulatorsBuilder_LastWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$type",
			expected:   bson.D{{Key: "$last", Value: "$type"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().LastWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Min(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "minPrice", Value: bson.D{{Key: "$min", Value: "$price"}}}}, NewBuilder().Min("minPrice", "$price").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "minPrice", Value: bson.D{{Key: "$min", Value: "$price"}}}, bson.E{Key: "minFee", Value: bson.D{{Key: "$min", Value: "$fee"}}}}, NewBuilder().Min("minPrice", "$price").Min("minFee", "$fee").Build())
	})
}
func Test_accumulatorsBuilder_MinWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$price",
			expected:   bson.D{{Key: "$min", Value: "$price"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().MinWithoutKey(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Max(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "maxPrice", Value: bson.D{{Key: "$max", Value: "$price"}}}}, NewBuilder().Max("maxPrice", "$price").Build())
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "maxPrice", Value: bson.D{{Key: "$max", Value: "$price"}}}, bson.E{Key: "maxFee", Value: bson.D{{Key: "$max", Value: "$fee"}}}}, NewBuilder().Max("maxPrice", "$price").Max("maxFee", "$fee").Build())
	})
}

func Test_accumulatorsBuilder_MaxWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal",
			expression: "$price",
			expected:   bson.D{{Key: "$max", Value: "$price"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().MaxWithoutKey(tc.expression).Build())
		})
	}
}
