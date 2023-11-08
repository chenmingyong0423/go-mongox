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

	"github.com/chenmingyong0423/go-mongox/converter"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_accumulatorsBuilder_Sum(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Sum(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_SumMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$price", "$fee"},
			expected:    bson.D{{Key: "$sum", Value: []any{"$price", "$fee"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().SumMany(tc.expressions...).Build())
		})
	}
}

func Test_accumulatorsBuilder_Push(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name: "normal",
			// { item: "$item", quantity: "$quantity" }
			expression: BsonBuilder().AddKeyValues(converter.KeyValue[any]("item", "$item"), converter.KeyValue[any]("quantity", "$quantity")).Build(),
			expected:   bson.D{{Key: "$push", Value: bson.D{{Key: "item", Value: "$item"}, {Key: "quantity", Value: "$quantity"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Push(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Avg(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name: "normal",
			// { $multiply: [ "$price", "$quantity" ] }
			expression: BsonBuilder().Multiply("$price", "$quantity").Build(),
			expected:   bson.D{{Key: "$avg", Value: bson.D{{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Avg(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_AvgMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$price", "$fee"},
			expected:    bson.D{{Key: "$avg", Value: []any{"$price", "$fee"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().AvgMany(tc.expressions...).Build())
		})
	}
}

func Test_accumulatorsBuilder_First(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().First(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Last(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Last(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_Min(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Min(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_MinMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$price", "$fee"},
			expected:    bson.D{{Key: "$min", Value: []any{"$price", "$fee"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().MinMany(tc.expressions...).Build())
		})
	}
}

func Test_accumulatorsBuilder_Max(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Max(tc.expression).Build())
		})
	}
}

func Test_accumulatorsBuilder_MaxMany(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$price", "$fee"},
			expected:    bson.D{{Key: "$max", Value: []any{"$price", "$fee"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().MaxMany(tc.expressions...).Build())
		})
	}
}
