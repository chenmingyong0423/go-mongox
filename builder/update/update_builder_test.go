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

package update

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBuilder_Add(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{},
		},
		{
			name:      "odd params",
			keyValues: []any{"name", "cmy", "age"},
			want:      bson.D{},
		},
		{
			name:      "keys contain non-string",
			keyValues: []any{"name", "cmy", 2, "age"},
			want:      bson.D{bson.E{Key: "name", Value: "cmy"}},
		},
		{
			name:      "normal params",
			keyValues: []any{"name", "cmy", "age", 18, "scores", []int{100, 99, 98}},
			want:      bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 18}, bson.E{Key: "scores", Value: []int{100, 99, 98}}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, BsonBuilder().Add(testCase.keyValues...).Build())
		})
	}
}
