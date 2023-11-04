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
			assert.Equal(t, tc.expected, BsonBuilder().Add(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Multiply(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Multiply(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Subtract(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		start    int64
		length   int64
		expected bson.D
	}{
		{
			name:     "normal",
			s:        "$quarter",
			start:    0,
			length:   2,
			expected: bson.D{bson.E{Key: "$subtract", Value: []any{"$quarter", int64(0), int64(2)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Subtract(tc.s, tc.start, tc.length).Build())
		})
	}
}

func Test_arithmeticBuilder_Divide(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Divide(tc.expressions...).Build())
		})
	}
}

func Test_arithmeticBuilder_Mod(t *testing.T) {
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
			assert.Equal(t, tc.expected, BsonBuilder().Mod(tc.expressions...).Build())
		})
	}
}
