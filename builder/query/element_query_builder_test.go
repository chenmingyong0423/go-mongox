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

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func Test_elementQueryBuilder_Exists(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}, BsonBuilder().Exists("name", true).Build())
}

func Test_elementQueryBuilder_Type(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: bson.TypeString}}}}, BsonBuilder().Type("name", bson.TypeString).Build())
}

func Test_elementQueryBuilder_TypeAlias(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$type", Value: "string"}}}}, BsonBuilder().TypeAlias("name", "string").Build())
}

func TestBsonBuilder_TypeArray(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []bsontype.Type

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: ([]bsontype.Type)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []bsontype.Type{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []bsontype.Type{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []bsontype.Type{bson.TypeString},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []bsontype.Type{bson.TypeString}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []bsontype.Type{bson.TypeString, bson.TypeInt32},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []bsontype.Type{bson.TypeString, bson.TypeInt32}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().TypeArray(tc.key, tc.ts...).Build())
		})
	}
}

func TestBsonBuilder_TypeArrayAlias(t *testing.T) {

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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: ([]string)(nil)}}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []string{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []string{}}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []string{"string"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []string{"string"}}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []string{"string", "int32"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.Type, Value: []string{"string", "int32"}}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().TypeArrayAlias(tc.key, tc.ts...).Build())
		})
	}
}
