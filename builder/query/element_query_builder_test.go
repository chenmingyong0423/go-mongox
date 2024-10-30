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

func Test_elementQueryBuilder_Exists(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}, NewBuilder().Exists("name", true).Build())
}

func Test_elementQueryBuilder_Type(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: bson.TypeString}}}}, NewBuilder().Type("name", bson.TypeString).Build())
}

func Test_elementQueryBuilder_TypeAlias(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: "string"}}}}, NewBuilder().TypeAlias("name", "string").Build())
}

func TestNewBuilder_TypeArray(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []bson.Type

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: ([]bson.Type)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []bson.Type{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []bson.Type{bson.TypeString},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{bson.TypeString}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []bson.Type{bson.TypeString, bson.TypeInt32},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []bson.Type{bson.TypeString, bson.TypeInt32}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBuilder().TypeArray(tc.key, tc.ts...).Build())
		})
	}
}

func TestNewBuilder_TypeArrayAlias(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []string

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: ([]string)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []string{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []string{"string"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{"string"}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []string{"string", "int32"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: TypeOp, Value: []string{"string", "int32"}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBuilder().TypeArrayAlias(tc.key, tc.ts...).Build())
		})
	}
}
