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

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_logicalBuilder_And(t *testing.T) {
	t.Run("test And", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}}}}, NewBuilder().And("item", bsonx.D("$gt", []any{"$qty", 100}), bsonx.D("$lt", []any{"$qty", 250})).Build())
	})
}

func Test_logicalBuilder_AndWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{NewBuilder().GtWithoutKey("$qty", 100).Build(), NewBuilder().LtWithoutKey("$qty", 250).Build()},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().AndWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_logicalBuilder_Not(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}}}}, NewBuilder().Not("item", bsonx.D("$gt", []any{"$qty", 250})).Build())
}

func Test_logicalBuilder_NotWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{NewBuilder().GtWithoutKey("$qty", 250).Build()},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().NotWithoutKey(tc.expressions...).Build())
		})
	}
}

func Test_logicalBuilder_Or(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 50}}}}}}}}, NewBuilder().Or("item", bsonx.D("$gt", []any{"$qty", 250}), bsonx.D("$lt", []any{"$qty", 50})).Build())
}

func Test_logicalBuilder_OrWithoutKey(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{query.NewBuilder().Eq("x", 0).Build(), query.NewBuilder().Expr(NewBuilder().EqWithoutKey(NewBuilder().DivideWithoutKey(1, "$x").Build(), 3).Build()).Build()},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "x", Value: bson.D{bson.E{Key: "$eq", Value: 0}}}}, bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$eq", Value: []any{bson.D{bson.E{Key: "$divide", Value: []any{1, "$x"}}}, 3}}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, NewBuilder().OrWithoutKey(tc.expressions...).Build())
		})
	}
}
