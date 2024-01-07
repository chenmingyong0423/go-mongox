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

func Test_stringBuilder_Concat(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$item", " - ", "$description"}}}}}, BsonBuilder().Concat("item", "$item", " - ", "$description").Build())
	})
}

func Test_stringBuilder_ConcatWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{{Key: "$concat", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{{Key: "$concat", Value: []any{}}},
		},
		{
			name:        "normal expression",
			expressions: []any{"$item", " - ", "$description"},
			expected:    bson.D{{Key: "$concat", Value: []any{"$item", " - ", "$description"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ConcatWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_stringBuilder_SubstrBytes(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "quarterSubtring", Value: bson.D{{Key: "$substrBytes", Value: []any{"$quarter", int64(0), int64(2)}}}}}, BsonBuilder().SubstrBytes("quarterSubtring", "$quarter", int64(0), int64(2)).Build())
	})
}

func Test_stringBuilder_SubstrBytesWithoutKey(t *testing.T) {
	testCases := []struct {
		name             string
		stringExpression string
		byteIndex        int64
		byteCount        int64
		expected         bson.D
	}{
		{
			name:             "normal expression",
			stringExpression: "$quarter",
			byteIndex:        0,
			byteCount:        2,
			expected:         bson.D{{Key: "$substrBytes", Value: []any{"$quarter", int64(0), int64(2)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().SubstrBytesWithoutKey(tc.stringExpression, tc.byteIndex, tc.byteCount).Build())
		})
	}
}

func Test_stringBuilder_ToLower(t *testing.T) {
	t.Run("test ToLower", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{{Key: "$toLower", Value: "$item"}}}}, BsonBuilder().ToLower("item", "$item").Build())
	})
}

func Test_stringBuilder_ToLowerWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal expression",
			expression: "$item",
			expected:   bson.D{{Key: "$toLower", Value: "$item"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ToLowerWithoutKey(tc.expression).Build())
		})
	}
}

func Test_stringBuilder_ToUpper(t *testing.T) {
	t.Run("test ToUpper", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{{Key: "$toUpper", Value: "$item"}}}}, BsonBuilder().ToUpper("item", "$item").Build())
	})
}

func Test_stringBuilder_ToUpperWithoutKey(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "normal expression",
			expression: "$item",
			expected:   bson.D{{Key: "$toUpper", Value: "$item"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ToUpperWithoutKey(tc.expression).Build())
		})
	}
}

func Test_stringBuilder_Contact(t *testing.T) {
	t.Run("test Contact", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$item", " - ", "$description"}}}}}, BsonBuilder().Contact("item", "$item", " - ", "$description").Build())
	})
}

func Test_stringBuilder_ContactWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "normal",
			expressions: []any{"$item", " - ", "$description"},
			expected:    bson.D{{Key: "$concat", Value: []any{"$item", " - ", "$description"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ContactWithoutKey(tc.expressions...).Build())
		})
	}
}
