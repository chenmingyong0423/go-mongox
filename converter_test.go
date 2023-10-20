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

package mongox

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestStructToSetBson(t *testing.T) {
	type testData struct {
		Name    string `bson:"name"`
		Age     int
		Unknown string `bson:"-"`
	}
	testCases := []struct {
		name  string
		data  any
		wantD bson.D
	}{
		{
			name: "nil data",
			data: nil,

			wantD: nil,
		},
		{
			name: "struct with zero-value",
			data: testData{},
			wantD: bson.D{
				bson.E{Key: set, Value: bson.D{
					bson.E{Key: "name", Value: ""},
					bson.E{Key: "age", Value: 0},
				}},
			},
		},
		{
			name: "struct with no zero-value",
			data: testData{Name: "cmy", Age: 24},
			wantD: bson.D{
				bson.E{Key: set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: 24},
				}},
			},
		},
		{
			name: "struct pointer with empty-value",
			data: &testData{},
			wantD: bson.D{
				bson.E{Key: set, Value: bson.D{
					bson.E{Key: "name", Value: ""},
					bson.E{Key: "age", Value: 0},
				}},
			},
		},
		{
			name: "struct pointer with no empty-value",
			data: &testData{Name: "cmy", Age: 24},
			wantD: bson.D{
				bson.E{Key: set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: 24},
				}},
			},
		},
		{
			name:  "not struct",
			data:  1,
			wantD: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if StructToSetBson(tc.data) != nil {
				assert.Equal(t, tc.wantD[0].Value, StructToSetBson(tc.data)[0].Value)
			} else {
				assert.Equal(t, tc.wantD, StructToSetBson(tc.data))
			}
		})
	}
}