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

package filter

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestM(t *testing.T) {
	assert.Equal(t, bson.M{"name": "chenmingyong"}, M("name", "chenmingyong"))
}

func TestD(t *testing.T) {
	testCases := []struct {
		name     string
		input    []types.KeyValue
		expected bson.D
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: bson.D{},
		},
		{
			name:     "empty input",
			input:    []types.KeyValue{},
			expected: bson.D{},
		},
		{
			name: "one element",
			input: []types.KeyValue{
				KV("name", "chenmingyong"),
			},
			expected: bson.D{
				{Key: "name", Value: "chenmingyong"},
			},
		},
		{
			name: "many elements",
			input: []types.KeyValue{
				KV("name", "chenmingyong"),
				KV("age", 24),
			},
			expected: bson.D{
				{Key: "name", Value: "chenmingyong"},
				{Key: "age", Value: 24},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, D(tc.input...))
		})
	}
}
