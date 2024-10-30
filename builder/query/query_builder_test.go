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

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestQuery(t *testing.T) {
	query := NewBuilder()
	assert.NotNil(t, query)
	assert.Equal(t, bson.D{}, query.Build())
}

func TestQueryBuilder_Id(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "_id", Value: "123"}}, NewBuilder().Id("123").Build())
}

func TestQueryBuilder_KeyValue(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 18}, bson.E{Key: "scores", Value: []int{100, 99, 98}}}, NewBuilder().KeyValue("name", "cmy").KeyValue("age", 18).KeyValue("scores", []int{100, 99, 98}).Build())
}

func TestBuilder_TryMergeValue(t *testing.T) {
	testCases := []struct {
		name     string
		builder  *Builder
		key      string
		value    bson.E
		wantBool bool
		wantBson bson.D
	}{
		{
			name:     "not merge when key is not exist",
			builder:  NewBuilder(),
			key:      "age",
			value:    bson.E{Key: LtOp, Value: 25},
			wantBool: false,
			wantBson: bson.D{},
		},
		{
			name:     "not merge when key is different",
			builder:  NewBuilder().Gt("age", 18),
			key:      "name",
			value:    bson.E{Key: EqOp, Value: "cmy"},
			wantBool: false,
			wantBson: bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}}}},
		},
		{
			name:     "merge when key is same",
			builder:  NewBuilder().Gt("age", 18),
			key:      "age",
			value:    bson.E{Key: LtOp, Value: 25},
			wantBool: true,
			wantBson: bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: GtOp, Value: 18}, bson.E{Key: LtOp, Value: 25}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantBool, tc.builder.tryMergeValue(tc.key, tc.value))
			assert.Equal(t, tc.wantBson, tc.builder.Build())
		})
	}
}
